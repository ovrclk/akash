package kube

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ovrclk/akash/manifest"
	mtypes "github.com/ovrclk/akash/x/market/types"
)

type nsBuilder struct {
	builder
}

func newNSBuilder(settings Settings, lid mtypes.LeaseID, group *manifest.Group) *nsBuilder {
	return &nsBuilder{builder: builder{settings: settings, lid: lid, group: group}}
}

func (b *nsBuilder) name() string {
	return b.ns()
}

func (b *nsBuilder) labels() map[string]string {
	return appendLeaseLabels(b.lid, b.builder.labels())
}

func (b *nsBuilder) create() (*corev1.Namespace, error) { // nolint:golint,unparam
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   b.ns(),
			Labels: b.labels(),
		},
	}, nil
}

func (b *nsBuilder) update(obj *corev1.Namespace) (*corev1.Namespace, error) { // nolint:golint,unparam
	obj.Name = b.ns()
	obj.Labels = b.labels()
	return obj, nil
}
