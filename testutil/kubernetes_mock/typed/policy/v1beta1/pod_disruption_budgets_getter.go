// Code generated by mockery v1.1.2. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
)

// PodDisruptionBudgetsGetter is an autogenerated mock type for the PodDisruptionBudgetsGetter type
type PodDisruptionBudgetsGetter struct {
	mock.Mock
}

// PodDisruptionBudgets provides a mock function with given fields: namespace
func (_m *PodDisruptionBudgetsGetter) PodDisruptionBudgets(namespace string) v1beta1.PodDisruptionBudgetInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.PodDisruptionBudgetInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.PodDisruptionBudgetInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.PodDisruptionBudgetInterface)
		}
	}

	return r0
}