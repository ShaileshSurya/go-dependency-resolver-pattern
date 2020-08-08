package externalfactory

import "github.com/ShaileshSurya/go-dependency-resolver-pattern/external"

// DependencyResolver ...
type DependencyResolver interface {
	DependencyFunction() error
}

// GetResolver ...
func GetResolver(env string) DependencyResolver {
	if env == "test" {
		return &mockResolver{}
	}
	return &resolver{}
}

// EmptyStruct ...
type resolver struct {
}

// DependencyFunction ...
func (r *resolver) DependencyFunction() error {
	return external.DependencyFunction()
}
