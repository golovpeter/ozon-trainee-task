// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package urls is a generated GoMock package.
package urls

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	sqlx "github.com/jmoiron/sqlx"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetOriginalURL mocks base method.
func (m *MockRepository) GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOriginalURL", ctx, in)
	ret0, _ := ret[0].(*GetOriginalURlOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOriginalURL indicates an expected call of GetOriginalURL.
func (mr *MockRepositoryMockRecorder) GetOriginalURL(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOriginalURL", reflect.TypeOf((*MockRepository)(nil).GetOriginalURL), ctx, in)
}

// SaveShortenedURL mocks base method.
func (m *MockRepository) SaveShortenedURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveShortenedURL", ctx, in)
	ret0, _ := ret[0].(*ShortenURLOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveShortenedURL indicates an expected call of SaveShortenedURL.
func (mr *MockRepositoryMockRecorder) SaveShortenedURL(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveShortenedURL", reflect.TypeOf((*MockRepository)(nil).SaveShortenedURL), ctx, in)
}

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// BeginTxx mocks base method.
func (m *MockDatabase) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTxx", ctx, opts)
	ret0, _ := ret[0].(*sqlx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTxx indicates an expected call of BeginTxx.
func (mr *MockDatabaseMockRecorder) BeginTxx(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTxx", reflect.TypeOf((*MockDatabase)(nil).BeginTxx), ctx, opts)
}

// ExecContext mocks base method.
func (m *MockDatabase) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockDatabaseMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockDatabase)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockDatabase) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockDatabaseMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockDatabase)(nil).GetContext), varargs...)
}
