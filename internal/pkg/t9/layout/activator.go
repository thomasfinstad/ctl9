package layout

import (
	"github.com/thomasfinstad/ctl9/internal/pkg/constraint"
)

type activatorConstraint constraint.Constraint

type activator struct {
	Constraint        activatorConstraint
	RequireConstraint bool
}

func (a *activator) Active() bool {
	if a.Constraint == nil && !a.RequireConstraint {
		return true
	}

	return a.Constraint.IsValid()
}
