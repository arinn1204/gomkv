// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	io "github.com/arinn1204/gomkv/internal/io"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// EbmlReader is an autogenerated mock type for the EbmlReader type
type EbmlReader struct {
	mock.Mock
}

// Read provides a mock function with given fields: f, startPos, buf
func (_m *EbmlReader) Read(f *io.File, startPos uint, buf []byte) int {
	ret := _m.Called(f, startPos, buf)

	var r0 int
	if rf, ok := ret.Get(0).(func(*io.File, uint, []byte) int); ok {
		r0 = rf(f, startPos, buf)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NewEbmlReader creates a new instance of EbmlReader. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewEbmlReader(t testing.TB) *EbmlReader {
	mock := &EbmlReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}