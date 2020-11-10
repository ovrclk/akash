// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/api/networking/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// IngressClassInterface is an autogenerated mock type for the IngressClassInterface type
type IngressClassInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, ingressClass, opts
func (_m *IngressClassInterface) Create(ctx context.Context, ingressClass *v1.IngressClass, opts metav1.CreateOptions) (*v1.IngressClass, error) {
	ret := _m.Called(ctx, ingressClass, opts)

	var r0 *v1.IngressClass
	if rf, ok := ret.Get(0).(func(context.Context, *v1.IngressClass, metav1.CreateOptions) *v1.IngressClass); ok {
		r0 = rf(ctx, ingressClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.IngressClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.IngressClass, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, ingressClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *IngressClassInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *IngressClassInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *IngressClassInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.IngressClass, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *v1.IngressClass
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *v1.IngressClass); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.IngressClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *IngressClassInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1.IngressClassList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *v1.IngressClassList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *v1.IngressClassList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.IngressClassList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *IngressClassInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1.IngressClass, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.IngressClass
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *v1.IngressClass); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.IngressClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ingressClass, opts
func (_m *IngressClassInterface) Update(ctx context.Context, ingressClass *v1.IngressClass, opts metav1.UpdateOptions) (*v1.IngressClass, error) {
	ret := _m.Called(ctx, ingressClass, opts)

	var r0 *v1.IngressClass
	if rf, ok := ret.Get(0).(func(context.Context, *v1.IngressClass, metav1.UpdateOptions) *v1.IngressClass); ok {
		r0 = rf(ctx, ingressClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.IngressClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.IngressClass, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, ingressClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *IngressClassInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
