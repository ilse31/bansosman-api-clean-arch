// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	users "bansosman/bussiness/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: user, id
func (_m *Repository) Delete(user *users.Domain, id int) (string, error) {
	ret := _m.Called(user, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(*users.Domain, int) string); ok {
		r0 = rf(user, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain, int) error); ok {
		r1 = rf(user, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *Repository) FindAll() ([]users.Domain, error) {
	ret := _m.Called()

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func() []users.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
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

// FindByID provides a mock function with given fields: id
func (_m *Repository) FindByID(id int) (*users.Domain, error) {
	ret := _m.Called(id)

	var r0 *users.Domain
	if rf, ok := ret.Get(0).(func(int) *users.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *Repository) GetByName(name string) (users.Domain, error) {
	ret := _m.Called(name)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(string) users.Domain); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: user
func (_m *Repository) Insert(user *users.Domain) (*users.Domain, error) {
	ret := _m.Called(user)

	var r0 *users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) *users.Domain); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: user, id
func (_m *Repository) Update(user *users.Domain, id int) (*users.Domain, error) {
	ret := _m.Called(user, id)

	var r0 *users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain, int) *users.Domain); ok {
		r0 = rf(user, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain, int) error); ok {
		r1 = rf(user, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
