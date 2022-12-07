// Code generated by MockGen. DO NOT EDIT.
// Source: product.go

// Package handlers is a generated GoMock package.
package handlers

import (
	http "net/http"
	reflect "reflect"
	schema "temp/schema"

	gomock "github.com/golang/mock/gomock"
)

// Mockprods is a mock of prods interface.
type Mockprods struct {
	ctrl     *gomock.Controller
	recorder *MockprodsMockRecorder
}

// MockprodsMockRecorder is the mock recorder for Mockprods.
type MockprodsMockRecorder struct {
	mock *Mockprods
}

// NewMockprods creates a new mock instance.
func NewMockprods(ctrl *gomock.Controller) *Mockprods {
	mock := &Mockprods{ctrl: ctrl}
	mock.recorder = &MockprodsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockprods) EXPECT() *MockprodsMockRecorder {
	return m.recorder
}

// Createproduct mocks base method.
func (m *Mockprods) Createproduct(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Createproduct", w, r)
}

// Createproduct indicates an expected call of Createproduct.
func (mr *MockprodsMockRecorder) Createproduct(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Createproduct", reflect.TypeOf((*Mockprods)(nil).Createproduct), w, r)
}

// get mocks base method.
func (m *Mockprods) get(arg0 schema.Product) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "get", arg0)
}

// get indicates an expected call of get.
func (mr *MockprodsMockRecorder) get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "get", reflect.TypeOf((*Mockprods)(nil).get), arg0)
}
