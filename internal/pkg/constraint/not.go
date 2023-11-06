package constraint

import "github.com/hajimehoshi/ebiten/v2"

/*
NewNot creates a new "not" constraint-type,
this will only be vaild if the sub-constraint is invalid.

Valid constraint structure is a constraint map:
{constraint-type : constraint-config}
*/
func NewNot(controllerID ebiten.GamepadID, constraint map[string]any) *Not {
	return &Not{
		constraint: NewConstraint(controllerID, constraint),
	}
}

// Not reverses the sub constraints value
type Not struct {
	constraint Constraint
}

// IsValid returns the oposite of the child constraint validity
func (c *Not) IsValid() bool {
	return !c.constraint.IsValid()
}
