// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "k8s.io/api/authorization/v1"
)

// LocalSubjectAccessReviewInterface is an autogenerated mock type for the LocalSubjectAccessReviewInterface type
type LocalSubjectAccessReviewInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, localSubjectAccessReview, opts
func (_m *LocalSubjectAccessReviewInterface) Create(ctx context.Context, localSubjectAccessReview *v1.LocalSubjectAccessReview, opts metav1.CreateOptions) (*v1.LocalSubjectAccessReview, error) {
	ret := _m.Called(ctx, localSubjectAccessReview, opts)

	var r0 *v1.LocalSubjectAccessReview
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LocalSubjectAccessReview, metav1.CreateOptions) *v1.LocalSubjectAccessReview); ok {
		r0 = rf(ctx, localSubjectAccessReview, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.LocalSubjectAccessReview)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.LocalSubjectAccessReview, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, localSubjectAccessReview, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
