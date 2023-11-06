package constraint

import "github.com/hajimehoshi/ebiten/v2"

/*
NewXor creates a new "xor" constraint-type,
this will be vaild if EXACTLY ONE the sub-constraints are valid.

Valid constraint structure is a list with constraint maps:
[	{constraint-type : constraint-config}, ...]
*/
func NewXor(controllerID ebiten.GamepadID, constraints []any) *Xor {
	c := &Xor{}
	for _, v := range constraints {
		c.constraints = append(c.constraints, NewConstraint(controllerID, v.(map[string]any)))
	}
	return c
}

// Xor checks if exactly one child constraint is valid
type Xor struct {
	constraints []Constraint
}

// IsValid checks if exactly least one sub constraint evaluates to true
func (c *Xor) IsValid() bool {
	passedCheck := false
	for _, c := range c.constraints {
		if c.IsValid() {
			if passedCheck {
				return false
			}
			passedCheck = true
		}
	}
	return passedCheck
}
