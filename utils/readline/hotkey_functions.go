package readline

import "fmt"

func HkFnMoveToStartOfLine(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.line.RuneLen() == 0 {
		return
	}
	rl.clearHelpers()
	rl.line.SetCellPos(0)
	rl.echo()
	moveCursorForwards(1)
}

func HkFnMoveToEndOfLine(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.line.RuneLen() == 0 {
		return
	}
	rl.clearHelpers()
	rl.line.SetRunePos(rl.line.RuneLen())
	rl.echo()
	moveCursorForwards(1)
}

func HkFnClearAfterCursor(rl *Instance) {
	if rl.line.RuneLen() == 0 {
		return
	}
	rl.clearHelpers()
	rl.line.Set(rl.line.Runes()[:rl.line.RunePos()])
	rl.echo()
	moveCursorForwards(1)
}

func HkFnClearScreen(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.previewMode != previewModeClosed {
		HkFnPreviewToggle(rl)
	}
	print(seqSetCursorPosTopLeft + seqClearScreen)
	rl.echo()
	rl.renderHelpers()
}

func HkFnClearLine(rl *Instance) {
	rl.clearPrompt()
	rl.resetHelpers()
}

func HkFnFuzzyFind(rl *Instance) {
	rl.viUndoSkipAppend = true
	if !rl.modeTabCompletion {
		rl.modeAutoFind = true
		rl.getTabCompletion()
	}

	rl.modeTabFind = true
	rl.updateTabFind([]rune{})
}

func HkFnSearchHistory(rl *Instance) {
	rl.viUndoSkipAppend = true
	rl.modeAutoFind = true
	rl.tcOffset = 0
	rl.modeTabCompletion = true
	rl.tcDisplayType = TabDisplayMap
	rl.tabMutex.Lock()
	rl.tcSuggestions, rl.tcDescriptions = rl.autocompleteHistory()
	rl.tabMutex.Unlock()
	rl.initTabCompletion()

	rl.modeTabFind = true
	rl.updateTabFind([]rune{})
}

func HkFnAutocomplete(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.modeTabCompletion {
		rl.moveTabCompletionHighlight(1, 0)
	} else {
		rl.getTabCompletion()
	}

	rl.renderHelpers()
}

func HkFnJumpForwards(rl *Instance) {
	rl.viUndoSkipAppend = true
	rl.moveCursorByRuneAdjust(rl.viJumpE(tokeniseLine))
}

func HkFnJumpBackwards(rl *Instance) {
	rl.viUndoSkipAppend = true
	rl.moveCursorByRuneAdjust(rl.viJumpB(tokeniseLine))
}

func HkFnCancelAction(rl *Instance) {
	rl.viUndoSkipAppend = true
	switch {
	case rl.modeAutoFind:
		rl.clearPreview()
		rl.resetTabFind()
		rl.clearHelpers()
		rl.resetTabCompletion()
		rl.renderHelpers()

	case rl.modeTabFind:
		rl.resetTabFind()

	case rl.modeTabCompletion:
		rl.clearPreview()
		rl.clearHelpers()
		rl.resetTabCompletion()
		rl.renderHelpers()

	default:
		if rl.line.RunePos() == rl.line.RuneLen() && rl.line.RuneLen() > 0 {
			rl.line.SetRunePos(rl.line.RunePos() - 1)
			moveCursorBackwards(1)
		}
		rl.modeViMode = vimKeys
		rl.viIteration = ""
		rl.viHintMessage()
	}
}

func HkFnRecallWord1(rl *Instance)  { hkFnRecallWord(rl, 1) }
func HkFnRecallWord2(rl *Instance)  { hkFnRecallWord(rl, 2) }
func HkFnRecallWord3(rl *Instance)  { hkFnRecallWord(rl, 3) }
func HkFnRecallWord4(rl *Instance)  { hkFnRecallWord(rl, 4) }
func HkFnRecallWord5(rl *Instance)  { hkFnRecallWord(rl, 5) }
func HkFnRecallWord6(rl *Instance)  { hkFnRecallWord(rl, 6) }
func HkFnRecallWord7(rl *Instance)  { hkFnRecallWord(rl, 7) }
func HkFnRecallWord8(rl *Instance)  { hkFnRecallWord(rl, 8) }
func HkFnRecallWord9(rl *Instance)  { hkFnRecallWord(rl, 9) }
func HkFnRecallWord10(rl *Instance) { hkFnRecallWord(rl, 10) }
func HkFnRecallWord11(rl *Instance) { hkFnRecallWord(rl, 11) }
func HkFnRecallWord12(rl *Instance) { hkFnRecallWord(rl, 12) }

const errCannotRecallWord = "Cannot recall word"

func hkFnRecallWord(rl *Instance, i int) {
	line, err := rl.History.GetLine(rl.History.Len() - 1)
	if err != nil {
		rl.ForceHintTextUpdate(fmt.Sprintf("%s %d: empty history", errCannotRecallWord, i))
		return
	}

	tokens, _, _ := tokeniseSplitSpaces([]rune(line), 0)
	if i > len(tokens) {
		rl.ForceHintTextUpdate(fmt.Sprintf("%s %d: previous line contained fewer words", errCannotRecallWord, i))
		return
	}

	rl.insert([]rune(tokens[i-1] + " "))
}

func HkFnPreviewToggle(rl *Instance) {
	rl.viUndoSkipAppend = true

	switch rl.previewMode {
	case previewModeClosed:
		print(seqSaveBuffer)
		rl.previewMode++
	case previewModeOpen:
		rl.previewMode = previewModeClosed
		print(seqRestoreBuffer)
	case previewModeAutocomplete:
		if rl.modeTabFind {
			rl.resetTabFind()
		}
		HkFnCancelAction(rl)
	}

	rl.echo()
	rl.renderHelpers()
}

func HkFnUndo(rl *Instance) {
	rl.viUndoSkipAppend = true
	if len(rl.viUndoHistory) == 0 {
		return
	}
	rl.undoLast()
	rl.viUndoSkipAppend = true
	rl.line.SetRunePos(rl.line.RuneLen())
	moveCursorForwards(1)
}
