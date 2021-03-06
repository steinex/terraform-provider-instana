// Code generated by MockGen. DO NOT EDIT.
// Source: instana/restapi/application-configs-api.go

// Package mock_restapi is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	restapi "github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	gomock "github.com/golang/mock/gomock"
)

// MockApplicationConfigResource is a mock of ApplicationConfigResource interface
type MockApplicationConfigResource struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationConfigResourceMockRecorder
}

// MockApplicationConfigResourceMockRecorder is the mock recorder for MockApplicationConfigResource
type MockApplicationConfigResourceMockRecorder struct {
	mock *MockApplicationConfigResource
}

// NewMockApplicationConfigResource creates a new mock instance
func NewMockApplicationConfigResource(ctrl *gomock.Controller) *MockApplicationConfigResource {
	mock := &MockApplicationConfigResource{ctrl: ctrl}
	mock.recorder = &MockApplicationConfigResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationConfigResource) EXPECT() *MockApplicationConfigResourceMockRecorder {
	return m.recorder
}

// GetOne mocks base method
func (m *MockApplicationConfigResource) GetOne(id string) (restapi.ApplicationConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", id)
	ret0, _ := ret[0].(restapi.ApplicationConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockApplicationConfigResourceMockRecorder) GetOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockApplicationConfigResource)(nil).GetOne), id)
}

// Upsert mocks base method
func (m *MockApplicationConfigResource) Upsert(rule restapi.ApplicationConfig) (restapi.ApplicationConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", rule)
	ret0, _ := ret[0].(restapi.ApplicationConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockApplicationConfigResourceMockRecorder) Upsert(rule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockApplicationConfigResource)(nil).Upsert), rule)
}

// Delete mocks base method
func (m *MockApplicationConfigResource) Delete(rule restapi.ApplicationConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockApplicationConfigResourceMockRecorder) Delete(rule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApplicationConfigResource)(nil).Delete), rule)
}

// DeleteByID mocks base method
func (m *MockApplicationConfigResource) DeleteByID(applicationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", applicationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockApplicationConfigResourceMockRecorder) DeleteByID(applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockApplicationConfigResource)(nil).DeleteByID), applicationID)
}

// MockMatchExpression is a mock of MatchExpression interface
type MockMatchExpression struct {
	ctrl     *gomock.Controller
	recorder *MockMatchExpressionMockRecorder
}

// MockMatchExpressionMockRecorder is the mock recorder for MockMatchExpression
type MockMatchExpressionMockRecorder struct {
	mock *MockMatchExpression
}

// NewMockMatchExpression creates a new mock instance
func NewMockMatchExpression(ctrl *gomock.Controller) *MockMatchExpression {
	mock := &MockMatchExpression{ctrl: ctrl}
	mock.recorder = &MockMatchExpressionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMatchExpression) EXPECT() *MockMatchExpressionMockRecorder {
	return m.recorder
}

// GetType mocks base method
func (m *MockMatchExpression) GetType() restapi.MatchExpressionType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(restapi.MatchExpressionType)
	return ret0
}

// GetType indicates an expected call of GetType
func (mr *MockMatchExpressionMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockMatchExpression)(nil).GetType))
}

// Validate mocks base method
func (m *MockMatchExpression) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockMatchExpressionMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockMatchExpression)(nil).Validate))
}
