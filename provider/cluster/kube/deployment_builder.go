package kube

import (
	"github.com/tendermint/tendermint/libs/log"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ovrclk/akash/manifest"
	mtypes "github.com/ovrclk/akash/x/market/types"
)

// deployment
type deploymentBuilder struct {
	workloadBuilder
}

func newDeploymentBuilder(log log.Logger, settings Settings, lid mtypes.LeaseID, group *manifest.Group, service *manifest.Service) *deploymentBuilder {
	return &deploymentBuilder{
		workloadBuilder: newWorkloadBuilder(log, settings, lid, group, service),
	}
}

func (b *deploymentBuilder) create() (*appsv1.Deployment, error) { // nolint:golint,unparam
	replicas := int32(b.service.Count)
	falseValue := false

	var effectiveRuntimeClassName *string
	if len(b.runtimeClassName) != 0 && b.runtimeClassName != runtimeClassNoneValue {
		effectiveRuntimeClassName = &b.runtimeClassName
	}

	kdeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:   b.name(),
			Labels: b.labels(),
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: b.labels(),
			},
			Replicas: &replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: b.labels(),
				},
				Spec: corev1.PodSpec{
					RuntimeClassName: effectiveRuntimeClassName,
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: &falseValue,
					},
					AutomountServiceAccountToken: &falseValue,
					Containers:                   []corev1.Container{b.container()},
				},
			},
		},
	}

	return kdeployment, nil
}

func (b *deploymentBuilder) update(obj *appsv1.Deployment) (*appsv1.Deployment, error) { // nolint:golint,unparam
	replicas := int32(b.service.Count)
	obj.Labels = b.labels()
	obj.Spec.Selector.MatchLabels = b.labels()
	obj.Spec.Replicas = &replicas
	obj.Spec.Template.Labels = b.labels()
	obj.Spec.Template.Spec.Containers = []corev1.Container{b.container()}

	return obj, nil
}
