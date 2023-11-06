package app

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/gookit/goutil/arrutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/thomasfinstad/ctl9/internal/pkg/t9"
)

// TODO fix assets path loading

//go:embed assets
var builtInAssets embed.FS

func controllerConnectionStatusHandler(kb *Application) {
	for _, newControllerID := range inpututil.AppendJustConnectedGamepadIDs([]ebiten.GamepadID{}) {
		log.Printf("New controller detected: [id: %d] [SDL GUID: %s]", newControllerID, ebiten.GamepadSDLID(newControllerID))
		kb.controllerIDList = append(kb.controllerIDList, newControllerID)
	}

	for _, controllerID := range kb.controllerIDList {
		if inpututil.IsGamepadJustDisconnected(controllerID) {
			log.Printf("Controller disconnection detected: [id: %d]", controllerID)
			kb.controllerIDList = arrutil.Remove[ebiten.GamepadID](kb.controllerIDList, controllerID)

			if kb.controllerIDActive == &controllerID {
				kb.controllerIDActive = nil
			}
		}
	}

	if kb.controllerIDActive == nil && len(kb.controllerIDList) > 0 {
		kb.controllerIDActive = &kb.controllerIDList[0]

		assestsDir, err := fs.Sub(builtInAssets, "assets")
		if err != nil {
			log.Fatal("unable to open assets directory", err)
		}

		kb.t9 = t9.NewT9(*kb.controllerIDActive, assestsDir)
	}
}

/*
*******************

Informational helpers

*******************
*/

var standardControllerButtonStr = map[ebiten.StandardGamepadButton]string{
	// Backside Left
	ebiten.StandardGamepadButtonFrontTopLeft:    "FTL",
	ebiten.StandardGamepadButtonFrontBottomLeft: "FBL",

	// Backside Right
	ebiten.StandardGamepadButtonFrontTopRight:    "FTR",
	ebiten.StandardGamepadButtonFrontBottomRight: "FBR",

	// D-Pad
	ebiten.StandardGamepadButtonLeftTop:    "LT",
	ebiten.StandardGamepadButtonLeftRight:  "LR",
	ebiten.StandardGamepadButtonLeftBottom: "LB",
	ebiten.StandardGamepadButtonLeftLeft:   "LL",

	// Action keys
	ebiten.StandardGamepadButtonRightTop:    "RT",
	ebiten.StandardGamepadButtonRightRight:  "RR",
	ebiten.StandardGamepadButtonRightBottom: "RB",
	ebiten.StandardGamepadButtonRightLeft:   "RL",

	// Center / menu keys
	ebiten.StandardGamepadButtonCenterCenter: "CC",
	ebiten.StandardGamepadButtonCenterLeft:   "CL",
	ebiten.StandardGamepadButtonCenterRight:  "CR",

	// Stick press
	ebiten.StandardGamepadButtonLeftStick:  "LS",
	ebiten.StandardGamepadButtonRightStick: "RS",
}

var standardControllerAxisStr = map[ebiten.StandardGamepadAxis]string{
	// Left
	ebiten.StandardGamepadAxisLeftStickHorizontal: "LX",
	ebiten.StandardGamepadAxisLeftStickVertical:   "LY",

	// Right
	ebiten.StandardGamepadAxisRightStickHorizontal: "RX",
	ebiten.StandardGamepadAxisRightStickVertical:   "RY",
}

func normalControllerInfoStr(controllerID ebiten.GamepadID) string {

	controllerInfo := fmt.Sprintf(`--------------------------------
Ebiten ID: %d
SDL ID: %s
Name: %s
`,
		controllerID,
		ebiten.GamepadSDLID(controllerID),
		ebiten.GamepadName(controllerID))

	controllerLayout := `
    < FTL >                     < FTR >
    < FBL >                     < FBR >

    < LT >        < CC >        < RT >
< LL >  < LR > < CL >< CR > < RL > < RR >
    < LB >                    < RB >

    < LS >                    < RS >
LXNNN / LYNNN             RXNNN / RYNNN

`

	for button, buttonName := range standardControllerButtonStr {
		layoutStr := fmt.Sprintf("< %s >", buttonName)
		switch {
		case !ebiten.IsStandardGamepadButtonAvailable(controllerID, button):
			controllerLayout = strings.Replace(controllerLayout, layoutStr, "      ", 1)
		case ebiten.IsStandardGamepadButtonPressed(controllerID, button):
			controllerLayout = strings.Replace(controllerLayout, layoutStr, fmt.Sprintf("< %s >", buttonName), 1)
		default:
			controllerLayout = strings.Replace(controllerLayout, layoutStr, fmt.Sprintf("  %s  ", buttonName), 1)
		}
	}

	for axis, axisName := range standardControllerAxisStr {
		layoutStr := fmt.Sprintf("%sNNN", axisName)
		switch {
		case !ebiten.IsStandardGamepadAxisAvailable(controllerID, axis):
			controllerLayout = strings.Replace(controllerLayout, layoutStr, "     ", 1)
		default:

			controllerLayout = strings.Replace(
				controllerLayout,
				layoutStr,
				fmt.Sprintf("%+0.2f", ebiten.StandardGamepadAxisValue(controllerID, axis)),
				1)
		}
	}

	allAxesInfo := "\nAll axes: "
	for axisIndex := 0; axisIndex < ebiten.GamepadAxisCount(controllerID); axisIndex++ {
		allAxesInfo += fmt.Sprintf(
			" %d: %+0.2f |",
			axisIndex,
			ebiten.GamepadAxisValue(controllerID, axisIndex))
	}
	pressedSgp := "\nPressed standard gamepad:"
	for _, b := range inpututil.AppendPressedStandardGamepadButtons(controllerID, []ebiten.StandardGamepadButton{}) {
		pressedSgp += fmt.Sprintf(" %d", b)
	}

	pressedGp := "\nPressed gamepad:"
	for _, b := range inpututil.AppendPressedGamepadButtons(controllerID, []ebiten.GamepadButton{}) {
		pressedGp += fmt.Sprintf(" %d", b)
	}

	justPressedSgp := "\nJust pressed standard gamepad:"
	for _, b := range inpututil.AppendJustPressedStandardGamepadButtons(controllerID, []ebiten.StandardGamepadButton{}) {
		justPressedSgp += fmt.Sprintf(" %d", b)
	}

	justPressedGp := "\nJust pressed gamepad:"
	for _, b := range inpututil.AppendJustPressedGamepadButtons(controllerID, []ebiten.GamepadButton{}) {
		justPressedGp += fmt.Sprintf(" %d", b)
	}

	return controllerInfo +
		controllerLayout +
		allAxesInfo +
		pressedSgp +
		pressedGp +
		justPressedSgp +
		justPressedGp
}
