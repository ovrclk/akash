// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
)

// CertificateSigningRequestsGetter is an autogenerated mock type for the CertificateSigningRequestsGetter type
type CertificateSigningRequestsGetter struct {
	mock.Mock
}

// CertificateSigningRequests provides a mock function with given fields:
func (_m *CertificateSigningRequestsGetter) CertificateSigningRequests() v1beta1.CertificateSigningRequestInterface {
	ret := _m.Called()

	var r0 v1beta1.CertificateSigningRequestInterface
	if rf, ok := ret.Get(0).(func() v1beta1.CertificateSigningRequestInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.CertificateSigningRequestInterface)
		}
	}

	return r0
}
