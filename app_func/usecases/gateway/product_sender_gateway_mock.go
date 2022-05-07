package gateway

import (
	"go-kafka-clean-architecture/app_func/entities"

	"github.com/stretchr/testify/mock"
)

type ProductSenderGatewayMock struct {
	mock.Mock
}

func (m ProductSenderGatewayMock) Send(product entities.Product) error {
	ret := m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
