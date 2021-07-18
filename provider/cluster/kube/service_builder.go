package kube

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/ovrclk/akash/manifest"
	"github.com/ovrclk/akash/provider/cluster/util"
	mtypes "github.com/ovrclk/akash/x/market/types"
)

// service
type serviceBuilder struct {
	workloadBuilder
	requireNodePort bool
}

func newServiceBuilder(log log.Logger, settings Settings, lid mtypes.LeaseID, group *manifest.Group, service *manifest.Service, requireNodePort bool) *serviceBuilder {
	return &serviceBuilder{
		workloadBuilder: newWorkloadBuilder(log, settings, lid, group, service),
		requireNodePort: requireNodePort,
	}
}

func (b *serviceBuilder) name() string {
	basename := b.workloadBuilder.name()
	if b.requireNodePort {
		return makeGlobalServiceNameFromBasename(basename)
	}
	return basename
}

func (b *serviceBuilder) workloadServiceType() corev1.ServiceType {
	if b.requireNodePort {
		return corev1.ServiceTypeNodePort
	}
	return corev1.ServiceTypeClusterIP
}

func (b *serviceBuilder) create() (*corev1.Service, error) { // nolint:golint,unparam
	ports, err := b.ports()
	if err != nil {
		return nil, err
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:   b.name(),
			Labels: b.labels(),
		},
		Spec: corev1.ServiceSpec{
			Type:     b.workloadServiceType(),
			Selector: b.labels(),
			Ports:    ports,
		},
	}
	b.log.Debug("provider/cluster/kube/builder: created service", "service", service)

	return service, nil
}

func (b *serviceBuilder) update(obj *corev1.Service) (*corev1.Service, error) { // nolint:golint,unparam
	obj.Labels = b.labels()
	obj.Spec.Selector = b.labels()
	ports, err := b.ports()
	if err != nil {
		return nil, err
	}

	// retain provisioned NodePort values
	if b.requireNodePort {

		// for each newly-calculated port
		for i, port := range ports {

			// if there is a current (in-kube) port defined
			// with the same specified values
			for _, curport := range obj.Spec.Ports {
				if curport.Name == port.Name &&
					curport.Port == port.Port &&
					curport.TargetPort.IntValue() == port.TargetPort.IntValue() &&
					curport.Protocol == port.Protocol {

					// re-use current port
					ports[i] = curport
				}
			}
		}
	}

	obj.Spec.Ports = ports
	return obj, nil
}

func (b *serviceBuilder) any() bool {
	for _, expose := range b.service.Expose {
		exposeIsIngress := util.ShouldBeIngress(expose)
		if b.requireNodePort && exposeIsIngress {
			continue
		}

		if !b.requireNodePort && exposeIsIngress {
			return true
		}

		if expose.Global == b.requireNodePort {
			return true
		}
	}
	return false
}

var errUnsupportedProtocol = errors.New("Unsupported protocol for service")
var errInvalidServiceBuilder = errors.New("service builder invalid")

func (b *serviceBuilder) ports() ([]corev1.ServicePort, error) {
	ports := make([]corev1.ServicePort, 0, len(b.service.Expose))
	for i, expose := range b.service.Expose {
		shouldBeIngress := util.ShouldBeIngress(expose)
		if expose.Global == b.requireNodePort || (!b.requireNodePort && shouldBeIngress) {
			if b.requireNodePort && shouldBeIngress {
				continue
			}

			var exposeProtocol corev1.Protocol
			switch expose.Proto {
			case manifest.TCP:
				exposeProtocol = corev1.ProtocolTCP
			case manifest.UDP:
				exposeProtocol = corev1.ProtocolUDP
			default:
				return nil, errUnsupportedProtocol
			}
			externalPort := util.ExposeExternalPort(b.service.Expose[i])
			ports = append(ports, corev1.ServicePort{
				Name:       fmt.Sprintf("%d-%d", i, int(externalPort)),
				Port:       externalPort,
				TargetPort: intstr.FromInt(int(expose.Port)),
				Protocol:   exposeProtocol,
			})
		}
	}

	if len(ports) == 0 {
		b.log.Debug("provider/cluster/kube/builder: created 0 ports", "requireNodePort", b.requireNodePort, "serviceExpose", b.service.Expose)
		return nil, errInvalidServiceBuilder
	}
	return ports, nil
}
