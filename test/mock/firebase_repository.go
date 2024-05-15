// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	authentication "authenticator-backend/domain/model/authentication"

	mock "github.com/stretchr/testify/mock"
)

// FirebaseRepository is an autogenerated mock type for the FirebaseRepository type
type FirebaseRepository struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: uid, newPassword
func (_m *FirebaseRepository) ChangePassword(uid string, newPassword authentication.Password) error {
	ret := _m.Called(uid, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, authentication.Password) error); ok {
		r0 = rf(uid, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RefreshToken provides a mock function with given fields: refreshToken
func (_m *FirebaseRepository) RefreshToken(refreshToken string) (string, error) {
	ret := _m.Called(refreshToken)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(refreshToken)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(refreshToken)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignInWithPassword provides a mock function with given fields: email, password
func (_m *FirebaseRepository) SignInWithPassword(email string, password string) (authentication.LoginResult, error) {
	ret := _m.Called(email, password)

	var r0 authentication.LoginResult
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (authentication.LoginResult, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) authentication.LoginResult); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(authentication.LoginResult)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyIDToken provides a mock function with given fields: idToken
func (_m *FirebaseRepository) VerifyIDToken(idToken string) (authentication.Claims, error) {
	ret := _m.Called(idToken)

	var r0 authentication.Claims
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (authentication.Claims, error)); ok {
		return rf(idToken)
	}
	if rf, ok := ret.Get(0).(func(string) authentication.Claims); ok {
		r0 = rf(idToken)
	} else {
		r0 = ret.Get(0).(authentication.Claims)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFirebaseRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewFirebaseRepository creates a new instance of FirebaseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFirebaseRepository(t mockConstructorTestingTNewFirebaseRepository) *FirebaseRepository {
	mock := &FirebaseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
