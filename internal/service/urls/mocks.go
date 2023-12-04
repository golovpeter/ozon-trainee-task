// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package urls is a generated GoMock package.
package urls

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUrlService is a mock of UrlService interface.
type MockUrlService struct {
	ctrl     *gomock.Controller
	recorder *MockUrlServiceMockRecorder
}

// MockUrlServiceMockRecorder is the mock recorder for MockUrlService.
type MockUrlServiceMockRecorder struct {
	mock *MockUrlService
}

// NewMockUrlService creates a new mock instance.
func NewMockUrlService(ctrl *gomock.Controller) *MockUrlService {
	mock := &MockUrlService{ctrl: ctrl}
	mock.recorder = &MockUrlServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlService) EXPECT() *MockUrlServiceMockRecorder {
	return m.recorder
}

// GetOriginalURL mocks base method.
func (m *MockUrlService) GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOriginalURL", ctx, in)
	ret0, _ := ret[0].(*GetOriginalURlOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOriginalURL indicates an expected call of GetOriginalURL.
func (mr *MockUrlServiceMockRecorder) GetOriginalURL(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOriginalURL", reflect.TypeOf((*MockUrlService)(nil).GetOriginalURL), ctx, in)
}

// ShortenURL mocks base method.
func (m *MockUrlService) ShortenURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShortenURL", ctx, in)
	ret0, _ := ret[0].(*ShortenURLOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShortenURL indicates an expected call of ShortenURL.
func (mr *MockUrlServiceMockRecorder) ShortenURL(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShortenURL", reflect.TypeOf((*MockUrlService)(nil).ShortenURL), ctx, in)
}