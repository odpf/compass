// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	asset "github.com/odpf/compass/core/asset"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// AssetService is an autogenerated mock type for the AssetService type
type AssetService struct {
	mock.Mock
}

type AssetService_Expecter struct {
	mock *mock.Mock
}

func (_m *AssetService) EXPECT() *AssetService_Expecter {
	return &AssetService_Expecter{mock: &_m.Mock}
}

// DeleteAsset provides a mock function with given fields: _a0, _a1
func (_m *AssetService) DeleteAsset(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetService_DeleteAsset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAsset'
type AssetService_DeleteAsset_Call struct {
	*mock.Call
}

// DeleteAsset is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 string
func (_e *AssetService_Expecter) DeleteAsset(_a0 interface{}, _a1 interface{}) *AssetService_DeleteAsset_Call {
	return &AssetService_DeleteAsset_Call{Call: _e.mock.On("DeleteAsset", _a0, _a1)}
}

func (_c *AssetService_DeleteAsset_Call) Run(run func(_a0 context.Context, _a1 string)) *AssetService_DeleteAsset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AssetService_DeleteAsset_Call) Return(_a0 error) *AssetService_DeleteAsset_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetAllAssets provides a mock function with given fields: _a0, _a1, _a2
func (_m *AssetService) GetAllAssets(_a0 context.Context, _a1 asset.Filter, _a2 bool) ([]asset.Asset, uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []asset.Asset
	if rf, ok := ret.Get(0).(func(context.Context, asset.Filter, bool) []asset.Asset); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]asset.Asset)
		}
	}

	var r1 uint32
	if rf, ok := ret.Get(1).(func(context.Context, asset.Filter, bool) uint32); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, asset.Filter, bool) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AssetService_GetAllAssets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllAssets'
type AssetService_GetAllAssets_Call struct {
	*mock.Call
}

// GetAllAssets is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 asset.Filter
//  - _a2 bool
func (_e *AssetService_Expecter) GetAllAssets(_a0 interface{}, _a1 interface{}, _a2 interface{}) *AssetService_GetAllAssets_Call {
	return &AssetService_GetAllAssets_Call{Call: _e.mock.On("GetAllAssets", _a0, _a1, _a2)}
}

func (_c *AssetService_GetAllAssets_Call) Run(run func(_a0 context.Context, _a1 asset.Filter, _a2 bool)) *AssetService_GetAllAssets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(asset.Filter), args[2].(bool))
	})
	return _c
}

func (_c *AssetService_GetAllAssets_Call) Return(_a0 []asset.Asset, _a1 uint32, _a2 error) *AssetService_GetAllAssets_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

// GetAssetByID provides a mock function with given fields: ctx, id
func (_m *AssetService) GetAssetByID(ctx context.Context, id string) (asset.Asset, error) {
	ret := _m.Called(ctx, id)

	var r0 asset.Asset
	if rf, ok := ret.Get(0).(func(context.Context, string) asset.Asset); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(asset.Asset)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_GetAssetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAssetByID'
type AssetService_GetAssetByID_Call struct {
	*mock.Call
}

// GetAssetByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id string
func (_e *AssetService_Expecter) GetAssetByID(ctx interface{}, id interface{}) *AssetService_GetAssetByID_Call {
	return &AssetService_GetAssetByID_Call{Call: _e.mock.On("GetAssetByID", ctx, id)}
}

func (_c *AssetService_GetAssetByID_Call) Run(run func(ctx context.Context, id string)) *AssetService_GetAssetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AssetService_GetAssetByID_Call) Return(_a0 asset.Asset, _a1 error) *AssetService_GetAssetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAssetByURN provides a mock function with given fields: ctx, urn, typ, service
func (_m *AssetService) GetAssetByURN(ctx context.Context, urn string, typ asset.Type, service string) (asset.Asset, error) {
	ret := _m.Called(ctx, urn, typ, service)

	var r0 asset.Asset
	if rf, ok := ret.Get(0).(func(context.Context, string, asset.Type, string) asset.Asset); ok {
		r0 = rf(ctx, urn, typ, service)
	} else {
		r0 = ret.Get(0).(asset.Asset)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, asset.Type, string) error); ok {
		r1 = rf(ctx, urn, typ, service)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_GetAssetByURN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAssetByURN'
type AssetService_GetAssetByURN_Call struct {
	*mock.Call
}

// GetAssetByURN is a helper method to define mock.On call
//  - ctx context.Context
//  - urn string
//  - typ asset.Type
//  - service string
func (_e *AssetService_Expecter) GetAssetByURN(ctx interface{}, urn interface{}, typ interface{}, service interface{}) *AssetService_GetAssetByURN_Call {
	return &AssetService_GetAssetByURN_Call{Call: _e.mock.On("GetAssetByURN", ctx, urn, typ, service)}
}

func (_c *AssetService_GetAssetByURN_Call) Run(run func(ctx context.Context, urn string, typ asset.Type, service string)) *AssetService_GetAssetByURN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(asset.Type), args[3].(string))
	})
	return _c
}

