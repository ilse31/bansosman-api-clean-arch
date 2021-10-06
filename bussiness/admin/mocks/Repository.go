// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	admin "bansosman/bussiness/admin"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id uint) (admin.Domain, error) {
	ret := _m.Called(id)

	var r0 admin.Domain
	if rf, ok := ret.Get(0).(func(uint) admin.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(admin.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *Repository) GetByUsername(username string) (admin.Domain, error) {
	ret := _m.Called(username)

	var r0 admin.Domain
	if rf, ok := ret.Get(0).(func(string) admin.Domain); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(admin.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: userData
func (_m *Repository) Register(userData *admin.Domain) (admin.Domain, error) {
	ret := _m.Called(userData)

	var r0 admin.Domain
	if rf, ok := ret.Get(0).(func(*admin.Domain) admin.Domain); ok {
		r0 = rf(userData)
	} else {
		r0 = ret.Get(0).(admin.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*admin.Domain) error); ok {
		r1 = rf(userData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}