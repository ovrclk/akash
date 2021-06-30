package kube

import (
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ovrclk/akash/manifest"
	akashv1 "github.com/ovrclk/akash/pkg/apis/akash.network/v1"
	mtypes "github.com/ovrclk/akash/x/market/types"
)

// manifestBuilder composes the k8s akashv1.Manifest type from LeaseID and
// manifest.Group data.
type manifestBuilder struct {
	builder
	mns string // Q: is this supposed to be the k8s Namespace? It's the Object name now.
}

func newManifestBuilder(log log.Logger, settings Settings, ns string, lid mtypes.LeaseID, group *manifest.Group) *manifestBuilder {
	return &manifestBuilder{
		builder: builder{
			log:      log.With("module", "kube-builder"),
			settings: settings,
			lid:      lid,
			group:    group,
		},
		mns: ns,
	}
}

func (b *manifestBuilder) labels() map[string]string {
	return appendLeaseLabels(b.lid, b.builder.labels())
}

func (b *manifestBuilder) ns() string {
	return b.mns
}

func (b *manifestBuilder) create() (*akashv1.Manifest, error) {
	obj, err := akashv1.NewManifest(b.name(), b.lid, b.group)
	if err != nil {
		return nil, err
	}
	obj.Labels = b.labels()
	return obj, nil
}

func (b *manifestBuilder) update(obj *akashv1.Manifest) (*akashv1.Manifest, error) {
	m, err := akashv1.NewManifest(b.name(), b.lid, b.group)
	if err != nil {
		return nil, err
	}
	obj.Spec = m.Spec
	obj.Labels = b.labels()
	return obj, nil
}

func (b *manifestBuilder) name() string {
	return lidNS(b.lid)
}
