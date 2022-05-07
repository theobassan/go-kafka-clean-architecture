package api

import (
	"github.com/stretchr/testify/mock"
)

type EventApiMock struct {
	mock.Mock
}

func (m *EventApiMock) Bind(topic string, value []byte) interface{} {
	ret := m.Called(topic, value)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, []byte) interface{}); ok {
		r0 = rf(topic, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0)
		}
	}

	return r0
}

func (m *EventApiMock) WriteMessage(i interface{}) error {
	ret := m.Called(i)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Error(0)
		}
	}

	return r0
}
