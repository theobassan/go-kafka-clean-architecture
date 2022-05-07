package translator

import (
	"go-kafka-clean-architecture/app_func/entities"

	"github.com/stretchr/testify/mock"
)

type ProductTranslatorMock struct {
	mock.Mock
}

func (m ProductTranslatorMock) Translate(product entities.Product) entities.Product {
	ret := m.Called(product)

	var r0 entities.Product
	if rf, ok := ret.Get(0).(func(entities.Product) entities.Product); ok {
		r0 = rf(product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entities.Product)
		}
	}

	return r0
}
