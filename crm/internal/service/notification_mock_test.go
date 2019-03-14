// Code generated by MockGen. DO NOT EDIT.
// Source: crm/internal/service/notification.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	service0 "github.com/crusttech/crust/system/service"
	types "github.com/crusttech/crust/system/types"
	gomock "github.com/golang/mock/gomock"
	mail_v2 "gopkg.in/mail.v2"
	reflect "reflect"
)

// MockNotificationService is a mock of NotificationService interface
type MockNotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceMockRecorder
}

// MockNotificationServiceMockRecorder is the mock recorder for MockNotificationService
type MockNotificationServiceMockRecorder struct {
	mock *MockNotificationService
}

// NewMockNotificationService creates a new mock instance
func NewMockNotificationService(ctrl *gomock.Controller) *MockNotificationService {
	mock := &MockNotificationService{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotificationService) EXPECT() *MockNotificationServiceMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockNotificationService) With(ctx context.Context) NotificationService {
	ret := m.ctrl.Call(m, "With", ctx)
	ret0, _ := ret[0].(NotificationService)
	return ret0
}

// With indicates an expected call of With
func (mr *MockNotificationServiceMockRecorder) With(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockNotificationService)(nil).With), ctx)
}

// SendEmail mocks base method
func (m *MockNotificationService) SendEmail(message *mail_v2.Message) error {
	ret := m.ctrl.Call(m, "SendEmail", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmail indicates an expected call of SendEmail
func (mr *MockNotificationServiceMockRecorder) SendEmail(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockNotificationService)(nil).SendEmail), message)
}

// AttachEmailRecipients mocks base method
func (m *MockNotificationService) AttachEmailRecipients(message *mail_v2.Message, field string, recipients ...string) error {
	varargs := []interface{}{message, field}
	for _, a := range recipients {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AttachEmailRecipients", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachEmailRecipients indicates an expected call of AttachEmailRecipients
func (mr *MockNotificationServiceMockRecorder) AttachEmailRecipients(message, field interface{}, recipients ...interface{}) *gomock.Call {
	varargs := append([]interface{}{message, field}, recipients...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachEmailRecipients", reflect.TypeOf((*MockNotificationService)(nil).AttachEmailRecipients), varargs...)
}

// MocknotificationUserService is a mock of notificationUserService interface
type MocknotificationUserService struct {
	ctrl     *gomock.Controller
	recorder *MocknotificationUserServiceMockRecorder
}

// MocknotificationUserServiceMockRecorder is the mock recorder for MocknotificationUserService
type MocknotificationUserServiceMockRecorder struct {
	mock *MocknotificationUserService
}

// NewMocknotificationUserService creates a new mock instance
func NewMocknotificationUserService(ctrl *gomock.Controller) *MocknotificationUserService {
	mock := &MocknotificationUserService{ctrl: ctrl}
	mock.recorder = &MocknotificationUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocknotificationUserService) EXPECT() *MocknotificationUserServiceMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MocknotificationUserService) With(ctx context.Context) service0.UserService {
	ret := m.ctrl.Call(m, "With", ctx)
	ret0, _ := ret[0].(service0.UserService)
	return ret0
}

// With indicates an expected call of With
func (mr *MocknotificationUserServiceMockRecorder) With(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MocknotificationUserService)(nil).With), ctx)
}

// FindByID mocks base method
func (m *MocknotificationUserService) FindByID(userID uint64) (*types.User, error) {
	ret := m.ctrl.Call(m, "FindByID", userID)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MocknotificationUserServiceMockRecorder) FindByID(userID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MocknotificationUserService)(nil).FindByID), userID)
}
