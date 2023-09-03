package expressions

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/alter"
)

func scalarNameDetokenised(r []rune) []rune {
	if r[1] == '(' {
		return r[2 : len(r)-1]
	} else {
		return r[1:]
	}
}

func convertScalarToBareword(node *astNodeT) {
	if node.key == symbols.Scalar && len(node.value) > 1 &&
		node.value[0] == '$' && node.value[1] != '{' {

		node.key = symbols.Bareword
		node.value = scalarNameDetokenised(node.value)
	}
}

/*func expAssign(tree *ParserT, _ bool) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	convertScalarToBareword(left)

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key <= symbols.Bareword {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	var v interface{}
	switch right.dt.Primitive {
	case primitives.Array, primitives.Object:
		b, err := json.Marshal(right.dt.Value)
		if err != nil {
			return err
		}
		v = string(b)

	default:
		v = right.dt.Value
	}

	err = tree.setVar(left.value, v, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		},
	})
}*/

func expAssign(tree *ParserT, overwriteType bool) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	convertScalarToBareword(left)

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key <= symbols.Bareword {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	var (
		v  interface{}
		dt string
	)

	switch right.dt.Primitive {
	case primitives.Array, primitives.Object:
		if overwriteType {
			dt = types.Json
		} else {
			dt = tree.p.Variables.GetDataType(left.Value())
			if dt == "" {
				dt = types.Json
			}
		}

		// this is ugly but Go's JSON marshaller is better behaved than Murexes on with empty values
		if dt == types.Json {
			b, err := json.Marshal(right.dt.Value)
			if err != nil {
				raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
			}
			v = string(b)
		} else {
			b, err := lang.MarshalData(tree.p, dt, right.dt.Value)
			if err != nil {
				raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
			}
			v = string(b)
		}

	default:
		if overwriteType {
			dt = right.dt.DataType()
			v = right.dt.Value
		} else {
			dt = tree.p.Variables.GetDataType(left.Value())
			//panic("-->" + dt + "<--")
			if dt == "" || dt == types.Null {
				dt = right.dt.DataType()
				v = right.dt.Value
			} else {
				//panic(fmt.Sprintf("-->%s||%v<--", dt, right.dt.Value))
				v, err = types.ConvertGoType(right.dt.Value, dt)
				if err != nil {
					raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
				}
			}
		}
	}

	err = tree.setVar(left.value, v, dt)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}

func expAssignAdd(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	convertScalarToBareword(left)

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	/*if right.key != symbols.Number {
		return raiseError(tree.expression,tree.currentSymbol(), fmt.Sprintf(
			"right side should not be a %s", right.key))
	}*/

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
			// var doesn't exist and we have strict types disabled so lets create var
			v, dt, err = float64(0), types.Number, nil
		} else {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}
	}

	var result interface{}

	switch dt {
	case types.Number, types.Float:
		if right.dt.Primitive != primitives.Number {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = v.(float64) + right.dt.Value().(float64)

	case types.Integer:
		if right.dt.Primitive != primitives.Number {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = float64(v.(int)) + right.dt.Value().(float64)

	case types.Boolean:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))

	case types.Null:
		switch right.dt.Primitive {
		case primitives.String:
			result = right.dt.Value().(string)
		case primitives.Number:
			result = right.dt.Value().(float64)
		default:
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}

	default:
		if right.dt.Primitive != primitives.String {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = v.(string) + right.dt.Value().(string)
	}

	err = tree.setVar(left.value, result, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}

func expAssignSubtract(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	convertScalarToBareword(left)

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
			// var doesn't exist and we have strict types disabled so lets create var
			v, dt, err = float64(0), types.Number, nil
		} else {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) - right.dt.Value().(float64)
	case types.Integer:
		f = float64(v.(int)) - right.dt.Value().(float64)
	case types.Null:
		f = 0 - right.dt.Value().(float64)
	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(left.value, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}

func expAssignMultiply(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	convertScalarToBareword(left)

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
			// var doesn't exist and we have strict types disabled so lets create var
			v, dt, err = float64(0), types.Number, nil
		} else {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) * right.dt.Value().(float64)
	case types.Integer:
		f = float64(v.(int)) * right.dt.Value().(float64)
	case types.Null:
		f = 0
	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(left.value, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}

func expAssignDivide(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	convertScalarToBareword(left)

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
			// var doesn't exist and we have strict types disabled so lets create var
			v, dt, err = float64(0), types.Number, nil
		} else {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) / right.dt.Value().(float64)
	case types.Integer:
		f = float64(v.(int)) / right.dt.Value().(float64)
	case types.Null:
		f = 0 / right.dt.Value().(float64)
	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(left.value, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}

func expAssignMerge(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	convertScalarToBareword(left)

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	rightVal := right.dt.Value()
	if right.dt.Primitive != primitives.String && reflect.TypeOf(rightVal).Kind() == reflect.String {
		rightVal, err = lang.UnmarshalDataBuffered(tree.p, []byte(rightVal.(string)), right.dt.MxDT)
		if err != nil {
			return err
		}
	}

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
			// var doesn't exist and we have strict types disabled so lets create var
			err = tree.setVar(left.value, rightVal, right.dt.DataType())
			if err != nil {
				return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
			}
			return tree.foldAst(&astNodeT{
				key: symbols.Calculated,
				pos: tree.ast[tree.astPos].pos,
				dt:  primitives.NewPrimitive(primitives.Null, nil),
			})
		} else {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}
	}

	merged, err := alter.Merge(tree.p.Context, v, nil, rightVal)
	if err != nil {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"cannot perform merge '%s' into '%s': %s",
			right.Value(), left.Value(),
			err.Error()))
	}

	err = tree.setVar(left.value, merged, dt)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}
