package constraint

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
NewButton creates a new "button" constraint-type,
this will be vaild if the specified ebiten standard button is pressed.

Valid constraint structure is an int with the standard button id:

*ebiten.StandardGamepadButton (int)
*/
func NewButton(controllerID ebiten.GamepadID, constraint any) *Button {
	id := ebiten.StandardGamepadButton(constraint.(float64))
	return &Button{
		gamepadID:        controllerID,
		controllerButton: &id,
	}
}

/*
NewButtonOnce creates a new "button-once" constraint-type,
this will be vaild if the specified ebiten standard button is pressed, but only the first time it checks after being pressed.

Valid constraint structure is an int with the standard button id:

*ebiten.StandardGamepadButton (int)
*/
func NewButtonOnce(controllerID ebiten.GamepadID, constraint any) *Button {
	id := ebiten.StandardGamepadButton(constraint.(float64))
	return &Button{
		gamepadID:            controllerID,
		controllerButtonOnce: &id,
	}
}

func newControllerNonStandardButton(controllerID ebiten.GamepadID, requiredButton *ebiten.GamepadButton) *Button {
	return &Button{
		gamepadID:                   controllerID,
		nonStandardControllerButton: requiredButton,
	}
}

func newMouseButton(requiredButton *ebiten.MouseButton) *Button {
	return &Button{
		mouseButton: requiredButton,
	}
}

/*
NewKey creates a new "key" constraint-type,
this will be vaild if the specified ebiten keyboard key is pressed.

Valid constraint structure is an int with the keyboard key id:

*ebiten.Key
*/
func NewKey(constraint any) *Button {
	id := ebiten.Key(constraint.(float64))
	return &Button{
		keyboardKey: &id,
	}
}

// Button checks if a standard game controller button is pressed
type Button struct {
	gamepadID ebiten.GamepadID

	controllerButton            *ebiten.StandardGamepadButton
	controllerButtonOnce        *ebiten.StandardGamepadButton
	nonStandardControllerButton *ebiten.GamepadButton
	mouseButton                 *ebiten.MouseButton
	keyboardKey                 *ebiten.Key
}

// IsValid checks if the button state is satisfactory
func (c *Button) IsValid() bool {
	switch {
	case c.controllerButton != nil:
		return ebiten.IsStandardGamepadButtonPressed(c.gamepadID, *c.controllerButton)
	case c.controllerButtonOnce != nil:
		return inpututil.IsStandardGamepadButtonJustPressed(c.gamepadID, *c.controllerButtonOnce)
	case c.nonStandardControllerButton != nil:
		return ebiten.IsGamepadButtonPressed(c.gamepadID, *c.nonStandardControllerButton)
	case c.mouseButton != nil:
		return ebiten.IsMouseButtonPressed(*c.mouseButton)
	case c.keyboardKey != nil:
		return ebiten.IsKeyPressed(*c.keyboardKey)
	default:
		log.Fatal(fmt.Errorf("unknown button kind"))
		return false
	}
}
