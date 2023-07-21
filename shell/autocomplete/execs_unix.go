//go:build !windows && !plan9
// +build !windows,!plan9

package autocomplete

import (
	"os"
	"strings"

	"github.com/lmorg/murex/utils/consts"
	"github.com/phayes/permbits"
)

func listExes(path string, exes map[string]bool) {
	// if WSL then using a Windows lookup rather than POSIX
	for i := range wslMounts {
		if strings.HasPrefix(path, wslMounts[i]) {
			listExesWindows(path, exes)
			return
		}
	}

	// POSIX lookup
	files, _ := os.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fi, _ := f.Info()
		perm := permbits.FileMode(fi.Mode())
		switch {
		case perm.OtherExecute() && fi.Mode().IsRegular():
			exes[f.Name()] = true
		case perm.OtherExecute() && fi.Mode()&os.ModeSymlink != 0:
			ln, err := os.Readlink(path + consts.PathSlash + f.Name())
			if err != nil {
				continue
			}
			if ln[0] != consts.PathSlash[0] {
				ln = path + consts.PathSlash + ln
			}
			info, err := os.Stat(ln)
			if err != nil {
				continue
			}
			perm := permbits.FileMode(info.Mode())
			if perm.OtherExecute() && info.Mode().IsRegular() {
				exes[f.Name()] = true
			}

		default:
			/*|| perm.GroupExecute()||perm.UserExecute() need to check what user and groups you are in first */
		}
	}
}

func matchExes(s string, exes map[string]bool, includeColon bool) (items []string) {
	colon := ""
	// We only want a colon added if the exe is the function call rather than a
	// functions parameter (eg `some-exec` vs `sudo some-exec`).
	if includeColon {
		colon = ":"
	}

	for name := range exes {
		if strings.HasPrefix(name, s) {
			if name != consts.NamedPipeProcName {
				if !isSpecialBuiltin(name) {
					name = name + colon
				}
				items = append(items, name[len(s):])
			}
		}
	}

	return
}
