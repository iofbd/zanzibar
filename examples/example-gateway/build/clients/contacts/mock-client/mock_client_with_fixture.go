// Code generated by zanzibar
// @generated

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

package clientmock

import (
	"github.com/golang/mock/gomock"
)

// MockClientWithFixture is a mock of Client interface with preset fixture
type MockClientWithFixture struct {
	*MockClient
	fixture *ClientFixture

	saveContactsMock *SaveContactsMock
}

// Call is a thin wrapper around gomock.Call for exposing the methods that do not mutate the fixture related information
// like Return().
type Call struct {
	call *gomock.Call
}

// MaxTimes marks a fixture as callable up to a maximum number of times.
func (c Call) MaxTimes(max int) {
	c.call.MaxTimes(max)
}

// MinTimes marks a fixture as must be called a minimum number of times.
func (c Call) MinTimes(min int) {
	c.call.MinTimes(min)
}

// New creates a new mock instance
func New(ctrl *gomock.Controller, fixture *ClientFixture) *MockClientWithFixture {
	return &MockClientWithFixture{
		MockClient: NewMockClient(ctrl),
		fixture:    fixture,
	}
}

// EXPECT shadows the EXPECT method on the underlying mock client.
// It should not be called directly.
func (m *MockClientWithFixture) EXPECT() {
	panic("should not call EXPECT directly.")
}

// SaveContactsMock mocks the SaveContacts method
type SaveContactsMock struct {
	scenarios  *SaveContactsScenarios
	mockClient *MockClient
}

// ExpectSaveContacts returns an object that allows the caller to choose expected scenario for SaveContacts
func (m *MockClientWithFixture) ExpectSaveContacts() *SaveContactsMock {
	if m.saveContactsMock == nil {
		m.saveContactsMock = &SaveContactsMock{
			scenarios:  m.fixture.SaveContacts,
			mockClient: m.MockClient,
		}
	}
	return m.saveContactsMock
}

// Success sets the expected scenario as defined in the concrete fixture package
// github.com/uber/zanzibar/examples/example-gateway/clients/contacts/fixture
func (s *SaveContactsMock) Success() Call {
	f := s.scenarios.Success

	var arg0 interface{}
	arg0 = f.Arg0
	if f.Arg0Any {
		arg0 = gomock.Any()
	}
	var arg1 interface{}
	arg1 = f.Arg1
	if f.Arg1Any {
		arg1 = gomock.Any()
	}
	var arg2 interface{}
	arg2 = f.Arg2
	if f.Arg2Any {
		arg2 = gomock.Any()
	}
	var arg3 interface{}
	arg3 = f.Arg3
	if f.Arg3Any {
		arg3 = gomock.Any()
	}

	ret0 := f.Ret0
	ret1 := f.Ret1
	ret2 := f.Ret2
	ret3 := f.Ret3

	return Call{call: s.mockClient.EXPECT().SaveContacts(arg0, arg1, arg2, arg3).Return(ret0, ret1, ret2, ret3)}
}
