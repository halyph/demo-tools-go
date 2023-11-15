// Code generated by mockery v2.35.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MyFoo is an autogenerated mock type for the MyFoo type
type MyFoo struct {
	mock.Mock
}

// Process provides a mock function with given fields: input
func (_m *MyFoo) Process(input int) int {
	ret := _m.Called(input)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NewMyFoo creates a new instance of MyFoo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMyFoo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MyFoo {
	mock := &MyFoo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}