package expressions

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/mkarray"
	_ "github.com/lmorg/murex/builtins/types/generic"
	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseObject(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ObjectBegin,
		tests: []expTestT{
			{
				input:    `%{foo:}`,
				expected: `{"foo":null}`,
				pos:      5,
			},
			{
				input:    `%{foo: bar}`,
				expected: `{"foo":"bar"}`,
				pos:      9,
			},
			{
				input:    `%{a:b}`,
				expected: `{"a":"b"}`,
				pos:      4,
			},
			{
				input:    `%{a:[1,2]}`,
				expected: `{"a":[1,2]}`,
				pos:      8,
			},
			{
				input:    `%{a:%[1,2]}`,
				expected: `{"a":[1,2]}`,
				pos:      9,
			},
			{
				input:    `%{neg:-1}`,
				expected: `{"neg":-1}`,
				pos:      7,
			},
			{
				input:    `%{-2:-2,1:1,0:0,3.4:3.4}`,
				expected: `{"-2":-2,"0":0,"1":1,"3.4":3.4}`,
				pos:      22,
			},
			{
				input:    `%{'foo':"bar"}`,
				expected: `{"foo":"bar"}`,
				pos:      12,
			},
			{
				input:    `%{'foo':"bar",a:{1:a,2:b,3:c}}`,
				expected: `{"a":{"1":"a","2":"b","3":"c"},"foo":"bar"}`,
				pos:      28,
			},
			{
				input:    `%{'foo':"bar",a:%{1:a,2:b,3:c}}`,
				expected: `{"a":{"1":"a","2":"b","3":"c"},"foo":"bar"}`,
				pos:      29,
			},
			{
				input:    `%{a:$a,b:@b}`,
				expected: `{"a":null,"b":null}`,
				pos:      10,
			},
			{
				input:    `%{a:$a,b:[@b]}`,
				expected: `{"a":null,"b":null}`,
				pos:      12,
			},
			{
				input:    `%{a:$a,b:%[@b]}`,
				expected: `{"a":null,"b":null}`,
				pos:      13,
			},
			{
				input:    `%{nan:-}`,
				expected: `{"NaN":"-"}`,
				pos:      6,
			},
			{
				input:    `%{nan:-one}`,
				expected: `{"NaN":"-one"}`,
				pos:      9,
			},
		},
	}

	testParserObject(t, tests)
}

func TestParseObjectBadGrammar(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ObjectBegin,
		tests: []expTestT{
			{
				input: `%{foo}`,
				error: true,
			},
			{
				input: `%{foo::bar}`,
				error: true,
			},
			{
				input: `%{foo:bar,,}`,
				error: true,
			},
			{
				input: `%{foo:bar`,
				error: true,
			},
		},
	}

	testParserObject(t, tests)
}

func TestParseObjectLf(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ObjectBegin,
		tests: []expTestT{
			{
				input:    "%{\nfoo:bar}",
				expected: `{"foo":"bar"}`,
				pos:      9,
			},
			{
				input:    "%{\n foo:bar}",
				expected: `{"foo":"bar"}`,
				pos:      10,
			},
		},
	}

	testParserObject(t, tests)
}

func TestParseObjectBool(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ObjectBegin,
		tests: []expTestT{
			{
				input:    "%{foo:true}",
				expected: `{"foo":true}`,
				pos:      9,
			},
			{
				input:    "%{foo:false}",
				expected: `{"foo":false}`,
				pos:      10,
			},
		},
	}

	testParserObject(t, tests)
}
