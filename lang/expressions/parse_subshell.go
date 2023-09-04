package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

func (tree *ParserT) parseSubShell(exec bool, prefix rune, strOrVal varFormatting) ([]rune, primitives.FunctionT, error) {
	start := tree.charPos

	tree.charPos += 2

	_, err := tree.parseBlockQuote()
	if err != nil {
		return nil, nil, err
	}

	value := tree.expression[start : tree.charPos+1]
	block := tree.expression[start+2 : tree.charPos]

	if !exec {
		return value, nil, nil
	}

	switch prefix {
	case '$':
		fn := func() (any, string, error) {
			return execSubShellScalar(tree, block, strOrVal)
		}
		return value, fn, nil

	case '@':
		fn := func() (any, string, error) {
			v, err := execSubShellArray(tree, block, strOrVal)
			return v, types.Json, err
		}
		return value, fn, nil
	default:
		err = fmt.Errorf("invalid prefix in expression '%s'. %s", string(prefix), consts.IssueTrackerURL)
		return nil, nil, err
	}
}

func execSubShellScalar(tree *ParserT, block []rune, strOrVal varFormatting) (interface{}, string, error) {
	if tree.p == nil {
		panic("`tree.p` is undefined")
	}
	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_PARENT_VARTABLE)
	exitNum, err := fork.Execute(block)
	if err != nil {
		return nil, "", fmt.Errorf("subshell failed: %s", err.Error())
	}
	if exitNum > 0 && tree.p.RunMode.IsStrict() {
		return nil, "", fmt.Errorf("subshell exit status %d", exitNum)
	}
	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return nil, "", fmt.Errorf("cannot read from subshell's STDOUT: %s", err.Error())
	}

	b = utils.CrLfTrim(b)
	dataType := fork.Stdout.GetDataType()

	v, err := formatBytes(b, dataType, strOrVal)
	return v, dataType, err
}

func execSubShellArray(tree *ParserT, block []rune, strOrVal varFormatting) ([]interface{}, error) {
	var slice []interface{}

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_PARENT_VARTABLE)
	exitNum, err := fork.Execute(block)
	if err != nil {
		return nil, fmt.Errorf("subshell failed: %s", err.Error())
	}
	if exitNum > 0 && tree.p.RunMode.IsStrict() {
		return nil, fmt.Errorf("subshell exit status %d", exitNum)
	}

	switch strOrVal {
	case varAsString:
		err = fork.Stdout.ReadArray(tree.p.Context, func(b []byte) {
			slice = append(slice, string(b))
		})
	case varAsValue:
		err = fork.Stdout.ReadArrayWithType(tree.p.Context, func(v interface{}, _ string) {
			slice = append(slice, v)
		})
	default:
		panic("invalid value set for strOrVal")
	}
	if err != nil {
		return nil, fmt.Errorf("cannot read from subshell's STDOUT: %s", err.Error())
	}

	if len(slice) == 0 && tree.StrictArrays() {
		return nil, fmt.Errorf(errEmptyArray, "{"+string(block)+"}")
	}

	return slice, nil
}
