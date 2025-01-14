// Code generated by "stringer -type=RunMode"; DO NOT EDIT.

package runmode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Default-0]
	_ = x[Normal-1]
	_ = x[Evil-2]
	_ = x[BlockTry-3]
	_ = x[BlockTryPipe-4]
	_ = x[ModuleTry-5]
	_ = x[ModuleTryPipe-6]
	_ = x[FunctionTry-7]
	_ = x[FunctionTryPipe-8]
}

const _RunMode_name = "DefaultNormalEvilBlockTryBlockTryPipeModuleTryModuleTryPipeFunctionTryFunctionTryPipe"

var _RunMode_index = [...]uint8{0, 7, 13, 17, 25, 37, 46, 59, 70, 85}

func (i RunMode) String() string {
	if i < 0 || i >= RunMode(len(_RunMode_index)-1) {
		return "RunMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RunMode_name[_RunMode_index[i]:_RunMode_index[i+1]]
}
