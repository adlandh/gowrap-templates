package main

import (
	"fmt"

	"go.uber.org/zap"
)

//go:generate gowrap gen -i SomeInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/zap.gotmpl -o main_gen.go -l "" -g

type SomeInterface[T any] interface {
	SetGreeting(greeting T) (err error)
	Run(name T) (err error)
}

var _ SomeInterface[string] = (*SomeStruct)(nil)

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

	someApp := NewSomeInterfaceWithZap[string](&SomeStruct{}, logger)

	// happy run
	_ = someApp.SetGreeting("Hello")
	_ = someApp.Run("Man")

	// empty greeting
	_ = someApp.SetGreeting("")
	// empty name
	_ = someApp.Run("")
}
