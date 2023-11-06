package constraint

import "github.com/hajimehoshi/ebiten/v2"

/*
NewAnd creates a new "and" constraint-type,
this will be vaild if ALL the sub-constraints are valid.

Valid constraint structure is a list with constraint maps:
[	{constraint-type : constraint-config}, ...]
*/
func NewAnd(controllerID ebiten.GamepadID, constraints []any) *And {
	c := &And{}
	for _, v := range constraints {
		c.constraints = append(c.constraints, NewConstraint(controllerID, v.(map[string]any)))
	}
	return c
}

// And requires all sub constraints to be valid
type And struct {
	constraints []Constraint
}

// IsValid checks if all sub constraints are valid
func (c *And) IsValid() bool {
	for _, c := range c.constraints {
		if !c.IsValid() {
			return false
		}
	}
	return true
}
