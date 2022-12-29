// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/bangumi/server/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

type Repo_Expecter struct {
	mock *mock.Mock
}

func (_m *Repo) EXPECT() *Repo_Expecter {
	return &Repo_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx, userID
func (_m *Repo) Count(ctx context.Context, userID model.UserID) (int64, error) {
	ret := _m.Called(ctx, userID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID) int64); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repo_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type Repo_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
//   - userID model.UserID
func (_e *Repo_Expecter) Count(ctx interface{}, userID interface{}) *Repo_Count_Call {
	return &Repo_Count_Call{Call: _e.mock.On("Count", ctx, userID)}
}

func (_c *Repo_Count_Call) Run(run func(ctx context.Context, userID model.UserID)) *Repo_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID))
	})
	return _c
}

func (_c *Repo_Count_Call) Return(_a0 int64, _a1 error) *Repo_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepo(t mockConstructorTestingTNewRepo) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}