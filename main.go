package main

import (
	demo "github.com/saschagrunert/demo"
)

func main() {
	d := demo.New()

	d.Add(sample(), "demo-0", "just an example demo run")

	d.Run()
}

func sample() *demo.Run {
	r := demo.NewRun(
		"Demo Title",
		"Some additional description",
	)

	r.Step(demo.S(
		"A comment...",
	), nil)

	// A single step can consist of a description and a command to be executed
	r.Step(demo.S(
		"Execute 'ps'...",
	), demo.S(
		"ps",
	))

	return r
}
