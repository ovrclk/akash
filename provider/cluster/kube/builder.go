package kube

// nolint:deadcode,golint

import (
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"strconv"
	"strings"

	"github.com/tendermint/tendermint/libs/log"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/ovrclk/akash/manifest"
	mtypes "github.com/ovrclk/akash/x/market/types"
)

const (
	akashManagedLabelName         = "akash.network"
	akashNetworkNamespace         = "akash.network/namespace"
	akashManifestServiceLabelName = "akash.network/manifest-service"

	akashLeaseOwnerLabelName    = "akash.network/lease.id.owner"
	akashLeaseDSeqLabelName     = "akash.network/lease.id.dseq"
	akashLeaseGSeqLabelName     = "akash.network/lease.id.gseq"
	akashLeaseOSeqLabelName     = "akash.network/lease.id.oseq"
	akashLeaseProviderLabelName = "akash.network/lease.id.provider"

	akashDeploymentPolicyName = "akash-deployment-restrictions"
)

var (
	dnsPort     = intstr.FromInt(53)
	dnsProtocol = corev1.Protocol("UDP")
)

type builder struct {
	log      log.Logger
	settings Settings
	lid      mtypes.LeaseID
	group    *manifest.Group
}

func (b *builder) ns() string {
	return lidNS(b.lid)
}

func (b *builder) labels() map[string]string {
	return map[string]string{
		akashManagedLabelName: "true",
		akashNetworkNamespace: lidNS(b.lid),
	}
}

// TODO: re-enable.  see #946
// pspRestrictedBuilder produces restrictive PodSecurityPolicies for tenant Namespaces.
// Restricted PSP source: https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/policy/restricted-psp.yaml
// type pspRestrictedBuilder struct {
// 	builder
// }
//
// func newPspBuilder(settings Settings, lid mtypes.LeaseID, group *manifest.Group) *pspRestrictedBuilder { // nolint:golint,unparam
// 	return &pspRestrictedBuilder{builder: builder{settings: settings, lid: lid, group: group}}
// }
//
// func (p *pspRestrictedBuilder) name() string {
// 	return p.ns()
// }
//
// func (p *pspRestrictedBuilder) create() (*v1beta1.PodSecurityPolicy, error) { // nolint:golint,unparam
// 	falseVal := false
// 	return &v1beta1.PodSecurityPolicy{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      p.name(),
// 			Namespace: p.name(),
// 			Labels:    p.labels(),
// 			Annotations: map[string]string{
// 				"seccomp.security.alpha.kubernetes.io/allowedProfileNames": "docker/default,runtime/default",
// 				"apparmor.security.beta.kubernetes.io/allowedProfileNames": "runtime/default",
// 				"seccomp.security.alpha.kubernetes.io/defaultProfileName":  "runtime/default",
// 				"apparmor.security.beta.kubernetes.io/defaultProfileName":  "runtime/default",
// 			},
// 		},
// 		Spec: v1beta1.PodSecurityPolicySpec{
// 			Privileged:               false,
// 			AllowPrivilegeEscalation: &falseVal,
// 			RequiredDropCapabilities: []corev1.Capability{
// 				"ALL",
// 			},
// 			Volumes: []v1beta1.FSType{
// 				v1beta1.EmptyDir,
// 				v1beta1.PersistentVolumeClaim, // evaluate necessity later
// 			},
// 			HostNetwork: false,
// 			HostIPC:     false,
// 			HostPID:     false,
// 			RunAsUser: v1beta1.RunAsUserStrategyOptions{
// 				// fixme(#946): previous value RunAsUserStrategyMustRunAsNonRoot was interfering with
// 				// (b *deploymentBuilder) create() RunAsNonRoot: false
// 				// allow any user at this moment till revise all security debris of kube api
// 				Rule: v1beta1.RunAsUserStrategyRunAsAny,
// 			},
// 			SELinux: v1beta1.SELinuxStrategyOptions{
// 				Rule: v1beta1.SELinuxStrategyRunAsAny,
// 			},
// 			SupplementalGroups: v1beta1.SupplementalGroupsStrategyOptions{
// 				Rule: v1beta1.SupplementalGroupsStrategyRunAsAny,
// 			},
// 			FSGroup: v1beta1.FSGroupStrategyOptions{
// 				Rule: v1beta1.FSGroupStrategyMustRunAs,
// 				Ranges: []v1beta1.IDRange{
// 					{
// 						Min: int64(1),
// 						Max: int64(65535),
// 					},
// 				},
// 			},
// 			ReadOnlyRootFilesystem: false,
// 		},
// 	}, nil
// }
//
// func (p *pspRestrictedBuilder) update(obj *v1beta1.PodSecurityPolicy) (*v1beta1.PodSecurityPolicy, error) { // nolint:golint,unparam
// 	obj.Name = p.ns()
// 	obj.Labels = p.labels()
// 	return obj, nil
// }

const runtimeClassNoneValue = "none"

const (
	envVarAkashGroupSequence         = "AKASH_GROUP_SEQUENCE"
	envVarAkashDeploymentSequence    = "AKASH_DEPLOYMENT_SEQUENCE"
	envVarAkashOrderSequence         = "AKASH_ORDER_SEQUENCE"
	envVarAkashOwner                 = "AKASH_OWNER"
	envVarAkashProvider              = "AKASH_PROVIDER"
	envVarAkashClusterPublicHostname = "AKASH_CLUSTER_PUBLIC_HOSTNAME"
)

func addIfNotPresent(envVarsAlreadyAdded map[string]int, env []corev1.EnvVar, key string, value interface{}) []corev1.EnvVar {
	_, exists := envVarsAlreadyAdded[key]
	if exists {
		return env
	}

	env = append(env, corev1.EnvVar{Name: key, Value: fmt.Sprintf("%v", value)})
	return env
}

const suffixForNodePortServiceName = "-np"

func makeGlobalServiceNameFromBasename(basename string) string {
	return fmt.Sprintf("%s%s", basename, suffixForNodePortServiceName)
}

// lidNS generates a unique sha256 sum for identifying a provider's object name.
func lidNS(lid mtypes.LeaseID) string {
	path := lid.String()
	// DNS-1123 label must consist of lower case alphanumeric characters or '-',
	// and must start and end with an alphanumeric character
	// (e.g. 'my-name',  or '123-abc', regex used for validation
	// is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')
	sha := sha256.Sum224([]byte(path))
	return strings.ToLower(base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(sha[:]))
}

func appendLeaseLabels(lid mtypes.LeaseID, labels map[string]string) map[string]string {
	labels[akashLeaseOwnerLabelName] = lid.Owner
	labels[akashLeaseDSeqLabelName] = strconv.FormatUint(lid.DSeq, 10)
	labels[akashLeaseGSeqLabelName] = strconv.FormatUint(uint64(lid.GSeq), 10)
	labels[akashLeaseOSeqLabelName] = strconv.FormatUint(uint64(lid.OSeq), 10)
	labels[akashLeaseProviderLabelName] = lid.Provider
	return labels
}
