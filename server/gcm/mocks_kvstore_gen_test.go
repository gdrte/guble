// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/smancke/guble/server/kvstore (interfaces: KVStore)

package gcm

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of KVStore interface
type MockKVStore struct {
	ctrl     *gomock.Controller
	recorder *_MockKVStoreRecorder
}

// Recorder for MockKVStore (not exported)
type _MockKVStoreRecorder struct {
	mock *MockKVStore
}

func NewMockKVStore(ctrl *gomock.Controller) *MockKVStore {
	mock := &MockKVStore{ctrl: ctrl}
	mock.recorder = &_MockKVStoreRecorder{mock}
	return mock
}

func (_m *MockKVStore) EXPECT() *_MockKVStoreRecorder {
	return _m.recorder
}

func (_m *MockKVStore) Delete(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKVStoreRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1)
}

func (_m *MockKVStore) Get(_param0 string, _param1 string) ([]byte, bool, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockKVStoreRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockKVStore) Iterate(_param0 string, _param1 string) chan [2]string {
	ret := _m.ctrl.Call(_m, "Iterate", _param0, _param1)
	ret0, _ := ret[0].(chan [2]string)
	return ret0
}

func (_mr *_MockKVStoreRecorder) Iterate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Iterate", arg0, arg1)
}

func (_m *MockKVStore) IterateKeys(_param0 string, _param1 string) chan string {
	ret := _m.ctrl.Call(_m, "IterateKeys", _param0, _param1)
	ret0, _ := ret[0].(chan string)
	return ret0
}

func (_mr *_MockKVStoreRecorder) IterateKeys(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IterateKeys", arg0, arg1)
}

func (_m *MockKVStore) Put(_param0 string, _param1 string, _param2 []byte) error {
	ret := _m.ctrl.Call(_m, "Put", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKVStoreRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0, arg1, arg2)
}
