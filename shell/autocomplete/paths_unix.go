//go:build !windows && !plan9
// +build !windows,!plan9

package autocomplete

import (
	"os"
	"strings"

	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/consts"
	"github.com/phayes/permbits"
)

func pathIsLocal(s string) bool {
	return strings.HasPrefix(s, consts.PathSlash) ||
		strings.HasPrefix(s, "."+consts.PathSlash) ||
		strings.HasPrefix(s, ".."+consts.PathSlash)
}

func matchDirsOnce(s string) (items []string) {
	s = variables.ExpandString(s)
	path, partial := partialPath(s)

	var dirs []string

	files, _ := os.ReadDir(path)
	for _, f := range files {
		if f.IsDir() && (f.Name()[0] != '.' ||
			(len(partial) > 0 && partial[0] == '.')) {
			dirs = append(dirs, f.Name()+consts.PathSlash)
			continue
		}

		fi, _ := f.Info()
		perm := permbits.FileMode(fi.Mode())
		switch {
		case perm.OtherExecute() && fi.Mode()&os.ModeSymlink != 0:
			ln, err := os.Readlink(path + consts.PathSlash + f.Name())
			if err != nil {
				continue
			}
			if ln[0] != consts.PathSlash[0] {
				ln = path + consts.PathSlash + ln
			}
			info, err := os.Lstat(ln)
			if err != nil {
				continue
			}
			perm := permbits.FileMode(info.Mode())
			if perm.OtherExecute() && info.Mode().IsDir() {
				dirs = append(dirs, f.Name()+consts.PathSlash)
			}

		default:
			/*|| perm.GroupExecute()||perm.UserExecute() need to check what user and groups you are in first */
		}
	}

	if path != consts.PathSlash {
		dirs = append(dirs, ".."+consts.PathSlash)
	}

	for i := range dirs {
		if strings.HasPrefix(dirs[i], partial) {
			items = append(items, dirs[i][len(partial):])
		}
	}

	return
}
