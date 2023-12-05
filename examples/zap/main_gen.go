// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/adlandh/gowrap-templates/main/zap.gotmpl
// gowrap: http://github.com/hexdigest/gowrap

package main

import (
	"go.uber.org/zap"
)

// SomeInterfaceWithZap implements SomeInterface that is instrumented with zap logger
type SomeInterfaceWithZap[T any] struct {
	_base SomeInterface[T]
	_log  *zap.Logger
}

// NewSomeInterfaceWithZap instruments an implementation of the SomeInterface with simple logging
func NewSomeInterfaceWithZap[T any](base SomeInterface[T], log *zap.Logger) SomeInterfaceWithZap[T] {
	return SomeInterfaceWithZap[T]{
		_base: base,
		_log:  log,
	}
}

// Run implements SomeInterface
func (_d SomeInterfaceWithZap[T]) Run(name T) (err error) {
	_d._log.Debug("SomeInterfaceWithZap: calling Run", zap.Any("params", map[string]interface{}{
		"name": name}))
	defer func() {
		if err != nil {
			_d._log.Warn("SomeInterfaceWithZap: method Run returned an error", zap.Error(err), zap.Any("result", map[string]interface{}{
				"err": err}))
		} else {
			_d._log.Debug("SomeInterfaceWithZap: method Run finished", zap.Any("result", map[string]interface{}{
				"err": err}))
		}
	}()
	return _d._base.Run(name)
}

// SetGreeting implements SomeInterface
func (_d SomeInterfaceWithZap[T]) SetGreeting(greeting T) (err error) {
	_d._log.Debug("SomeInterfaceWithZap: calling SetGreeting", zap.Any("params", map[string]interface{}{
		"greeting": greeting}))
	defer func() {
		if err != nil {
			_d._log.Warn("SomeInterfaceWithZap: method SetGreeting returned an error", zap.Error(err), zap.Any("result", map[string]interface{}{
				"err": err}))
		} else {
			_d._log.Debug("SomeInterfaceWithZap: method SetGreeting finished", zap.Any("result", map[string]interface{}{
				"err": err}))
		}
	}()
	return _d._base.SetGreeting(greeting)
}
