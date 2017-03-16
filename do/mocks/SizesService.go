// Generated: please do not edit by hand

package mocks

import do "github.com/digitalocean/doctl/do"
import mock "github.com/stretchr/testify/mock"

// SizesService is an autogenerated mock type for the SizesService type
type SizesService struct {
	mock.Mock
}

// List provides a mock function with given fields:
func (_m *SizesService) List() (do.Sizes, error) {
	ret := _m.Called()

	var r0 do.Sizes
	if rf, ok := ret.Get(0).(func() do.Sizes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.Sizes)
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
