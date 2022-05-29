// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// Read provides a mock function with given fields: startPos, buf
func (_m *Reader) Read(startPos uint, buf []byte) (int, error) {
	ret := _m.Called(startPos, buf)

	var r0 int
	if rf, ok := ret.Get(0).(func(uint, []byte) int); ok {
		r0 = rf(startPos, buf)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, []byte) error); ok {
		r1 = rf(startPos, buf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewReaderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReader(t NewReaderT) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}