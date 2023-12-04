// Code generated by mockery v2.25.1. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// Broadcaster is an autogenerated mock type for the Broadcaster type
type Broadcaster struct {
	mock.Mock
}

// OnAdd provides a mock function with given fields: _a0, _a1
func (_m *Broadcaster) OnAdd(_a0 interface{}, _a1 bool) {
	_m.Called(_a0, _a1)
}

// OnDelete provides a mock function with given fields: _a0
func (_m *Broadcaster) OnDelete(_a0 interface{}) {
	_m.Called(_a0)
}

// OnUpdate provides a mock function with given fields: _a0, _a1
func (_m *Broadcaster) OnUpdate(_a0 interface{}, _a1 interface{}) {
	_m.Called(_a0, _a1)
}

// Subscribe provides a mock function with given fields: ch, filters
func (_m *Broadcaster) Subscribe(ch chan *v1alpha1.ApplicationWatchEvent, filters ...func(*v1alpha1.ApplicationWatchEvent) bool) func() {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ch)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 func()
	if rf, ok := ret.Get(0).(func(chan *v1alpha1.ApplicationWatchEvent, ...func(*v1alpha1.ApplicationWatchEvent) bool) func()); ok {
		r0 = rf(ch, filters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}

type mockConstructorTestingTNewBroadcaster interface {
	mock.TestingT
	Cleanup(func())
}

// NewBroadcaster creates a new instance of Broadcaster. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBroadcaster(t mockConstructorTestingTNewBroadcaster) *Broadcaster {
	mock := &Broadcaster{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
