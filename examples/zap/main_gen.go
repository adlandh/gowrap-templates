package main

// Code generated by gowrap. DO NOT EDIT.
// template: ../../zap.gotmpl
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p github.com/adlandh/gowrap-templates/examples/zap -i SomeInterface -t ../../zap.gotmpl -o main_gen.go -l ""

import (
	"go.uber.org/zap"
)

// SomeInterfaceWithZap implements SomeInterface that is instrumented with zap logger
type SomeInterfaceWithZap struct {
	_log  *zap.Logger
	_base SomeInterface
}

// NewSomeInterfaceWithZap instruments an implementation of the SomeInterface with simple logging
func NewSomeInterfaceWithZap(base SomeInterface, log *zap.Logger) SomeInterfaceWithZap {
	return SomeInterfaceWithZap{
		_base: base,
		_log:  log,
	}
}

// Run implements SomeInterface
func (_d SomeInterfaceWithZap) Run(name string) (err error) {
	_d._log.Debug("SomeInterfaceWithZap: calling Run", zap.Any("params", map[string]interface{}{
		"name": name}))
	defer func() {
		if err != nil {
			_d._log.Error("SomeInterfaceWithZap: method Run returned an error", zap.Error(err), zap.Any("result", map[string]interface{}{
				"err": err}))
		} else {
			_d._log.Debug("SomeInterfaceWithZap: method Run finished", zap.Any("result", map[string]interface{}{
				"err": err}))
		}
	}()
	return _d._base.Run(name)
}

// SetGreeting implements SomeInterface
func (_d SomeInterfaceWithZap) SetGreeting(greeting string) (err error) {
	_d._log.Debug("SomeInterfaceWithZap: calling SetGreeting", zap.Any("params", map[string]interface{}{
		"greeting": greeting}))
	defer func() {
		if err != nil {
			_d._log.Error("SomeInterfaceWithZap: method SetGreeting returned an error", zap.Error(err), zap.Any("result", map[string]interface{}{
				"err": err}))
		} else {
			_d._log.Debug("SomeInterfaceWithZap: method SetGreeting finished", zap.Any("result", map[string]interface{}{
				"err": err}))
		}
	}()
	return _d._base.SetGreeting(greeting)
}
