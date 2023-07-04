package parameters_test

import (
	"testing"

	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type flagTest struct {
	Parameters []string
	Arguments  parameters.Arguments

	ExpFlags      map[string]string
	ExpAdditional []string

	Error bool
}

func TestFlags(t *testing.T) {
	tests := []flagTest{
		{
			Parameters: []string{"--string", "--test", "--number", "-5"},
			Arguments: parameters.Arguments{
				//AllowAdditional: true,
				Flags: map[string]string{
					"--string": types.String,
					"--number": types.Number,
				},
			},
			ExpFlags: map[string]string{
				"--string": "--test",
				"--number": "-5",
			},
			ExpAdditional: nil,
			Error:         false,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.ExpFlags == nil {
			test.ExpFlags = make(map[string]string)
		}
		if test.ExpAdditional == nil {
			test.ExpAdditional = make([]string, 0)
		}

		flags, additional, err := parameters.ParseFlags(test.Parameters, &test.Arguments)

		if flags == nil {
			flags = make(map[string]string)
		}
		if additional == nil {
			additional = make([]string, 0)
		}

		var failed bool

		if (err != nil) != test.Error {
			t.Errorf("Unexpected error in test %d:", i)
			failed = true
		}

		if json.LazyLogging(test.ExpFlags) != json.LazyLogging(flags) {
			t.Errorf("Flags doesn't match expected in test %d:", i)
			failed = true
		}

		if json.LazyLogging(test.ExpAdditional) != json.LazyLogging(additional) {
			t.Errorf("Additional doesn't match expected in test %d:", i)
			failed = true
		}

		if failed {
			t.Logf("  Parameters: %s", json.LazyLogging(test.Parameters))
			t.Logf("  Arguments:  %s", json.LazyLogging(test.Arguments))
			t.Logf("  exp flags:  %s", json.LazyLogging(test.ExpFlags))
			t.Logf("  act flags:  %s", json.LazyLogging(flags))
			t.Logf("  exp aditnl: %s", json.LazyLogging(test.ExpAdditional))
			t.Logf("  act aditnl: %s", json.LazyLogging(additional))
			t.Logf("  exp error:  %v", test.Error)
			t.Logf("  act error:  %v", err)
		}
	}
}
