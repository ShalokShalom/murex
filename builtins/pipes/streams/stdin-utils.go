package streams

import (
	"github.com/lmorg/murex/lang/types"
)

// Shamelessly stolen from https://blog.golang.org/go-slices-usage-and-internals
// (it works well so why reinvent the wheel?)
func appendBytes(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// IsTTY returns false because the Stdin stream is not a pseudo-TTY
func (stdin *Stdin) IsTTY() bool { return false }

// Stats provides real time stream stats. Useful for progress bars etc.
func (stdin *Stdin) Stats() (bytesWritten, bytesRead uint64) {
	stdin.mutex.Lock()
	bytesWritten = stdin.bWritten
	bytesRead = stdin.bRead
	stdin.mutex.Unlock()
	return
}

// GetDataType returns the murex data type for the stream.Io interface
func (stdin *Stdin) GetDataType() (dt string) {
	for {
		select {
		case <-stdin.ctx.Done():
			return types.Generic
		default:
		}

		stdin.dtLock.Lock()
		dt = stdin.dataType
		stdin.dtLock.Unlock()
		if dt != "" {
			return
		}
	}
}

// SetDataType defines the murex data type for the stream.Io interface
func (stdin *Stdin) SetDataType(dt string) {
	stdin.dtLock.Lock()
	stdin.dataType = dt
	stdin.dtLock.Unlock()
	return
}

// DefaultDataType defines the murex data type for the stream.Io interface if it's not already set
func (stdin *Stdin) DefaultDataType(err bool) {
	stdin.dtLock.Lock()
	dt := stdin.dataType
	stdin.dtLock.Unlock()

	if dt == "" {
		if err {
			stdin.dtLock.Lock()
			stdin.dataType = types.Null
			stdin.dtLock.Unlock()
		} else {
			stdin.dtLock.Lock()
			stdin.dataType = types.Generic
			stdin.dtLock.Unlock()
		}
	}
}