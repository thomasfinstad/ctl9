package constraint

import (
	"log"

	"github.com/gookit/goutil/maputil"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
NewAxis creates a new "axis" constraint-type,
this will be vaild if the specified ebiten standard axis is within the specified bounds.

Valid constraint structure is a map with the axis id, minimum and maximum values:

	{
		id: *ebiten.StandardGamepadAxis
		min: float64
		max: float64
	}
*/
func NewAxis(controllerID ebiten.GamepadID, constraint map[string]any) *Axis {
	a := &Axis{
		gamepadID: controllerID,
	}

	if ok, missing := maputil.HasAllKeys(constraint, "id", "min", "max"); !ok {
		log.Fatalf("constraint for axis is missing config key: %#v", missing)
	}

	id := ebiten.StandardGamepadAxis(constraint["id"].(float64))
	a.controllerAxis = &id

	a.min = constraint["min"].(float64)
	a.max = constraint["max"].(float64)

	return a
}

// Axis checks if a given standard game controller axis is between the min and max values
type Axis struct {
	gamepadID ebiten.GamepadID

	controllerAxis *ebiten.StandardGamepadAxis
	//controllerNonStandardAxis *int

	min float64
	max float64
}

// IsValid checks if constraint is fullfilled
func (c *Axis) IsValid() bool {
	switch {
	case c.controllerAxis != nil:
		v := ebiten.GamepadAxisValue(c.gamepadID, int(*c.controllerAxis))
		return c.min <= v && v <= c.max
	default:
		log.Fatalf("unknown axis kind")
		return false
	}
}
