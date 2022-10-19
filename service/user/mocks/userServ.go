// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	user "github.com/quangpham789/golang-assessment/service/user"
	mock "github.com/stretchr/testify/mock"
)

// UserServ is an autogenerated mock type for the UserServ type
type UserServ struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, input
func (_m *UserServ) CreateUser(ctx context.Context, input user.CreateUserInput) (user.UserResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 user.UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, user.CreateUserInput) user.UserResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(user.UserResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, user.CreateUserInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserServ interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserServ creates a new instance of UserServ. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserServ(t mockConstructorTestingTNewUserServ) *UserServ {
	mock := &UserServ{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
