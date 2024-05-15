// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// IPlantHandler is an autogenerated mock type for the IPlantHandler type
type IPlantHandler struct {
	mock.Mock
}

// GetPlant provides a mock function with given fields: c
func (_m *IPlantHandler) GetPlant(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutPlant provides a mock function with given fields: c
func (_m *IPlantHandler) PutPlant(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIPlantHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPlantHandler creates a new instance of IPlantHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPlantHandler(t mockConstructorTestingTNewIPlantHandler) *IPlantHandler {
	mock := &IPlantHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}