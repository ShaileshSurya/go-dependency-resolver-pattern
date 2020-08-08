package main

import (
	"flag"
	"fmt"

	"github.com/ShaileshSurya/go-dependency-resolver-pattern/factory"
)

func main() {
	env := flag.String("env", "dev", "Environment Test Or Development")
	flag.Parse()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", *env)
	depResolver := factory.GetResolver(*env)
	err := depResolver.DependencyFunction()
	if err != nil {
		fmt.Println("Error from dependency function")
	}
}
