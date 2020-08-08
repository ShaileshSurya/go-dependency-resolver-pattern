package externalfactory

import "errors"

// EmptyStruct ...
type mockResolver struct {
}

// DependencyFunction ...
func (mr *mockResolver) DependencyFunction() error {
	return errors.New("Mock Resolver returning error")
}
