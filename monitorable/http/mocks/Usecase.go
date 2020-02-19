// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-monitorable

package mocks

import (
	monitorormodels "github.com/monitoror/monitoror/models"
	models "github.com/monitoror/monitoror/monitorable/http/models"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// HTTPFormatted provides a mock function with given fields: params
func (_m *Usecase) HTTPFormatted(params *models.HTTPFormattedParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.HTTPFormattedParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.HTTPFormattedParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HTTPProxy provides a mock function with given fields: params
func (_m *Usecase) HTTPProxy(params *models.HTTPProxyParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.HTTPProxyParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.HTTPProxyParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HTTPRaw provides a mock function with given fields: params
func (_m *Usecase) HTTPRaw(params *models.HTTPRawParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.HTTPRawParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.HTTPRawParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HTTPStatus provides a mock function with given fields: params
func (_m *Usecase) HTTPStatus(params *models.HTTPStatusParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.HTTPStatusParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.HTTPStatusParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
