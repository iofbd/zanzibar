// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/zanzibar/examples/example-gateway/build/clients/bar (interfaces: Client)

// Package clientmock is a generated GoMock package.
package clientmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	runtime "github.com/uber/zanzibar/runtime"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ArgNotStruct mocks base method
func (m *MockClient) ArgNotStruct(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_ArgNotStruct_Args) (map[string]string, error) {
	ret := m.ctrl.Call(m, "ArgNotStruct", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArgNotStruct indicates an expected call of ArgNotStruct
func (mr *MockClientMockRecorder) ArgNotStruct(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArgNotStruct", reflect.TypeOf((*MockClient)(nil).ArgNotStruct), arg0, arg1, arg2)
}

// ArgWithHeaders mocks base method
func (m *MockClient) ArgWithHeaders(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_ArgWithHeaders_Args) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "ArgWithHeaders", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ArgWithHeaders indicates an expected call of ArgWithHeaders
func (mr *MockClientMockRecorder) ArgWithHeaders(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArgWithHeaders", reflect.TypeOf((*MockClient)(nil).ArgWithHeaders), arg0, arg1, arg2)
}

// ArgWithManyQueryParams mocks base method
func (m *MockClient) ArgWithManyQueryParams(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_ArgWithManyQueryParams_Args) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "ArgWithManyQueryParams", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ArgWithManyQueryParams indicates an expected call of ArgWithManyQueryParams
func (mr *MockClientMockRecorder) ArgWithManyQueryParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArgWithManyQueryParams", reflect.TypeOf((*MockClient)(nil).ArgWithManyQueryParams), arg0, arg1, arg2)
}

// ArgWithNestedQueryParams mocks base method
func (m *MockClient) ArgWithNestedQueryParams(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_ArgWithNestedQueryParams_Args) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "ArgWithNestedQueryParams", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ArgWithNestedQueryParams indicates an expected call of ArgWithNestedQueryParams
func (mr *MockClientMockRecorder) ArgWithNestedQueryParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArgWithNestedQueryParams", reflect.TypeOf((*MockClient)(nil).ArgWithNestedQueryParams), arg0, arg1, arg2)
}

// ArgWithParams mocks base method
func (m *MockClient) ArgWithParams(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_ArgWithParams_Args) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "ArgWithParams", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ArgWithParams indicates an expected call of ArgWithParams
func (mr *MockClientMockRecorder) ArgWithParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArgWithParams", reflect.TypeOf((*MockClient)(nil).ArgWithParams), arg0, arg1, arg2)
}

// ArgWithQueryHeader mocks base method
func (m *MockClient) ArgWithQueryHeader(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_ArgWithQueryHeader_Args) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "ArgWithQueryHeader", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ArgWithQueryHeader indicates an expected call of ArgWithQueryHeader
func (mr *MockClientMockRecorder) ArgWithQueryHeader(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArgWithQueryHeader", reflect.TypeOf((*MockClient)(nil).ArgWithQueryHeader), arg0, arg1, arg2)
}

// ArgWithQueryParams mocks base method
func (m *MockClient) ArgWithQueryParams(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_ArgWithQueryParams_Args) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "ArgWithQueryParams", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ArgWithQueryParams indicates an expected call of ArgWithQueryParams
func (mr *MockClientMockRecorder) ArgWithQueryParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArgWithQueryParams", reflect.TypeOf((*MockClient)(nil).ArgWithQueryParams), arg0, arg1, arg2)
}