func (_c *AssetService_GetAssetByURN_Call) Return(_a0 asset.Asset, _a1 error) *AssetService_GetAssetByURN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAssetByVersion provides a mock function with given fields: ctx, id, version
func (_m *AssetService) GetAssetByVersion(ctx context.Context, id string, version string) (asset.Asset, error) {
	ret := _m.Called(ctx, id, version)

	var r0 asset.Asset
	if rf, ok := ret.Get(0).(func(context.Context, string, string) asset.Asset); ok {
		r0 = rf(ctx, id, version)
	} else {
		r0 = ret.Get(0).(asset.Asset)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_GetAssetByVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAssetByVersion'
type AssetService_GetAssetByVersion_Call struct {
	*mock.Call
}

// GetAssetByVersion is a helper method to define mock.On call
//  - ctx context.Context
//  - id string
//  - version string
func (_e *AssetService_Expecter) GetAssetByVersion(ctx interface{}, id interface{}, version interface{}) *AssetService_GetAssetByVersion_Call {
	return &AssetService_GetAssetByVersion_Call{Call: _e.mock.On("GetAssetByVersion", ctx, id, version)}
}

func (_c *AssetService_GetAssetByVersion_Call) Run(run func(ctx context.Context, id string, version string)) *AssetService_GetAssetByVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AssetService_GetAssetByVersion_Call) Return(_a0 asset.Asset, _a1 error) *AssetService_GetAssetByVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAssetVersionHistory provides a mock function with given fields: ctx, flt, id
func (_m *AssetService) GetAssetVersionHistory(ctx context.Context, flt asset.Filter, id string) ([]asset.Asset, error) {
	ret := _m.Called(ctx, flt, id)

	var r0 []asset.Asset
	if rf, ok := ret.Get(0).(func(context.Context, asset.Filter, string) []asset.Asset); ok {
		r0 = rf(ctx, flt, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]asset.Asset)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, asset.Filter, string) error); ok {
		r1 = rf(ctx, flt, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_GetAssetVersionHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAssetVersionHistory'
type AssetService_GetAssetVersionHistory_Call struct {
	*mock.Call
}

// GetAssetVersionHistory is a helper method to define mock.On call
//  - ctx context.Context
//  - flt asset.Filter
//  - id string
func (_e *AssetService_Expecter) GetAssetVersionHistory(ctx interface{}, flt interface{}, id interface{}) *AssetService_GetAssetVersionHistory_Call {
	return &AssetService_GetAssetVersionHistory_Call{Call: _e.mock.On("GetAssetVersionHistory", ctx, flt, id)}
}

func (_c *AssetService_GetAssetVersionHistory_Call) Run(run func(ctx context.Context, flt asset.Filter, id string)) *AssetService_GetAssetVersionHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(asset.Filter), args[2].(string))
	})
	return _c
}

func (_c *AssetService_GetAssetVersionHistory_Call) Return(_a0 []asset.Asset, _a1 error) *AssetService_GetAssetVersionHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLineage provides a mock function with given fields: ctx, node
func (_m *AssetService) GetLineage(ctx context.Context, node asset.LineageNode) (asset.LineageGraph, error) {
	ret := _m.Called(ctx, node)

	var r0 asset.LineageGraph
	if rf, ok := ret.Get(0).(func(context.Context, asset.LineageNode) asset.LineageGraph); ok {
		r0 = rf(ctx, node)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(asset.LineageGraph)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, asset.LineageNode) error); ok {
		r1 = rf(ctx, node)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_GetLineage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLineage'
type AssetService_GetLineage_Call struct {
	*mock.Call
}

// GetLineage is a helper method to define mock.On call
//  - ctx context.Context
//  - node asset.LineageNode
func (_e *AssetService_Expecter) GetLineage(ctx interface{}, node interface{}) *AssetService_GetLineage_Call {
	return &AssetService_GetLineage_Call{Call: _e.mock.On("GetLineage", ctx, node)}
}

func (_c *AssetService_GetLineage_Call) Run(run func(ctx context.Context, node asset.LineageNode)) *AssetService_GetLineage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(asset.LineageNode))
	})
	return _c
}

func (_c *AssetService_GetLineage_Call) Return(_a0 asset.LineageGraph, _a1 error) *AssetService_GetLineage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTypes provides a mock function with given fields: ctx, flt
func (_m *AssetService) GetTypes(ctx context.Context, flt asset.Filter) (map[asset.Type]int, error) {
	ret := _m.Called(ctx, flt)

	var r0 map[asset.Type]int
	if rf, ok := ret.Get(0).(func(context.Context, asset.Filter) map[asset.Type]int); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[asset.Type]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, asset.Filter) error); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_GetTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTypes'
type AssetService_GetTypes_Call struct {
	*mock.Call
}

