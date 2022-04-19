// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	asset "github.com/odpf/compass/asset"

	mock "github.com/stretchr/testify/mock"
)

// DiscoveryRepository is an autogenerated mock type for the Repository type
type DiscoveryRepository struct {
	mock.Mock
}

type DiscoveryRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *DiscoveryRepository) EXPECT() *DiscoveryRepository_Expecter {
	return &DiscoveryRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, assetID
func (_m *DiscoveryRepository) Delete(ctx context.Context, assetID string) error {
	ret := _m.Called(ctx, assetID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, assetID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscoveryRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type DiscoveryRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - ctx context.Context
//  - assetID string
func (_e *DiscoveryRepository_Expecter) Delete(ctx interface{}, assetID interface{}) *DiscoveryRepository_Delete_Call {
	return &DiscoveryRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, assetID)}
}

func (_c *DiscoveryRepository_Delete_Call) Run(run func(ctx context.Context, assetID string)) *DiscoveryRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DiscoveryRepository_Delete_Call) Return(_a0 error) *DiscoveryRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *DiscoveryRepository) Upsert(_a0 context.Context, _a1 asset.Asset) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, asset.Asset) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscoveryRepository_Upsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upsert'
type DiscoveryRepository_Upsert_Call struct {
	*mock.Call
}

// Upsert is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 asset.Asset
func (_e *DiscoveryRepository_Expecter) Upsert(_a0 interface{}, _a1 interface{}) *DiscoveryRepository_Upsert_Call {
	return &DiscoveryRepository_Upsert_Call{Call: _e.mock.On("Upsert", _a0, _a1)}
}

func (_c *DiscoveryRepository_Upsert_Call) Run(run func(_a0 context.Context, _a1 asset.Asset)) *DiscoveryRepository_Upsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(asset.Asset))
	})
	return _c
}

func (_c *DiscoveryRepository_Upsert_Call) Return(_a0 error) *DiscoveryRepository_Upsert_Call {
	_c.Call.Return(_a0)
	return _c
}
