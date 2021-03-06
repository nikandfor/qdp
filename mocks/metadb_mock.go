// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/qiwitech/qdp/proto/metadbpb (interfaces: MetaDBServiceInterface)

package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	metadbpb "github.com/qiwitech/qdp/proto/metadbpb"
)

// Mock of MetaDBServiceInterface interface
type MockMetaDBServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockMetaDBServiceInterfaceRecorder
}

// Recorder for MockMetaDBServiceInterface (not exported)
type _MockMetaDBServiceInterfaceRecorder struct {
	mock *MockMetaDBServiceInterface
}

func NewMockMetaDBServiceInterface(ctrl *gomock.Controller) *MockMetaDBServiceInterface {
	mock := &MockMetaDBServiceInterface{ctrl: ctrl}
	mock.recorder = &_MockMetaDBServiceInterfaceRecorder{mock}
	return mock
}

func (_m *MockMetaDBServiceInterface) EXPECT() *_MockMetaDBServiceInterfaceRecorder {
	return _m.recorder
}

func (_m *MockMetaDBServiceInterface) Delete(_param0 context.Context, _param1 *metadbpb.DeleteRequest) (*metadbpb.DeleteResponse, error) {
	ret := _m.ctrl.Call(_m, "Delete", _param0, _param1)
	ret0, _ := ret[0].(*metadbpb.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMetaDBServiceInterfaceRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1)
}

func (_m *MockMetaDBServiceInterface) Get(_param0 context.Context, _param1 *metadbpb.GetRequest) (*metadbpb.GetResponse, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1)
	ret0, _ := ret[0].(*metadbpb.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMetaDBServiceInterfaceRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockMetaDBServiceInterface) GetMulti(_param0 context.Context, _param1 *metadbpb.GetMultiRequest) (*metadbpb.GetMultiResponse, error) {
	ret := _m.ctrl.Call(_m, "GetMulti", _param0, _param1)
	ret0, _ := ret[0].(*metadbpb.GetMultiResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMetaDBServiceInterfaceRecorder) GetMulti(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMulti", arg0, arg1)
}

func (_m *MockMetaDBServiceInterface) Put(_param0 context.Context, _param1 *metadbpb.PutRequest) (*metadbpb.PutResponse, error) {
	ret := _m.ctrl.Call(_m, "Put", _param0, _param1)
	ret0, _ := ret[0].(*metadbpb.PutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMetaDBServiceInterfaceRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0, arg1)
}

func (_m *MockMetaDBServiceInterface) Search(_param0 context.Context, _param1 *metadbpb.SearchRequest) (*metadbpb.SearchResponse, error) {
	ret := _m.ctrl.Call(_m, "Search", _param0, _param1)
	ret0, _ := ret[0].(*metadbpb.SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMetaDBServiceInterfaceRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Search", arg0, arg1)
}
