// Code generated by MockGen. DO NOT EDIT.
// Source: biz.go

// Package mbiz is a generated GoMock package.
package mbiz

import (
	context "context"
	biz "kratos-base/app/catalog/service/internal/biz"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCatalogUseCase is a mock of CatalogUseCase interface.
type MockCatalogUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCatalogUseCaseMockRecorder
}

// MockCatalogUseCaseMockRecorder is the mock recorder for MockCatalogUseCase.
type MockCatalogUseCaseMockRecorder struct {
	mock *MockCatalogUseCase
}

// NewMockCatalogUseCase creates a new mock instance.
func NewMockCatalogUseCase(ctrl *gomock.Controller) *MockCatalogUseCase {
	mock := &MockCatalogUseCase{ctrl: ctrl}
	mock.recorder = &MockCatalogUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCatalogUseCase) EXPECT() *MockCatalogUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCatalogUseCase) Create(ctx context.Context, u *biz.Beer) (*biz.Beer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, u)
	ret0, _ := ret[0].(*biz.Beer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCatalogUseCaseMockRecorder) Create(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCatalogUseCase)(nil).Create), ctx, u)
}

// Get mocks base method.
func (m *MockCatalogUseCase) Get(ctx context.Context, id int64) (*biz.Beer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*biz.Beer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCatalogUseCaseMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCatalogUseCase)(nil).Get), ctx, id)
}
