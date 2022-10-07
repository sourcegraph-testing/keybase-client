// Automatically generated by MockGen. DO NOT EDIT!
// Source: codec.go

package kbfscodec

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mock of Codec interface
type MockCodec struct {
	ctrl     *gomock.Controller
	recorder *_MockCodecRecorder
}

// Recorder for MockCodec (not exported)
type _MockCodecRecorder struct {
	mock *MockCodec
}

func NewMockCodec(ctrl *gomock.Controller) *MockCodec {
	mock := &MockCodec{ctrl: ctrl}
	mock.recorder = &_MockCodecRecorder{mock}
	return mock
}

func (_m *MockCodec) EXPECT() *_MockCodecRecorder {
	return _m.recorder
}

func (_m *MockCodec) Decode(buf []byte, obj any) error {
	ret := _m.ctrl.Call(_m, "Decode", buf, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCodecRecorder) Decode(arg0, arg1 any) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Decode", arg0, arg1)
}

func (_m *MockCodec) Encode(obj any) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "Encode", obj)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCodecRecorder) Encode(arg0 any) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Encode", arg0)
}

func (_m *MockCodec) RegisterType(rt reflect.Type, code ExtCode) {
	_m.ctrl.Call(_m, "RegisterType", rt, code)
}

func (_mr *_MockCodecRecorder) RegisterType(arg0, arg1 any) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterType", arg0, arg1)
}

func (_m *MockCodec) RegisterIfaceSliceType(rt reflect.Type, code ExtCode, typer func(any) reflect.Value) {
	_m.ctrl.Call(_m, "RegisterIfaceSliceType", rt, code, typer)
}

func (_mr *_MockCodecRecorder) RegisterIfaceSliceType(arg0, arg1, arg2 any) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterIfaceSliceType", arg0, arg1, arg2)
}
