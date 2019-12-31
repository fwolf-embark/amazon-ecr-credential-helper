// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/awslabs/amazon-ecr-credential-helper/ecr-login/api (interfaces: ECRAPI)

package mock_api

import (
	ecr "github.com/aws/aws-sdk-go/service/ecr"
	gomock "github.com/golang/mock/gomock"
)

// Mock of ECRAPI interface
type MockECRAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockECRAPIRecorder
}

// Recorder for MockECRAPI (not exported)
type _MockECRAPIRecorder struct {
	mock *MockECRAPI
}

func NewMockECRAPI(ctrl *gomock.Controller) *MockECRAPI {
	mock := &MockECRAPI{ctrl: ctrl}
	mock.recorder = &_MockECRAPIRecorder{mock}
	return mock
}

func (_m *MockECRAPI) EXPECT() *_MockECRAPIRecorder {
	return _m.recorder
}

func (_m *MockECRAPI) GetAuthorizationToken(_param0 *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorizationToken", _param0)
	ret0, _ := ret[0].(*ecr.GetAuthorizationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) GetAuthorizationToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizationToken", arg0)
}
