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
func (_m *Repository) Create(ctx context.Context, _a1 models.PVZ) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PVZ) error); ok {
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
//   - _a1 models.PVZ
func (_e *Repository_Expecter) Create(ctx interface{}, _a1 interface{}) *Repository_Create_Call {
	return &Repository_Create_Call{Call: _e.mock.On("Create", ctx, _a1)}
}

func (_c *Repository_Create_Call) Run(run func(ctx context.Context, _a1 models.PVZ)) *Repository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.PVZ))
	})
	return _c
}

func (_c *Repository_Create_Call) Return(_a0 error) *Repository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Create_Call) RunAndReturn(run func(context.Context, models.PVZ) error) *Repository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]models.PVZ, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.PVZ
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.PVZ, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.PVZ); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PVZ)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type Repository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Repository_Expecter) GetAll(ctx interface{}) *Repository_GetAll_Call {
	return &Repository_GetAll_Call{Call: _e.mock.On("GetAll", ctx)}
}

func (_c *Repository_GetAll_Call) Run(run func(ctx context.Context)) *Repository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Repository_GetAll_Call) Return(_a0 []models.PVZ, _a1 error) *Repository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetAll_Call) RunAndReturn(run func(context.Context) ([]models.PVZ, error)) *Repository_GetAll_Call {
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
