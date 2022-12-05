package expressions

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/generic"
	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseArray(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ArrayBegin,
		tests: []expTestT{
			{
				input:    `%[1 2 3]`,
				expected: `[1,2,3]`,
				pos:      6,
			},
			{
				input: `%[1 2 3`,
				error: true,
			},
			{
				input:    `%[1  2  3]`,
				expected: `[1,2,3]`,
				pos:      8,
			},
			{
				input:    `%[1,2,3]`,
				expected: `[1,2,3]`,
				pos:      6,
			},
			{
				input:    `%[1, 2, 3]`,
				expected: `[1,2,3]`,
				pos:      8,
			},
			{
				input:    `%[  1  ,  2  ,  3  ]`,
				expected: `[1,2,3]`,
				pos:      18,
			},
			/////
			{
				input:    `%[  foo  ,  bar  ,  baz  ]`,
				expected: `["foo","bar","baz"]`,
				pos:      24,
			},
			/////
			{
				input: `%[  [1 2 3]  ,  [ "foo" "bar" ]`,
				error: true,
			},
			{
				input:    `%[%]`,
				expected: `["%"]`,
				pos:      2,
			},
			{
				input:    `%[  [1 2 3]  ,  [ "foo" "bar" ] ]`,
				expected: `[[1,2,3],["foo","bar"]]`,
				pos:      31,
			},
			{
				input:    `%[  %[1 2 3]  ,  %[ "foo" "bar" ] ]`,
				expected: `[[1,2,3],["foo","bar"]]`,
				pos:      33,
			},
			/////
			{
				input:    "%[\n\t1\n\t2\n\t3\n]",
				expected: `[1,2,3]`,
				pos:      11,
			},
			/////
			{
				input:    "%[$TestParseArray]",
				expected: `[""]`,
				pos:      16,
			},
			{
				input:    "%[@TestParseArray]",
				expected: `null`,
				pos:      16,
			},
			{
				input:    "%[[@TestParseArray]]",
				expected: `[null]`,
				pos:      18,
			},
		},
	}

	testParserObject(t, tests)
}