package constraint

import (
	"log"

	"github.com/gookit/goutil/maputil"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
NewConstraint takes in a constraint tree and recursively creates constraints.
To see how to configure each constraint type see doc string from the
relevant New<type>() functions.

Supported constraint types are:
  - and
  - not
  - or
  - xor
  - button
  - button-once
  - key
  - axis
*/
func NewConstraint(controllerID ebiten.GamepadID, config map[string]any) Constraint {
	kinds := maputil.Keys(config)
	if len(kinds) != 1 {
		log.Fatalf("expected exactly one kind of constraint, got %d: %#v", len(kinds), kinds)
	}
	kind := kinds[0]

	switch kind {
	case "and":
		return NewAnd(controllerID, config[kind].([]any))
	case "not":
		return NewNot(controllerID, config[kind].(map[string]any))
	case "or":
		return NewOr(controllerID, config[kind].([]any))
	case "xor":
		return NewXor(controllerID, config[kind].([]any))
	case "button":
		return NewButton(controllerID, config[kind])
	case "button-once":
		return NewButtonOnce(controllerID, config[kind])
	case "key":
		return NewKey(config[kind])
	case "axis":
		return NewAxis(controllerID, config[kind].(map[string]any))
	default:
		log.Fatalf("unknown constraint kind %#v", kind)
		return nil
	}
}
