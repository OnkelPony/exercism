package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	Greet(name string) string
	LanguageName() string
}

func SayHello(name string, greeter Greeter) string {
