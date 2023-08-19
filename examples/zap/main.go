package main

import (
	"fmt"

	"go.uber.org/zap"
)

type SomeInterface interface {
	SetGreeting(greeting string) (err error)
	Run(name string) (err error)
}

var _ SomeInterface = (*SomeStruct)(nil)

type SomeStruct struct {
	Greeting string
}

func (s *SomeStruct) Run(name string) error {
	if name == "" {
		return fmt.Errorf("empty name")
	}

	fmt.Println(s.Greeting + ", " + name + "!")

	return nil
}

func (s *SomeStruct) SetGreeting(greeting string) error {
	if greeting == "" {
		return fmt.Errorf("empty greeting")
	}

	s.Greeting = greeting

	return nil
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	someApp := NewSomeInterfaceWithZap(&SomeStruct{}, logger)

	// happy run
	_ = someApp.SetGreeting("Hello")
	_ = someApp.Run("Man")

	// empty greeting
	_ = someApp.SetGreeting("")
	// empty name
	_ = someApp.Run("")
}
