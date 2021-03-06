// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	payments "RentSpace/businesses/payments"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddPayment provides a mock function with given fields: _a0, paymentDomain
func (_m *Repository) AddPayment(_a0 context.Context, paymentDomain *payments.Domain) error {
	ret := _m.Called(_a0, paymentDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *payments.Domain) error); ok {
		r0 = rf(_a0, paymentDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
