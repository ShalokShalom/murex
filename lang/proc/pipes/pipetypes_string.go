// Code generated by "stringer -type=PipeTypes"; DO NOT EDIT

package pipes

import "fmt"

const _PipeTypes_name = "pipeUndefinedpipeNullpipeStreampipeFileWriterpipeNetDialerpipeNetListener"

var _PipeTypes_index = [...]uint8{0, 13, 21, 31, 45, 58, 73}

func (i PipeTypes) String() string {
	if i < 0 || i >= PipeTypes(len(_PipeTypes_index)-1) {
		return fmt.Sprintf("PipeTypes(%d)", i)
	}
	return _PipeTypes_name[_PipeTypes_index[i]:_PipeTypes_index[i+1]]
}
