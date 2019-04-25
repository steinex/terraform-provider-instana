// Code generated by MockGen. DO NOT EDIT.
// Source: instana/restapi/custom-event-specficiations-api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	restapi "github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	gomock "github.com/golang/mock/gomock"
)

// MockCustomEventSpecificationResource is a mock of CustomEventSpecificationResource interface
type MockCustomEventSpecificationResource struct {
	ctrl     *gomock.Controller
	recorder *MockCustomEventSpecificationResourceMockRecorder
}

// MockCustomEventSpecificationResourceMockRecorder is the mock recorder for MockCustomEventSpecificationResource
type MockCustomEventSpecificationResourceMockRecorder struct {
	mock *MockCustomEventSpecificationResource
}

// NewMockCustomEventSpecificationResource creates a new mock instance
func NewMockCustomEventSpecificationResource(ctrl *gomock.Controller) *MockCustomEventSpecificationResource {
	mock := &MockCustomEventSpecificationResource{ctrl: ctrl}
	mock.recorder = &MockCustomEventSpecificationResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomEventSpecificationResource) EXPECT() *MockCustomEventSpecificationResourceMockRecorder {
	return m.recorder
}

// GetOne mocks base method
func (m *MockCustomEventSpecificationResource) GetOne(id string) (restapi.CustomEventSpecification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", id)
	ret0, _ := ret[0].(restapi.CustomEventSpecification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockCustomEventSpecificationResourceMockRecorder) GetOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockCustomEventSpecificationResource)(nil).GetOne), id)
}

// Upsert mocks base method
func (m *MockCustomEventSpecificationResource) Upsert(spec restapi.CustomEventSpecification) (restapi.CustomEventSpecification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", spec)
	ret0, _ := ret[0].(restapi.CustomEventSpecification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockCustomEventSpecificationResourceMockRecorder) Upsert(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockCustomEventSpecificationResource)(nil).Upsert), spec)
}

// Delete mocks base method
func (m *MockCustomEventSpecificationResource) Delete(spec restapi.CustomEventSpecification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCustomEventSpecificationResourceMockRecorder) Delete(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomEventSpecificationResource)(nil).Delete), spec)
}

// DeleteByID mocks base method
func (m *MockCustomEventSpecificationResource) DeleteByID(specID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", specID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockCustomEventSpecificationResourceMockRecorder) DeleteByID(specID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockCustomEventSpecificationResource)(nil).DeleteByID), specID)
}
