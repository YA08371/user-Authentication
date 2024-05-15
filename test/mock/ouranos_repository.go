// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	traceability "authenticator-backend/domain/model/traceability"
)

// OuranosRepository is an autogenerated mock type for the OuranosRepository type
type OuranosRepository struct {
	mock.Mock
}

// CreatePlant provides a mock function with given fields: e
func (_m *OuranosRepository) CreatePlant(e traceability.PlantEntityModel) (traceability.PlantEntityModel, error) {
	ret := _m.Called(e)

	var r0 traceability.PlantEntityModel
	var r1 error
	if rf, ok := ret.Get(0).(func(traceability.PlantEntityModel) (traceability.PlantEntityModel, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(traceability.PlantEntityModel) traceability.PlantEntityModel); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(traceability.PlantEntityModel)
	}

	if rf, ok := ret.Get(1).(func(traceability.PlantEntityModel) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCFPCertificateByCFPID provides a mock function with given fields: cfpID
func (_m *OuranosRepository) DeleteCFPCertificateByCFPID(cfpID string) error {
	ret := _m.Called(cfpID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(cfpID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCFPInformation provides a mock function with given fields: cfpID
func (_m *OuranosRepository) DeleteCFPInformation(cfpID string) error {
	ret := _m.Called(cfpID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(cfpID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteParts provides a mock function with given fields: traceID
func (_m *OuranosRepository) DeleteParts(traceID string) error {
	ret := _m.Called(traceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(traceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePartsStructure provides a mock function with given fields: traceID
func (_m *OuranosRepository) DeletePartsStructure(traceID string) error {
	ret := _m.Called(traceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(traceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePlant provides a mock function with given fields: plantID
func (_m *OuranosRepository) DeletePlant(plantID string) error {
	ret := _m.Called(plantID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(plantID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRequestStatusByTradeID provides a mock function with given fields: tradeID
func (_m *OuranosRepository) DeleteRequestStatusByTradeID(tradeID string) error {
	ret := _m.Called(tradeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(tradeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTrade provides a mock function with given fields: tradeID
func (_m *OuranosRepository) DeleteTrade(tradeID string) error {
	ret := _m.Called(tradeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(tradeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCFPInformation provides a mock function with given fields: traceID
func (_m *OuranosRepository) GetCFPInformation(traceID string) (traceability.CfpEntityModel, error) {
	ret := _m.Called(traceID)

	var r0 traceability.CfpEntityModel
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (traceability.CfpEntityModel, error)); ok {
		return rf(traceID)
	}
	if rf, ok := ret.Get(0).(func(string) traceability.CfpEntityModel); ok {
		r0 = rf(traceID)
	} else {
		r0 = ret.Get(0).(traceability.CfpEntityModel)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(traceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperator provides a mock function with given fields: operatorID
func (_m *OuranosRepository) GetOperator(operatorID string) (traceability.OperatorEntityModel, error) {
	ret := _m.Called(operatorID)

	var r0 traceability.OperatorEntityModel
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (traceability.OperatorEntityModel, error)); ok {
		return rf(operatorID)
	}
	if rf, ok := ret.Get(0).(func(string) traceability.OperatorEntityModel); ok {
		r0 = rf(operatorID)
	} else {
		r0 = ret.Get(0).(traceability.OperatorEntityModel)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(operatorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperatorByOpenOperatorID provides a mock function with given fields: openOperatorID
func (_m *OuranosRepository) GetOperatorByOpenOperatorID(openOperatorID string) (traceability.OperatorEntityModel, error) {
	ret := _m.Called(openOperatorID)

	var r0 traceability.OperatorEntityModel
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (traceability.OperatorEntityModel, error)); ok {
		return rf(openOperatorID)
	}
	if rf, ok := ret.Get(0).(func(string) traceability.OperatorEntityModel); ok {
		r0 = rf(openOperatorID)
	} else {
		r0 = ret.Get(0).(traceability.OperatorEntityModel)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(openOperatorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperators provides a mock function with given fields: operatorIDs
func (_m *OuranosRepository) GetOperators(operatorIDs []string) (traceability.OperatorEntityModels, error) {
	ret := _m.Called(operatorIDs)

	var r0 traceability.OperatorEntityModels
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) (traceability.OperatorEntityModels, error)); ok {
		return rf(operatorIDs)
	}
	if rf, ok := ret.Get(0).(func([]string) traceability.OperatorEntityModels); ok {
		r0 = rf(operatorIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(traceability.OperatorEntityModels)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(operatorIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlant provides a mock function with given fields: operatorID, plantID
func (_m *OuranosRepository) GetPlant(operatorID string, plantID string) (traceability.PlantEntityModel, error) {
	ret := _m.Called(operatorID, plantID)

	var r0 traceability.PlantEntityModel
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (traceability.PlantEntityModel, error)); ok {
		return rf(operatorID, plantID)
	}
	if rf, ok := ret.Get(0).(func(string, string) traceability.PlantEntityModel); ok {
		r0 = rf(operatorID, plantID)
	} else {
		r0 = ret.Get(0).(traceability.PlantEntityModel)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(operatorID, plantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListParts provides a mock function with given fields: getPlantPartsModel
func (_m *OuranosRepository) ListParts(getPlantPartsModel traceability.GetPartsModel) (traceability.PartsModels, error) {
	ret := _m.Called(getPlantPartsModel)

	var r0 traceability.PartsModels
	var r1 error
	if rf, ok := ret.Get(0).(func(traceability.GetPartsModel) (traceability.PartsModels, error)); ok {
		return rf(getPlantPartsModel)
	}
	if rf, ok := ret.Get(0).(func(traceability.GetPartsModel) traceability.PartsModels); ok {
		r0 = rf(getPlantPartsModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(traceability.PartsModels)
		}
	}

	if rf, ok := ret.Get(1).(func(traceability.GetPartsModel) error); ok {
		r1 = rf(getPlantPartsModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPlantsByOperatorID provides a mock function with given fields: operatorID
func (_m *OuranosRepository) ListPlantsByOperatorID(operatorID string) (traceability.PlantEntityModels, error) {
	ret := _m.Called(operatorID)

	var r0 traceability.PlantEntityModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (traceability.PlantEntityModels, error)); ok {
		return rf(operatorID)
	}
	if rf, ok := ret.Get(0).(func(string) traceability.PlantEntityModels); ok {
		r0 = rf(operatorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(traceability.PlantEntityModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(operatorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTradesByOperatorID provides a mock function with given fields: operatorID
func (_m *OuranosRepository) ListTradesByOperatorID(operatorID string) (traceability.TradeEntityModels, error) {
	ret := _m.Called(operatorID)

	var r0 traceability.TradeEntityModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (traceability.TradeEntityModels, error)); ok {
		return rf(operatorID)
	}
	if rf, ok := ret.Get(0).(func(string) traceability.TradeEntityModels); ok {
		r0 = rf(operatorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(traceability.TradeEntityModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(operatorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutOperator provides a mock function with given fields: e
func (_m *OuranosRepository) PutOperator(e traceability.OperatorEntityModel) (traceability.OperatorEntityModel, error) {
	ret := _m.Called(e)

	var r0 traceability.OperatorEntityModel
	var r1 error
	if rf, ok := ret.Get(0).(func(traceability.OperatorEntityModel) (traceability.OperatorEntityModel, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(traceability.OperatorEntityModel) traceability.OperatorEntityModel); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(traceability.OperatorEntityModel)
	}

	if rf, ok := ret.Get(1).(func(traceability.OperatorEntityModel) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlant provides a mock function with given fields: e
func (_m *OuranosRepository) UpdatePlant(e traceability.PlantEntityModel) (traceability.PlantEntityModel, error) {
	ret := _m.Called(e)

	var r0 traceability.PlantEntityModel
	var r1 error
	if rf, ok := ret.Get(0).(func(traceability.PlantEntityModel) (traceability.PlantEntityModel, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(traceability.PlantEntityModel) traceability.PlantEntityModel); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(traceability.PlantEntityModel)
	}

	if rf, ok := ret.Get(1).(func(traceability.PlantEntityModel) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOuranosRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewOuranosRepository creates a new instance of OuranosRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOuranosRepository(t mockConstructorTestingTNewOuranosRepository) *OuranosRepository {
	mock := &OuranosRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
