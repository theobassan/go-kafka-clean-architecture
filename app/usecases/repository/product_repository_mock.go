package repository

import (
	"go-kafka-clean-architecture/app/entities"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) Create(u *entities.Product) (*int64, error) {
	ret := m.Called(u)

	var r0 *int64
	if rf, ok := ret.Get(0).(func(*entities.Product) *int64); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Product) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *ProductRepositoryMock) FindAll() ([]*entities.Product, error) {
	ret := m.Called()

	var r0 []*entities.Product
	if rf, ok := ret.Get(0).(func() []*entities.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Product)
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
