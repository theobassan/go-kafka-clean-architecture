package usecases

import (
	"go-kafka-clean-architecture/app_func/entities"

	"github.com/stretchr/testify/mock"
)

type ProductInteractorMock struct {
	mock.Mock
}

func (m ProductInteractorMock) Create(id int64) (int64, error) {
	ret := m.Called(id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m ProductInteractorMock) FindAll() ([]entities.Product, error) {
	ret := m.Called()

	var r0 []entities.Product
	if rf, ok := ret.Get(0).(func() []entities.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Product)
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

func (m ProductInteractorMock) Get(productID int64) (entities.Product, error) {
	ret := m.Called(productID)

	var r0 entities.Product
	if rf, ok := ret.Get(0).(func(int64) entities.Product); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entities.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m ProductInteractorMock) CreateTranslated(product entities.Product) (int64, error) {
	ret := m.Called(product)

	var r0 int64
	if rf, ok := ret.Get(0).(func(entities.Product) int64); ok {
		r0 = rf(product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Product) error); ok {
		r1 = rf(product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m ProductInteractorMock) FindAllTranslated() ([]entities.Product, error) {
	ret := m.Called()

	var r0 []entities.Product
	if rf, ok := ret.Get(0).(func() []entities.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Product)
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
