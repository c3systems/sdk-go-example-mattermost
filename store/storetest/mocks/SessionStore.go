// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/c3systems/c3-sdk-go-example-mattermost/model"
import store "github.com/c3systems/c3-sdk-go-example-mattermost/store"

// SessionStore is an autogenerated mock type for the SessionStore type
type SessionStore struct {
	mock.Mock
}

// AnalyticsSessionCount provides a mock function with given fields:
func (_m *SessionStore) AnalyticsSessionCount() store.StoreChannel {
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

// Cleanup provides a mock function with given fields: expiryTime, batchSize
func (_m *SessionStore) Cleanup(expiryTime int64, batchSize int64) {
	_m.Called(expiryTime, batchSize)
}

// Get provides a mock function with given fields: sessionIdOrToken
func (_m *SessionStore) Get(sessionIdOrToken string) store.StoreChannel {
	ret := _m.Called(sessionIdOrToken)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(sessionIdOrToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetSessions provides a mock function with given fields: userId
func (_m *SessionStore) GetSessions(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetSessionsWithActiveDeviceIds provides a mock function with given fields: userId
func (_m *SessionStore) GetSessionsWithActiveDeviceIds(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteSessionsByUser provides a mock function with given fields: teamId
func (_m *SessionStore) PermanentDeleteSessionsByUser(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Remove provides a mock function with given fields: sessionIdOrToken
func (_m *SessionStore) Remove(sessionIdOrToken string) store.StoreChannel {
	ret := _m.Called(sessionIdOrToken)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(sessionIdOrToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// RemoveAllSessions provides a mock function with given fields:
func (_m *SessionStore) RemoveAllSessions() store.StoreChannel {
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

// Save provides a mock function with given fields: session
func (_m *SessionStore) Save(session *model.Session) store.StoreChannel {
	ret := _m.Called(session)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Session) store.StoreChannel); ok {
		r0 = rf(session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateDeviceId provides a mock function with given fields: id, deviceId, expiresAt
func (_m *SessionStore) UpdateDeviceId(id string, deviceId string, expiresAt int64) store.StoreChannel {
	ret := _m.Called(id, deviceId, expiresAt)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, int64) store.StoreChannel); ok {
		r0 = rf(id, deviceId, expiresAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateLastActivityAt provides a mock function with given fields: sessionId, time
func (_m *SessionStore) UpdateLastActivityAt(sessionId string, time int64) store.StoreChannel {
	ret := _m.Called(sessionId, time)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64) store.StoreChannel); ok {
		r0 = rf(sessionId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateRoles provides a mock function with given fields: userId, roles
func (_m *SessionStore) UpdateRoles(userId string, roles string) store.StoreChannel {
	ret := _m.Called(userId, roles)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(userId, roles)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
