package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func marshal(p *lang.Process, v any) ([]byte, error) {
	switch t := v.(type) {
	case [][]string:
		var i int
		table := make([]map[string]any, len(t)-1)
		err := types.Table2Map(t, func(m map[string]any) error {
			table[i] = m
			i++
			return nil
		})
		if err != nil {
			return nil, err
		}
		return json.Marshal(table, p.Stdout.IsTTY())
	default:
		return json.Marshal(v, p.Stdout.IsTTY())
	}
}

func unmarshal(p *lang.Process) (v any, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &v)

	// Initially I really liked the idea of JSON files automatically falling
	// back to jsonlines. However on reflection I think it is a bad idea
	// because it then means we need to cover any and all instances where JSON
	// is read but not calling unmarshal - which will be plentiful - else we
	// end up with inconsistent and confusing behavior. But in the current
	// modal all we need to do is the following (see below) so we're not
	// really saving a significant amount of effort.
	//
	//     open ~/jsonlines.json -> cast jsonl -> format json
	/*if err.Error() == "invalid character '{' after top-level value" {
		// ^ this needs a test case so we catch any failures should Go ever
		// change the reported error message
		if jsonl, errJl := unmarshalJsonLines(b); errJl != nil {
			debug.Json(err.Error(), jsonl)
			return jsonl, nil
		}
	}*/

	return
}

/*func unmarshalJsonLines(b []byte) (v interface{}, err error) {
	var jsonl []interface{}

	lines := bytes.Split(b, []byte{'\n'})
	for _, line := range lines {
		err = json.Unmarshal(line, &v)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal index %d in jsonlines: %s", len(jsonl), err)
		}
		jsonl = append(jsonl, v)
	}

	return jsonl, err
}
*/
