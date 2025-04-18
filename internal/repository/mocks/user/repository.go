// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	models "avito-pvz/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *Repository) Create(ctx context.Context, _a1 models.User) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Repository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 models.User
func (_e *Repository_Expecter) Create(ctx interface{}, _a1 interface{}) *Repository_Create_Call {
	return &Repository_Create_Call{Call: _e.mock.On("Create", ctx, _a1)}
}

func (_c *Repository_Create_Call) Run(run func(ctx context.Context, _a1 models.User)) *Repository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.User))
	})
	return _c
}

func (_c *Repository_Create_Call) Return(_a0 error) *Repository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Create_Call) RunAndReturn(run func(context.Context, models.User) error) *Repository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function with given fields: ctx, email
func (_m *Repository) Exists(ctx context.Context, email string) (bool, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type Repository_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *Repository_Expecter) Exists(ctx interface{}, email interface{}) *Repository_Exists_Call {
	return &Repository_Exists_Call{Call: _e.mock.On("Exists", ctx, email)}
}

func (_c *Repository_Exists_Call) Run(run func(ctx context.Context, email string)) *Repository_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Repository_Exists_Call) Return(_a0 bool, _a1 error) *Repository_Exists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Exists_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *Repository_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) GetByEmail(ctx context.Context, email string) (models.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type Repository_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *Repository_Expecter) GetByEmail(ctx interface{}, email interface{}) *Repository_GetByEmail_Call {
	return &Repository_GetByEmail_Call{Call: _e.mock.On("GetByEmail", ctx, email)}
}

func (_c *Repository_GetByEmail_Call) Run(run func(ctx context.Context, email string)) *Repository_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Repository_GetByEmail_Call) Return(_a0 models.User, _a1 error) *Repository_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetByEmail_Call) RunAndReturn(run func(context.Context, string) (models.User, error)) *Repository_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
