// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/c3systems/c3-sdk-go-example-mattermost/model"
import store "github.com/c3systems/c3-sdk-go-example-mattermost/store"

// SchemeStore is an autogenerated mock type for the SchemeStore type
type SchemeStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: schemeId
func (_m *SchemeStore) Delete(schemeId string) store.StoreChannel {
	ret := _m.Called(schemeId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(schemeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: schemeId
func (_m *SchemeStore) Get(schemeId string) store.StoreChannel {
	ret := _m.Called(schemeId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(schemeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllPage provides a mock function with given fields: scope, offset, limit
func (_m *SchemeStore) GetAllPage(scope string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(scope, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(scope, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByName provides a mock function with given fields: schemeName
func (_m *SchemeStore) GetByName(schemeName string) store.StoreChannel {
	ret := _m.Called(schemeName)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(schemeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteAll provides a mock function with given fields:
func (_m *SchemeStore) PermanentDeleteAll() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: scheme
func (_m *SchemeStore) Save(scheme *model.Scheme) store.StoreChannel {
	ret := _m.Called(scheme)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Scheme) store.StoreChannel); ok {
		r0 = rf(scheme)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
