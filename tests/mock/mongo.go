// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/interfaces/mongo.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	gomock "github.com/golang/mock/gomock"
	interfaces "reddit/pkg/interfaces"
	reflect "reflect"
)

// MockMongoSession is a mock of MongoSession interface
type MockMongoSession struct {
	ctrl     *gomock.Controller
	recorder *MockMongoSessionMockRecorder
}

// MockMongoSessionMockRecorder is the mock recorder for MockMongoSession
type MockMongoSessionMockRecorder struct {
	mock *MockMongoSession
}

// NewMockMongoSession creates a new mock instance
func NewMockMongoSession(ctrl *gomock.Controller) *MockMongoSession {
	mock := &MockMongoSession{ctrl: ctrl}
	mock.recorder = &MockMongoSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMongoSession) EXPECT() *MockMongoSessionMockRecorder {
	return m.recorder
}

// DB mocks base method
func (m *MockMongoSession) DB(arg0 string) interfaces.MongoDatabase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB", arg0)
	ret0, _ := ret[0].(interfaces.MongoDatabase)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockMongoSessionMockRecorder) DB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockMongoSession)(nil).DB), arg0)
}

// MockMongoDatabase is a mock of MongoDatabase interface
type MockMongoDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockMongoDatabaseMockRecorder
}

// MockMongoDatabaseMockRecorder is the mock recorder for MockMongoDatabase
type MockMongoDatabaseMockRecorder struct {
	mock *MockMongoDatabase
}

// NewMockMongoDatabase creates a new mock instance
func NewMockMongoDatabase(ctrl *gomock.Controller) *MockMongoDatabase {
	mock := &MockMongoDatabase{ctrl: ctrl}
	mock.recorder = &MockMongoDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMongoDatabase) EXPECT() *MockMongoDatabaseMockRecorder {
	return m.recorder
}

// C mocks base method
func (m *MockMongoDatabase) C(name string) interfaces.MongoCollection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "C", name)
	ret0, _ := ret[0].(interfaces.MongoCollection)
	return ret0
}

// C indicates an expected call of C
func (mr *MockMongoDatabaseMockRecorder) C(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockMongoDatabase)(nil).C), name)
}

// MockMongoCollection is a mock of MongoCollection interface
type MockMongoCollection struct {
	ctrl     *gomock.Controller
	recorder *MockMongoCollectionMockRecorder
}

// MockMongoCollectionMockRecorder is the mock recorder for MockMongoCollection
type MockMongoCollectionMockRecorder struct {
	mock *MockMongoCollection
}

// NewMockMongoCollection creates a new mock instance
func NewMockMongoCollection(ctrl *gomock.Controller) *MockMongoCollection {
	mock := &MockMongoCollection{ctrl: ctrl}
	mock.recorder = &MockMongoCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMongoCollection) EXPECT() *MockMongoCollectionMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockMongoCollection) Insert(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockMongoCollectionMockRecorder) Insert(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMongoCollection)(nil).Insert), arg0...)
}

// Find mocks base method
func (m *MockMongoCollection) Find(arg0 interface{}) interfaces.MongoQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(interfaces.MongoQuery)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockMongoCollectionMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMongoCollection)(nil).Find), arg0)
}

// Update mocks base method
func (m *MockMongoCollection) Update(arg0, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockMongoCollectionMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMongoCollection)(nil).Update), arg0, arg1)
}

// Remove mocks base method
func (m *MockMongoCollection) Remove(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockMongoCollectionMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockMongoCollection)(nil).Remove), arg0)
}

// MockMongoQuery is a mock of MongoQuery interface
type MockMongoQuery struct {
	ctrl     *gomock.Controller
	recorder *MockMongoQueryMockRecorder
}

// MockMongoQueryMockRecorder is the mock recorder for MockMongoQuery
type MockMongoQueryMockRecorder struct {
	mock *MockMongoQuery
}

// NewMockMongoQuery creates a new mock instance
func NewMockMongoQuery(ctrl *gomock.Controller) *MockMongoQuery {
	mock := &MockMongoQuery{ctrl: ctrl}
	mock.recorder = &MockMongoQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMongoQuery) EXPECT() *MockMongoQueryMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockMongoQuery) All(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockMongoQueryMockRecorder) All(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockMongoQuery)(nil).All), arg0)
}

// One mocks base method
func (m *MockMongoQuery) One(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// One indicates an expected call of One
func (mr *MockMongoQueryMockRecorder) One(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockMongoQuery)(nil).One), arg0)
}
