// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// RedisCache is an autogenerated mock type for the RedisCache type
type RedisCache struct {
	mock.Mock
}

type RedisCache_Expecter struct {
	mock *mock.Mock
}

func (_m *RedisCache) EXPECT() *RedisCache_Expecter {
	return &RedisCache_Expecter{mock: &_m.Mock}
}

// Del provides a mock function with given fields: ctx, keys
func (_m *RedisCache) Del(ctx context.Context, keys ...string) error {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) error); ok {
		r0 = rf(ctx, keys...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedisCache_Del_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Del'
type RedisCache_Del_Call struct {
	*mock.Call
}

// Del is a helper method to define mock.On call
//   - ctx context.Context
//   - keys ...string
func (_e *RedisCache_Expecter) Del(ctx interface{}, keys ...interface{}) *RedisCache_Del_Call {
	return &RedisCache_Del_Call{Call: _e.mock.On("Del",
		append([]interface{}{ctx}, keys...)...)}
}

func (_c *RedisCache_Del_Call) Run(run func(ctx context.Context, keys ...string)) *RedisCache_Del_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *RedisCache_Del_Call) Return(_a0 error) *RedisCache_Del_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RedisCache_Del_Call) RunAndReturn(run func(context.Context, ...string) error) *RedisCache_Del_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key, value
func (_m *RedisCache) Get(ctx context.Context, key string, value interface{}) (bool, error) {
	ret := _m.Called(ctx, key, value)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) (bool, error)); ok {
		return rf(ctx, key, value)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RedisCache_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RedisCache_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
func (_e *RedisCache_Expecter) Get(ctx interface{}, key interface{}, value interface{}) *RedisCache_Get_Call {
	return &RedisCache_Get_Call{Call: _e.mock.On("Get", ctx, key, value)}
}

func (_c *RedisCache_Get_Call) Run(run func(ctx context.Context, key string, value interface{})) *RedisCache_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *RedisCache_Get_Call) Return(_a0 bool, _a1 error) *RedisCache_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RedisCache_Get_Call) RunAndReturn(run func(context.Context, string, interface{}) (bool, error)) *RedisCache_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, value, ttl
func (_m *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	ret := _m.Called(ctx, key, value, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedisCache_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type RedisCache_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
//   - ttl time.Duration
func (_e *RedisCache_Expecter) Set(ctx interface{}, key interface{}, value interface{}, ttl interface{}) *RedisCache_Set_Call {
	return &RedisCache_Set_Call{Call: _e.mock.On("Set", ctx, key, value, ttl)}
}

func (_c *RedisCache_Set_Call) Run(run func(ctx context.Context, key string, value interface{}, ttl time.Duration)) *RedisCache_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *RedisCache_Set_Call) Return(_a0 error) *RedisCache_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RedisCache_Set_Call) RunAndReturn(run func(context.Context, string, interface{}, time.Duration) error) *RedisCache_Set_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRedisCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewRedisCache creates a new instance of RedisCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRedisCache(t mockConstructorTestingTNewRedisCache) *RedisCache {
	mock := &RedisCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
