// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
)

// RuntimeClassesGetter is an autogenerated mock type for the RuntimeClassesGetter type
type RuntimeClassesGetter struct {
	mock.Mock
}

// RuntimeClasses provides a mock function with given fields:
func (_m *RuntimeClassesGetter) RuntimeClasses() v1beta1.RuntimeClassInterface {
	ret := _m.Called()

	var r0 v1beta1.RuntimeClassInterface
	if rf, ok := ret.Get(0).(func() v1beta1.RuntimeClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.RuntimeClassInterface)
		}
	}

	return r0
}
