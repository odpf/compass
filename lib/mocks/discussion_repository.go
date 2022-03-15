// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	discussion "github.com/odpf/columbus/discussion"
	mock "github.com/stretchr/testify/mock"
)

// DiscussionRepository is an autogenerated mock type for the Repository type
type DiscussionRepository struct {
	mock.Mock
}

type DiscussionRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *DiscussionRepository) EXPECT() *DiscussionRepository_Expecter {
	return &DiscussionRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *DiscussionRepository) Create(ctx context.Context, _a1 *discussion.Discussion) (string, error) {
	ret := _m.Called(ctx, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *discussion.Discussion) string); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *discussion.Discussion) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiscussionRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type DiscussionRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 *discussion.Discussion
func (_e *DiscussionRepository_Expecter) Create(ctx interface{}, _a1 interface{}) *DiscussionRepository_Create_Call {
	return &DiscussionRepository_Create_Call{Call: _e.mock.On("Create", ctx, _a1)}
}

func (_c *DiscussionRepository_Create_Call) Run(run func(ctx context.Context, _a1 *discussion.Discussion)) *DiscussionRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discussion.Discussion))
	})
	return _c
}

func (_c *DiscussionRepository_Create_Call) Return(_a0 string, _a1 error) *DiscussionRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateComment provides a mock function with given fields: ctx, cmt
func (_m *DiscussionRepository) CreateComment(ctx context.Context, cmt *discussion.Comment) (string, error) {
	ret := _m.Called(ctx, cmt)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *discussion.Comment) string); ok {
		r0 = rf(ctx, cmt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *discussion.Comment) error); ok {
		r1 = rf(ctx, cmt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiscussionRepository_CreateComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateComment'
type DiscussionRepository_CreateComment_Call struct {
	*mock.Call
}

// CreateComment is a helper method to define mock.On call
//  - ctx context.Context
//  - cmt *discussion.Comment
func (_e *DiscussionRepository_Expecter) CreateComment(ctx interface{}, cmt interface{}) *DiscussionRepository_CreateComment_Call {
	return &DiscussionRepository_CreateComment_Call{Call: _e.mock.On("CreateComment", ctx, cmt)}
}

func (_c *DiscussionRepository_CreateComment_Call) Run(run func(ctx context.Context, cmt *discussion.Comment)) *DiscussionRepository_CreateComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discussion.Comment))
	})
	return _c
}

func (_c *DiscussionRepository_CreateComment_Call) Return(_a0 string, _a1 error) *DiscussionRepository_CreateComment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteComment provides a mock function with given fields: ctx, commentID, discussionID
func (_m *DiscussionRepository) DeleteComment(ctx context.Context, commentID string, discussionID string) error {
	ret := _m.Called(ctx, commentID, discussionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, commentID, discussionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscussionRepository_DeleteComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteComment'
type DiscussionRepository_DeleteComment_Call struct {
	*mock.Call
}

// DeleteComment is a helper method to define mock.On call
//  - ctx context.Context
//  - commentID string
//  - discussionID string
func (_e *DiscussionRepository_Expecter) DeleteComment(ctx interface{}, commentID interface{}, discussionID interface{}) *DiscussionRepository_DeleteComment_Call {
	return &DiscussionRepository_DeleteComment_Call{Call: _e.mock.On("DeleteComment", ctx, commentID, discussionID)}
}

func (_c *DiscussionRepository_DeleteComment_Call) Run(run func(ctx context.Context, commentID string, discussionID string)) *DiscussionRepository_DeleteComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *DiscussionRepository_DeleteComment_Call) Return(_a0 error) *DiscussionRepository_DeleteComment_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: ctx, did
func (_m *DiscussionRepository) Get(ctx context.Context, did string) (discussion.Discussion, error) {
	ret := _m.Called(ctx, did)

	var r0 discussion.Discussion
	if rf, ok := ret.Get(0).(func(context.Context, string) discussion.Discussion); ok {
		r0 = rf(ctx, did)
	} else {
		r0 = ret.Get(0).(discussion.Discussion)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, did)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiscussionRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type DiscussionRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - did string
func (_e *DiscussionRepository_Expecter) Get(ctx interface{}, did interface{}) *DiscussionRepository_Get_Call {
	return &DiscussionRepository_Get_Call{Call: _e.mock.On("Get", ctx, did)}
}

func (_c *DiscussionRepository_Get_Call) Run(run func(ctx context.Context, did string)) *DiscussionRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DiscussionRepository_Get_Call) Return(_a0 discussion.Discussion, _a1 error) *DiscussionRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAll provides a mock function with given fields: ctx, filter
func (_m *DiscussionRepository) GetAll(ctx context.Context, filter discussion.Filter) ([]discussion.Discussion, error) {
	ret := _m.Called(ctx, filter)

	var r0 []discussion.Discussion
	if rf, ok := ret.Get(0).(func(context.Context, discussion.Filter) []discussion.Discussion); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]discussion.Discussion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, discussion.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiscussionRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type DiscussionRepository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//  - ctx context.Context
//  - filter discussion.Filter
func (_e *DiscussionRepository_Expecter) GetAll(ctx interface{}, filter interface{}) *DiscussionRepository_GetAll_Call {
	return &DiscussionRepository_GetAll_Call{Call: _e.mock.On("GetAll", ctx, filter)}
}

func (_c *DiscussionRepository_GetAll_Call) Run(run func(ctx context.Context, filter discussion.Filter)) *DiscussionRepository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(discussion.Filter))
	})
	return _c
}

