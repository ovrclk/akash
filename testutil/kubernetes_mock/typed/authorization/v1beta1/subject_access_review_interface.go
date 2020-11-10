// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/api/authorization/v1beta1"
)

// SubjectAccessReviewInterface is an autogenerated mock type for the SubjectAccessReviewInterface type
type SubjectAccessReviewInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, subjectAccessReview, opts
func (_m *SubjectAccessReviewInterface) Create(ctx context.Context, subjectAccessReview *v1beta1.SubjectAccessReview, opts v1.CreateOptions) (*v1beta1.SubjectAccessReview, error) {
	ret := _m.Called(ctx, subjectAccessReview, opts)

	var r0 *v1beta1.SubjectAccessReview
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.SubjectAccessReview, v1.CreateOptions) *v1beta1.SubjectAccessReview); ok {
		r0 = rf(ctx, subjectAccessReview, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.SubjectAccessReview)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.SubjectAccessReview, v1.CreateOptions) error); ok {
		r1 = rf(ctx, subjectAccessReview, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
