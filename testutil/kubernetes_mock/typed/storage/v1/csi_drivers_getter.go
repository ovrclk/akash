// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

// CSIDriversGetter is an autogenerated mock type for the CSIDriversGetter type
type CSIDriversGetter struct {
	mock.Mock
}

// CSIDrivers provides a mock function with given fields:
func (_m *CSIDriversGetter) CSIDrivers() v1.CSIDriverInterface {
	ret := _m.Called()

	var r0 v1.CSIDriverInterface
	if rf, ok := ret.Get(0).(func() v1.CSIDriverInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.CSIDriverInterface)
		}
	}

	return r0
}