func (_c *DiscussionRepository_GetAll_Call) Return(_a0 []discussion.Discussion, _a1 error) *DiscussionRepository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAllComments provides a mock function with given fields: ctx, discussionID, flt
func (_m *DiscussionRepository) GetAllComments(ctx context.Context, discussionID string, flt discussion.Filter) ([]discussion.Comment, error) {
	ret := _m.Called(ctx, discussionID, flt)

	var r0 []discussion.Comment
	if rf, ok := ret.Get(0).(func(context.Context, string, discussion.Filter) []discussion.Comment); ok {
		r0 = rf(ctx, discussionID, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]discussion.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, discussion.Filter) error); ok {
		r1 = rf(ctx, discussionID, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiscussionRepository_GetAllComments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllComments'
type DiscussionRepository_GetAllComments_Call struct {
	*mock.Call
}

// GetAllComments is a helper method to define mock.On call
//  - ctx context.Context
//  - discussionID string
//  - flt discussion.Filter
func (_e *DiscussionRepository_Expecter) GetAllComments(ctx interface{}, discussionID interface{}, flt interface{}) *DiscussionRepository_GetAllComments_Call {
	return &DiscussionRepository_GetAllComments_Call{Call: _e.mock.On("GetAllComments", ctx, discussionID, flt)}
}

func (_c *DiscussionRepository_GetAllComments_Call) Run(run func(ctx context.Context, discussionID string, flt discussion.Filter)) *DiscussionRepository_GetAllComments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(discussion.Filter))
	})
	return _c
}

func (_c *DiscussionRepository_GetAllComments_Call) Return(_a0 []discussion.Comment, _a1 error) *DiscussionRepository_GetAllComments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetComment provides a mock function with given fields: ctx, commentID, discussionID
func (_m *DiscussionRepository) GetComment(ctx context.Context, commentID string, discussionID string) (discussion.Comment, error) {
	ret := _m.Called(ctx, commentID, discussionID)

	var r0 discussion.Comment
	if rf, ok := ret.Get(0).(func(context.Context, string, string) discussion.Comment); ok {
		r0 = rf(ctx, commentID, discussionID)
	} else {
		r0 = ret.Get(0).(discussion.Comment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, commentID, discussionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiscussionRepository_GetComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComment'
type DiscussionRepository_GetComment_Call struct {
	*mock.Call
}

// GetComment is a helper method to define mock.On call
//  - ctx context.Context
//  - commentID string
//  - discussionID string
func (_e *DiscussionRepository_Expecter) GetComment(ctx interface{}, commentID interface{}, discussionID interface{}) *DiscussionRepository_GetComment_Call {
	return &DiscussionRepository_GetComment_Call{Call: _e.mock.On("GetComment", ctx, commentID, discussionID)}
}

func (_c *DiscussionRepository_GetComment_Call) Run(run func(ctx context.Context, commentID string, discussionID string)) *DiscussionRepository_GetComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *DiscussionRepository_GetComment_Call) Return(_a0 discussion.Comment, _a1 error) *DiscussionRepository_GetComment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Patch provides a mock function with given fields: ctx, _a1
func (_m *DiscussionRepository) Patch(ctx context.Context, _a1 *discussion.Discussion) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discussion.Discussion) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscussionRepository_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type DiscussionRepository_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 *discussion.Discussion
func (_e *DiscussionRepository_Expecter) Patch(ctx interface{}, _a1 interface{}) *DiscussionRepository_Patch_Call {
	return &DiscussionRepository_Patch_Call{Call: _e.mock.On("Patch", ctx, _a1)}
}

func (_c *DiscussionRepository_Patch_Call) Run(run func(ctx context.Context, _a1 *discussion.Discussion)) *DiscussionRepository_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discussion.Discussion))
	})
	return _c
}

func (_c *DiscussionRepository_Patch_Call) Return(_a0 error) *DiscussionRepository_Patch_Call {
	_c.Call.Return(_a0)
	return _c
}

// UpdateComment provides a mock function with given fields: ctx, cmt
func (_m *DiscussionRepository) UpdateComment(ctx context.Context, cmt *discussion.Comment) error {
	ret := _m.Called(ctx, cmt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discussion.Comment) error); ok {
		r0 = rf(ctx, cmt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscussionRepository_UpdateComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateComment'
type DiscussionRepository_UpdateComment_Call struct {
	*mock.Call
}

// UpdateComment is a helper method to define mock.On call
//  - ctx context.Context
//  - cmt *discussion.Comment
func (_e *DiscussionRepository_Expecter) UpdateComment(ctx interface{}, cmt interface{}) *DiscussionRepository_UpdateComment_Call {
	return &DiscussionRepository_UpdateComment_Call{Call: _e.mock.On("UpdateComment", ctx, cmt)}
}

func (_c *DiscussionRepository_UpdateComment_Call) Run(run func(ctx context.Context, cmt *discussion.Comment)) *DiscussionRepository_UpdateComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discussion.Comment))
	})
	return _c
}

func (_c *DiscussionRepository_UpdateComment_Call) Return(_a0 error) *DiscussionRepository_UpdateComment_Call {
	_c.Call.Return(_a0)
	return _c
}
