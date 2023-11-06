package app

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
*******************

Informational helpers

*******************
*/

func keyboardInfoStr() string {
	keyboardInfo := "--------------------------------"

	pressedKb := "\nPressed keyboard:"
	for _, k := range inpututil.AppendPressedKeys([]ebiten.Key{}) {
		pressedKb += " " + k.String()
	}

	justPressedKb := "\nJust pressed keyboard:"
	for _, k := range inpututil.AppendJustPressedKeys([]ebiten.Key{}) {
		justPressedKb += " " + k.String()
	}

	justPressedTid := "\nJust pressed touch ID:"
	for _, t := range inpututil.AppendJustPressedTouchIDs([]ebiten.TouchID{}) {
		justPressedTid += fmt.Sprintf(" %d", t)
	}

	return keyboardInfo + pressedKb + justPressedKb + justPressedTid
}
