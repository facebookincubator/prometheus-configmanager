// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	alert "prometheus-configmanager/prometheus/alert"

	config "prometheus-configmanager/alertmanager/config"

	mock "github.com/stretchr/testify/mock"
)

// AlertmanagerClient is an autogenerated mock type for the AlertmanagerClient type
type AlertmanagerClient struct {
	mock.Mock
}

// CreateReceiver provides a mock function with given fields: tenantID, rec
func (_m *AlertmanagerClient) CreateReceiver(tenantID string, rec config.Receiver) error {
	ret := _m.Called(tenantID, rec)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, config.Receiver) error); ok {
		r0 = rf(tenantID, rec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReceiver provides a mock function with given fields: tenantID, receiverName
func (_m *AlertmanagerClient) DeleteReceiver(tenantID string, receiverName string) error {
	ret := _m.Called(tenantID, receiverName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(tenantID, receiverName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGlobalConfig provides a mock function with given fields:
func (_m *AlertmanagerClient) GetGlobalConfig() (*config.GlobalConfig, error) {
	ret := _m.Called()

	var r0 *config.GlobalConfig
	if rf, ok := ret.Get(0).(func() *config.GlobalConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.GlobalConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReceivers provides a mock function with given fields: tenantID
func (_m *AlertmanagerClient) GetReceivers(tenantID string) ([]config.Receiver, error) {
	ret := _m.Called(tenantID)

	var r0 []config.Receiver
	if rf, ok := ret.Get(0).(func(string) []config.Receiver); ok {
		r0 = rf(tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]config.Receiver)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoute provides a mock function with given fields: tenantID
func (_m *AlertmanagerClient) GetRoute(tenantID string) (*config.Route, error) {
	ret := _m.Called(tenantID)

	var r0 *config.Route
	if rf, ok := ret.Get(0).(func(string) *config.Route); ok {
		r0 = rf(tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyTenantRoute provides a mock function with given fields: tenantID, route
func (_m *AlertmanagerClient) ModifyTenantRoute(tenantID string, route *config.Route) error {
	ret := _m.Called(tenantID, route)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *config.Route) error); ok {
		r0 = rf(tenantID, route)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReloadAlertmanager provides a mock function with given fields:
func (_m *AlertmanagerClient) ReloadAlertmanager() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGlobalConfig provides a mock function with given fields: globalConfig
func (_m *AlertmanagerClient) SetGlobalConfig(globalConfig config.GlobalConfig) error {
	ret := _m.Called(globalConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(config.GlobalConfig) error); ok {
		r0 = rf(globalConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tenancy provides a mock function with given fields:
func (_m *AlertmanagerClient) Tenancy() *alert.TenancyConfig {
	ret := _m.Called()

	var r0 *alert.TenancyConfig
	if rf, ok := ret.Get(0).(func() *alert.TenancyConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*alert.TenancyConfig)
		}
	}

	return r0
}

// UpdateReceiver provides a mock function with given fields: tenantID, receiverName, newRec
func (_m *AlertmanagerClient) UpdateReceiver(tenantID string, receiverName string, newRec *config.Receiver) error {
	ret := _m.Called(tenantID, receiverName, newRec)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *config.Receiver) error); ok {
		r0 = rf(tenantID, receiverName, newRec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
