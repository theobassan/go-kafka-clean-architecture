package router

import (
	"github.com/stretchr/testify/mock"
)

type EventContextMock struct {
	mock.Mock
}

func (m *EventContextMock) Bind(v any) error {
	ret := m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(any) error); ok {
		r0 = rf(v)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}

	return r0
}

func (m *EventContextMock) Acknowledge() error {
	ret := m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}

	return r0
}
