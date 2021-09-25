// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	customers "RentSpace/businesses/customers"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddCustomer provides a mock function with given fields: ctx, customer
func (_m *Repository) AddCustomer(ctx context.Context, customer *customers.Domain) error {
	ret := _m.Called(ctx, customer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *customers.Domain) error); ok {
		r0 = rf(ctx, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCustomer provides a mock function with given fields: ctx
func (_m *Repository) GetAllCustomer(ctx context.Context) ([]customers.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []customers.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []customers.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]customers.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) GetByEmail(ctx context.Context, email string) (customers.Domain, error) {
	ret := _m.Called(ctx, email)

	var r0 customers.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) customers.Domain); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(customers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidationEmailPassword provides a mock function with given fields: ctx, email, password
func (_m *Repository) ValidationEmailPassword(ctx context.Context, email string, password string) (customers.Domain, error) {
	ret := _m.Called(ctx, email, password)

	var r0 customers.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) customers.Domain); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(customers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
