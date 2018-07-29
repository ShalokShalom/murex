package readline

type TabDisplayType int

const (
	TabDisplayGrid = iota
	TabDisplayList
	TabDisplayMap
)

func (rl *Instance) getTabCompletion() {
	rl.tcOffset = 0
	if rl.TabCompleter == nil {
		return
	}

	rl.tcPrefix, rl.tcSuggestions, rl.tcDescriptions, rl.tcDisplayType = rl.TabCompleter(rl.line, rl.pos)
	if len(rl.tcSuggestions) == 0 {
		return
	}

	if len(rl.tcDescriptions) == 0 {
		// probably not needed, but just in case someone doesn't initialise the
		// map in their API call.
		rl.tcDescriptions = make(map[string]string)
	}

	/*if len(rl.tcSuggestions) == 1 && !rl.modeTabCompletion {
		if len(rl.tcSuggestions[0]) == 0 || rl.tcSuggestions[0] == " " || rl.tcSuggestions[0] == "\t" {
			return
		}
		rl.insert([]byte(rl.tcSuggestions[0]))
		return
	}*/

	rl.initTabCompletion()
}

func (rl *Instance) initTabCompletion() {
	if rl.tcDisplayType == TabDisplayGrid {
		rl.initTabGrid()
	} else {
		rl.initTabMap()
	}
}

func (rl *Instance) moveTabCompletionHighlight(x, y int) {
	if rl.tcDisplayType == TabDisplayGrid {
		rl.moveTabGridHighlight(x, y)
	} else {
		rl.moveTabMapHighlight(x, y)
	}
}

func (rl *Instance) writeTabCompletion() {
	if !rl.modeTabCompletion {
		return
	}

	switch rl.tcDisplayType {
	case TabDisplayGrid:
		rl.writeTabGrid()

	case TabDisplayMap:
		rl.writeTabMap()

	case TabDisplayList:
		rl.writeTabMap()

	default:
		rl.writeTabGrid()
	}
}

func (rl *Instance) resetTabCompletion() {
	rl.modeTabCompletion = false
	rl.tcOffset = 0
	rl.tcUsedY = 0
	rl.modeTabFind = false
	rl.tfLine = []rune{}
}
