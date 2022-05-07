package gateway

import (
	"go-kafka-clean-architecture/app_func/entities"

	"github.com/stretchr/testify/mock"
)

type ProductFinderGatewayMock struct {
	mock.Mock
}

func (m ProductFinderGatewayMock) FindById(id int64) (entities.Product, error) {
	ret := m.Called(id)

	var r0 entities.Product
	if rf, ok := ret.Get(0).(func(int64) entities.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entities.Product)
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
