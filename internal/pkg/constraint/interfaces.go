package constraint

// Constraint is used to check the state of user inputs
type Constraint interface {
	IsValid() bool
}
