package constraint

import "github.com/hajimehoshi/ebiten/v2"

/*
NewOr creates a new "or" constraint-type,
this will be vaild if ANY the sub-constraints are valid.

Valid constraint structure is a list with constraint maps:
[	{constraint-type : constraint-config}, ...]
*/
func NewOr(controllerID ebiten.GamepadID, constraints []any) *Or {
	c := &Or{}
	for _, v := range constraints {
		c.constraints = append(c.constraints, NewConstraint(controllerID, v.(map[string]any)))
	}
	return c
}

// Or checks if at least one of the sub constraints are valid
type Or struct {
	constraints []Constraint
}

// IsValid checks if at least one of the sub constraints are valid
func (c *Or) IsValid() bool {
	for _, c := range c.constraints {
		if c.IsValid() {
			return true
		}
	}
	return false
}