// GetTypes is a helper method to define mock.On call
//  - ctx context.Context
//  - flt asset.Filter
func (_e *AssetService_Expecter) GetTypes(ctx interface{}, flt interface{}) *AssetService_GetTypes_Call {
	return &AssetService_GetTypes_Call{Call: _e.mock.On("GetTypes", ctx, flt)}
}

func (_c *AssetService_GetTypes_Call) Run(run func(ctx context.Context, flt asset.Filter)) *AssetService_GetTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(asset.Filter))
	})
	return _c
}

func (_c *AssetService_GetTypes_Call) Return(_a0 map[asset.Type]int, _a1 error) *AssetService_GetTypes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SearchAssets provides a mock function with given fields: ctx, cfg
func (_m *AssetService) SearchAssets(ctx context.Context, cfg asset.SearchConfig) ([]asset.SearchResult, error) {
	ret := _m.Called(ctx, cfg)

	var r0 []asset.SearchResult
	if rf, ok := ret.Get(0).(func(context.Context, asset.SearchConfig) []asset.SearchResult); ok {
		r0 = rf(ctx, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]asset.SearchResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, asset.SearchConfig) error); ok {
		r1 = rf(ctx, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_SearchAssets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchAssets'
type AssetService_SearchAssets_Call struct {
	*mock.Call
}

// SearchAssets is a helper method to define mock.On call
//  - ctx context.Context
//  - cfg asset.SearchConfig
func (_e *AssetService_Expecter) SearchAssets(ctx interface{}, cfg interface{}) *AssetService_SearchAssets_Call {
	return &AssetService_SearchAssets_Call{Call: _e.mock.On("SearchAssets", ctx, cfg)}
}

func (_c *AssetService_SearchAssets_Call) Run(run func(ctx context.Context, cfg asset.SearchConfig)) *AssetService_SearchAssets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(asset.SearchConfig))
	})
	return _c
}

func (_c *AssetService_SearchAssets_Call) Return(results []asset.SearchResult, err error) *AssetService_SearchAssets_Call {
	_c.Call.Return(results, err)
	return _c
}

// SuggestAssets provides a mock function with given fields: ctx, cfg
func (_m *AssetService) SuggestAssets(ctx context.Context, cfg asset.SearchConfig) ([]string, error) {
	ret := _m.Called(ctx, cfg)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, asset.SearchConfig) []string); ok {
		r0 = rf(ctx, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, asset.SearchConfig) error); ok {
		r1 = rf(ctx, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_SuggestAssets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuggestAssets'
type AssetService_SuggestAssets_Call struct {
	*mock.Call
}

// SuggestAssets is a helper method to define mock.On call
//  - ctx context.Context
//  - cfg asset.SearchConfig
func (_e *AssetService_Expecter) SuggestAssets(ctx interface{}, cfg interface{}) *AssetService_SuggestAssets_Call {
	return &AssetService_SuggestAssets_Call{Call: _e.mock.On("SuggestAssets", ctx, cfg)}
}

func (_c *AssetService_SuggestAssets_Call) Run(run func(ctx context.Context, cfg asset.SearchConfig)) *AssetService_SuggestAssets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(asset.SearchConfig))
	})
	return _c
}

func (_c *AssetService_SuggestAssets_Call) Return(suggestions []string, err error) *AssetService_SuggestAssets_Call {
	_c.Call.Return(suggestions, err)
	return _c
}

// UpsertPatchAsset provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *AssetService) UpsertPatchAsset(_a0 context.Context, _a1 *asset.Asset, _a2 []asset.LineageNode, _a3 []asset.LineageNode) (string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *asset.Asset, []asset.LineageNode, []asset.LineageNode) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *asset.Asset, []asset.LineageNode, []asset.LineageNode) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetService_UpsertPatchAsset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertPatchAsset'
type AssetService_UpsertPatchAsset_Call struct {
	*mock.Call
}

// UpsertPatchAsset is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *asset.Asset
//  - _a2 []asset.LineageNode
//  - _a3 []asset.LineageNode
func (_e *AssetService_Expecter) UpsertPatchAsset(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *AssetService_UpsertPatchAsset_Call {
	return &AssetService_UpsertPatchAsset_Call{Call: _e.mock.On("UpsertPatchAsset", _a0, _a1, _a2, _a3)}
}

func (_c *AssetService_UpsertPatchAsset_Call) Run(run func(_a0 context.Context, _a1 *asset.Asset, _a2 []asset.LineageNode, _a3 []asset.LineageNode)) *AssetService_UpsertPatchAsset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*asset.Asset), args[2].([]asset.LineageNode), args[3].([]asset.LineageNode))
	})
	return _c
}

func (_c *AssetService_UpsertPatchAsset_Call) Return(_a0 string, _a1 error) *AssetService_UpsertPatchAsset_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// NewAssetService creates a new instance of AssetService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewAssetService(t testing.TB) *AssetService {
	mock := &AssetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
