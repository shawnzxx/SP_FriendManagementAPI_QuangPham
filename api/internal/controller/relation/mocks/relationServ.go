// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/controller/relation"
	mock "github.com/stretchr/testify/mock"
)

// RelationServ is an autogenerated mock type for the RelationServ type
type RelationServ struct {
	mock.Mock
}

// CreateBlockRelation provides a mock function with given fields: ctx, input
func (_m *RelationServ) CreateBlockRelation(ctx context.Context, input relation.CreateRelationsInput) (relation.CreateRelationsResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 relation.CreateRelationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, relation.CreateRelationsInput) relation.CreateRelationsResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(relation.CreateRelationsResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, relation.CreateRelationsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFriendRelation provides a mock function with given fields: ctx, input
func (_m *RelationServ) CreateFriendRelation(ctx context.Context, input relation.CreateRelationsInput) (relation.CreateRelationsResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 relation.CreateRelationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, relation.CreateRelationsInput) relation.CreateRelationsResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(relation.CreateRelationsResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, relation.CreateRelationsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSubscriptionRelation provides a mock function with given fields: ctx, input
func (_m *RelationServ) CreateSubscriptionRelation(ctx context.Context, input relation.CreateRelationsInput) (relation.CreateRelationsResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 relation.CreateRelationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, relation.CreateRelationsInput) relation.CreateRelationsResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(relation.CreateRelationsResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, relation.CreateRelationsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllFriendsOfUser provides a mock function with given fields: ctx, input
func (_m *RelationServ) GetAllFriendsOfUser(ctx context.Context, input relation.GetAllFriendsInput) (relation.FriendListResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 relation.FriendListResponse
	if rf, ok := ret.Get(0).(func(context.Context, relation.GetAllFriendsInput) relation.FriendListResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(relation.FriendListResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, relation.GetAllFriendsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommonFriendList provides a mock function with given fields: ctx, input
func (_m *RelationServ) GetCommonFriendList(ctx context.Context, input relation.CommonFriendsInput) (relation.FriendListResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 relation.FriendListResponse
	if rf, ok := ret.Get(0).(func(context.Context, relation.CommonFriendsInput) relation.FriendListResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(relation.FriendListResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, relation.CommonFriendsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmailReceive provides a mock function with given fields: ctx, input
func (_m *RelationServ) GetEmailReceive(ctx context.Context, input relation.EmailReceiveInput) (relation.EmailReceiveResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 relation.EmailReceiveResponse
	if rf, ok := ret.Get(0).(func(context.Context, relation.EmailReceiveInput) relation.EmailReceiveResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(relation.EmailReceiveResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, relation.EmailReceiveInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRelationServ interface {
	mock.TestingT
	Cleanup(func())
}

// NewRelationServ creates a new instance of RelationServ. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRelationServ(t mockConstructorTestingTNewRelationServ) *RelationServ {
	mock := &RelationServ{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
