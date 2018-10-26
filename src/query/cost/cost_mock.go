// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/query/cost (interfaces: PerQueryEnforcer,PerQueryEnforcerFactory)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package cost is a generated GoMock package.
package cost

import (
	"reflect"

	cost0 "github.com/m3db/m3/src/x/cost"

	"github.com/golang/mock/gomock"
)

// MockPerQueryEnforcer is a mock of PerQueryEnforcer interface
type MockPerQueryEnforcer struct {
	ctrl     *gomock.Controller
	recorder *MockPerQueryEnforcerMockRecorder
}

// MockPerQueryEnforcerMockRecorder is the mock recorder for MockPerQueryEnforcer
type MockPerQueryEnforcerMockRecorder struct {
	mock *MockPerQueryEnforcer
}

// NewMockPerQueryEnforcer creates a new mock instance
func NewMockPerQueryEnforcer(ctrl *gomock.Controller) *MockPerQueryEnforcer {
	mock := &MockPerQueryEnforcer{ctrl: ctrl}
	mock.recorder = &MockPerQueryEnforcerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPerQueryEnforcer) EXPECT() *MockPerQueryEnforcerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockPerQueryEnforcer) Add(arg0 cost0.Cost) cost0.Report {
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(cost0.Report)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockPerQueryEnforcerMockRecorder) Add(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPerQueryEnforcer)(nil).Add), arg0)
}

// Release mocks base method
func (m *MockPerQueryEnforcer) Release() {
	m.ctrl.Call(m, "Release")
}

// Release indicates an expected call of Release
func (mr *MockPerQueryEnforcerMockRecorder) Release() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockPerQueryEnforcer)(nil).Release))
}

// Report mocks base method
func (m *MockPerQueryEnforcer) Report() {
	m.ctrl.Call(m, "Report")
}

// Report indicates an expected call of Report
func (mr *MockPerQueryEnforcerMockRecorder) Report() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Report", reflect.TypeOf((*MockPerQueryEnforcer)(nil).Report))
}

// State mocks base method
func (m *MockPerQueryEnforcer) State() (cost0.Report, cost0.Limit) {
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(cost0.Report)
	ret1, _ := ret[1].(cost0.Limit)
	return ret0, ret1
}

// State indicates an expected call of State
func (mr *MockPerQueryEnforcerMockRecorder) State() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockPerQueryEnforcer)(nil).State))
}

// MockPerQueryEnforcerFactory is a mock of PerQueryEnforcerFactory interface
type MockPerQueryEnforcerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPerQueryEnforcerFactoryMockRecorder
}

// MockPerQueryEnforcerFactoryMockRecorder is the mock recorder for MockPerQueryEnforcerFactory
type MockPerQueryEnforcerFactoryMockRecorder struct {
	mock *MockPerQueryEnforcerFactory
}

// NewMockPerQueryEnforcerFactory creates a new mock instance
func NewMockPerQueryEnforcerFactory(ctrl *gomock.Controller) *MockPerQueryEnforcerFactory {
	mock := &MockPerQueryEnforcerFactory{ctrl: ctrl}
	mock.recorder = &MockPerQueryEnforcerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPerQueryEnforcerFactory) EXPECT() *MockPerQueryEnforcerFactoryMockRecorder {
	return m.recorder
}

// GlobalEnforcer mocks base method
func (m *MockPerQueryEnforcerFactory) GlobalEnforcer() *cost0.Enforcer {
	ret := m.ctrl.Call(m, "GlobalEnforcer")
	ret0, _ := ret[0].(*cost0.Enforcer)
	return ret0
}

// GlobalEnforcer indicates an expected call of GlobalEnforcer
func (mr *MockPerQueryEnforcerFactoryMockRecorder) GlobalEnforcer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalEnforcer", reflect.TypeOf((*MockPerQueryEnforcerFactory)(nil).GlobalEnforcer))
}

// New mocks base method
func (m *MockPerQueryEnforcerFactory) New() PerQueryEnforcer {
	ret := m.ctrl.Call(m, "New")
	ret0, _ := ret[0].(PerQueryEnforcer)
	return ret0
}

// New indicates an expected call of New
func (mr *MockPerQueryEnforcerFactoryMockRecorder) New() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockPerQueryEnforcerFactory)(nil).New))
}
