//go:build !cgo_sqlite3
// +build !cgo_sqlite3

/*
	This file uses a pure Go driver for sqlite. Unlike lib_c.go, this one does
	not require cgo. For this reason it is the default option for custom builds
	however any pre-built binaries on Murex's website will be compiled against
	the C driver for sqlite.
*/

package sqlselect

import (
	_ "modernc.org/sqlite"
)

const driverName = "sqlite"