// EchoBinary mocks base method
func (m *MockClient) EchoBinary(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoBinary_Args) ([]byte, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoBinary", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoBinary indicates an expected call of EchoBinary
func (mr *MockClientMockRecorder) EchoBinary(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoBinary", reflect.TypeOf((*MockClient)(nil).EchoBinary), arg0, arg1, arg2)
}

// EchoBool mocks base method
func (m *MockClient) EchoBool(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoBool_Args) (bool, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoBool", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoBool indicates an expected call of EchoBool
func (mr *MockClientMockRecorder) EchoBool(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoBool", reflect.TypeOf((*MockClient)(nil).EchoBool), arg0, arg1, arg2)
}

// EchoDouble mocks base method
func (m *MockClient) EchoDouble(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoDouble_Args) (float64, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoDouble", arg0, arg1, arg2)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoDouble indicates an expected call of EchoDouble
func (mr *MockClientMockRecorder) EchoDouble(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoDouble", reflect.TypeOf((*MockClient)(nil).EchoDouble), arg0, arg1, arg2)
}

// EchoEnum mocks base method
func (m *MockClient) EchoEnum(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoEnum_Args) (bar.Fruit, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoEnum", arg0, arg1, arg2)
	ret0, _ := ret[0].(bar.Fruit)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoEnum indicates an expected call of EchoEnum
func (mr *MockClientMockRecorder) EchoEnum(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoEnum", reflect.TypeOf((*MockClient)(nil).EchoEnum), arg0, arg1, arg2)
}

// EchoI16 mocks base method
func (m *MockClient) EchoI16(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoI16_Args) (int16, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoI16", arg0, arg1, arg2)
	ret0, _ := ret[0].(int16)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoI16 indicates an expected call of EchoI16
func (mr *MockClientMockRecorder) EchoI16(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI16", reflect.TypeOf((*MockClient)(nil).EchoI16), arg0, arg1, arg2)
}

// EchoI32 mocks base method
func (m *MockClient) EchoI32(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoI32_Args) (int32, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoI32", arg0, arg1, arg2)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoI32 indicates an expected call of EchoI32
func (mr *MockClientMockRecorder) EchoI32(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI32", reflect.TypeOf((*MockClient)(nil).EchoI32), arg0, arg1, arg2)
}

// EchoI32Map mocks base method
func (m *MockClient) EchoI32Map(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoI32Map_Args) (map[int32]*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoI32Map", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[int32]*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoI32Map indicates an expected call of EchoI32Map
func (mr *MockClientMockRecorder) EchoI32Map(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI32Map", reflect.TypeOf((*MockClient)(nil).EchoI32Map), arg0, arg1, arg2)
}

// EchoI64 mocks base method
func (m *MockClient) EchoI64(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoI64_Args) (int64, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoI64", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoI64 indicates an expected call of EchoI64
func (mr *MockClientMockRecorder) EchoI64(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI64", reflect.TypeOf((*MockClient)(nil).EchoI64), arg0, arg1, arg2)
}

// EchoI8 mocks base method
func (m *MockClient) EchoI8(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoI8_Args) (int8, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoI8", arg0, arg1, arg2)
	ret0, _ := ret[0].(int8)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoI8 indicates an expected call of EchoI8
func (mr *MockClientMockRecorder) EchoI8(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI8", reflect.TypeOf((*MockClient)(nil).EchoI8), arg0, arg1, arg2)
}

// EchoString mocks base method
func (m *MockClient) EchoString(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoString_Args) (string, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoString", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoString indicates an expected call of EchoString
func (mr *MockClientMockRecorder) EchoString(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoString", reflect.TypeOf((*MockClient)(nil).EchoString), arg0, arg1, arg2)
}

// EchoStringList mocks base method
func (m *MockClient) EchoStringList(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoStringList_Args) ([]string, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoStringList", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoStringList indicates an expected call of EchoStringList
func (mr *MockClientMockRecorder) EchoStringList(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStringList", reflect.TypeOf((*MockClient)(nil).EchoStringList), arg0, arg1, arg2)
}

// EchoStringMap mocks base method
func (m *MockClient) EchoStringMap(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoStringMap_Args) (map[string]*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoStringMap", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoStringMap indicates an expected call of EchoStringMap
func (mr *MockClientMockRecorder) EchoStringMap(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStringMap", reflect.TypeOf((*MockClient)(nil).EchoStringMap), arg0, arg1, arg2)
}

// EchoStringSet mocks base method
func (m *MockClient) EchoStringSet(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoStringSet_Args) (map[string]struct{}, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoStringSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]struct{})
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoStringSet indicates an expected call of EchoStringSet
func (mr *MockClientMockRecorder) EchoStringSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStringSet", reflect.TypeOf((*MockClient)(nil).EchoStringSet), arg0, arg1, arg2)
}

// EchoStructList mocks base method
func (m *MockClient) EchoStructList(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoStructList_Args) ([]*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoStructList", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoStructList indicates an expected call of EchoStructList
func (mr *MockClientMockRecorder) EchoStructList(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStructList", reflect.TypeOf((*MockClient)(nil).EchoStructList), arg0, arg1, arg2)
}

// EchoStructSet mocks base method
func (m *MockClient) EchoStructSet(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoStructSet_Args) ([]*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoStructSet", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoStructSet indicates an expected call of EchoStructSet
func (mr *MockClientMockRecorder) EchoStructSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStructSet", reflect.TypeOf((*MockClient)(nil).EchoStructSet), arg0, arg1, arg2)
}

// EchoTypedef mocks base method
func (m *MockClient) EchoTypedef(arg0 context.Context, arg1 map[string]string, arg2 *bar.Echo_EchoTypedef_Args) (bar.UUID, map[string]string, error) {
	ret := m.ctrl.Call(m, "EchoTypedef", arg0, arg1, arg2)
	ret0, _ := ret[0].(bar.UUID)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoTypedef indicates an expected call of EchoTypedef
func (mr *MockClientMockRecorder) EchoTypedef(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoTypedef", reflect.TypeOf((*MockClient)(nil).EchoTypedef), arg0, arg1, arg2)
}

// HTTPClient mocks base method
func (m *MockClient) HTTPClient() *runtime.HTTPClient {
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*runtime.HTTPClient)
	return ret0
}

// HTTPClient indicates an expected call of HTTPClient
func (mr *MockClientMockRecorder) HTTPClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockClient)(nil).HTTPClient))
}

// Hello mocks base method
func (m *MockClient) Hello(arg0 context.Context, arg1 map[string]string) (string, map[string]string, error) {
	ret := m.ctrl.Call(m, "Hello", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Hello indicates an expected call of Hello
func (mr *MockClientMockRecorder) Hello(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockClient)(nil).Hello), arg0, arg1)
}

// MissingArg mocks base method
func (m *MockClient) MissingArg(arg0 context.Context, arg1 map[string]string) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "MissingArg", arg0, arg1)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MissingArg indicates an expected call of MissingArg
func (mr *MockClientMockRecorder) MissingArg(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MissingArg", reflect.TypeOf((*MockClient)(nil).MissingArg), arg0, arg1)
}

// NoRequest mocks base method
func (m *MockClient) NoRequest(arg0 context.Context, arg1 map[string]string) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "NoRequest", arg0, arg1)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NoRequest indicates an expected call of NoRequest
func (mr *MockClientMockRecorder) NoRequest(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoRequest", reflect.TypeOf((*MockClient)(nil).NoRequest), arg0, arg1)
}

// Normal mocks base method
func (m *MockClient) Normal(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_Normal_Args) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "Normal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Normal indicates an expected call of Normal
func (mr *MockClientMockRecorder) Normal(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Normal", reflect.TypeOf((*MockClient)(nil).Normal), arg0, arg1, arg2)
}

// NormalRecur mocks base method
func (m *MockClient) NormalRecur(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_NormalRecur_Args) (*bar.BarResponseRecur, map[string]string, error) {
	ret := m.ctrl.Call(m, "NormalRecur", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponseRecur)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NormalRecur indicates an expected call of NormalRecur
func (mr *MockClientMockRecorder) NormalRecur(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NormalRecur", reflect.TypeOf((*MockClient)(nil).NormalRecur), arg0, arg1, arg2)
}

// TooManyArgs mocks base method
func (m *MockClient) TooManyArgs(arg0 context.Context, arg1 map[string]string, arg2 *bar.Bar_TooManyArgs_Args) (*bar.BarResponse, map[string]string, error) {
	ret := m.ctrl.Call(m, "TooManyArgs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bar.BarResponse)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TooManyArgs indicates an expected call of TooManyArgs
func (mr *MockClientMockRecorder) TooManyArgs(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TooManyArgs", reflect.TypeOf((*MockClient)(nil).TooManyArgs), arg0, arg1, arg2)
}
