package factory

import "github.com/ShaileshSurya/go-dependency-resolver-pattern/factory/externalfactory"

// GetResolver ...
func GetResolver(env string) externalfactory.DependencyResolver {
	return externalfactory.GetResolver(env)
}
