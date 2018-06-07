package state

import (
	"github.com/gogo/protobuf/proto"
	"github.com/ovrclk/akash/types"
	"github.com/ovrclk/akash/types/base"
	"github.com/tendermint/iavl"
)

type DeploymentAdapter interface {
	Save(deployment *types.Deployment) error
	Get(base.Bytes) (*types.Deployment, error)
	GetMaxRange() (*types.Deployments, error)
	GetRangeWithProof(base.Bytes, base.Bytes, int) ([][]byte, *types.Deployments, iavl.KeyRangeProof, error)
	KeyFor(base.Bytes) base.Bytes
	SequenceFor(base.Bytes) Sequence
}

type deploymentAdapter struct {
	db DB
}

func NewDeploymentAdapter(db DB) DeploymentAdapter {
	return &deploymentAdapter{db}
}

func (d *deploymentAdapter) Save(deployment *types.Deployment) error {
	key := d.KeyFor(deployment.Address)

	dbytes, err := proto.Marshal(deployment)
	if err != nil {
		return err
	}

	d.db.Set(key, dbytes)
	return nil
}

func (d *deploymentAdapter) Get(address base.Bytes) (*types.Deployment, error) {

	dep := types.Deployment{}

	key := d.KeyFor(address)

	buf := d.db.Get(key)
	if buf == nil {
		return nil, nil
	}

	dep.Unmarshal(buf)

	return &dep, nil
}

func (d *deploymentAdapter) GetMaxRange() (*types.Deployments, error) {
	_, deps, _, err := d.GetRangeWithProof(MinAddress(), MaxAddress(), MaxRangeLimit)
	return deps, err
}

func (d *deploymentAdapter) GetRangeWithProof(startKey base.Bytes, endKey base.Bytes, limit int) ([][]byte, *types.Deployments, iavl.KeyRangeProof, error) {
	deps := types.Deployments{}
	proof := iavl.KeyRangeProof{}

	start := d.KeyFor(startKey)
	end := d.KeyFor(endKey)

	keys, dbytes, proof, err := d.db.GetRangeWithProof(start, end, limit)
	if err != nil {
		return nil, &deps, proof, err
	}
	if keys == nil {
		return nil, &deps, proof, nil
	}

	for _, d := range dbytes {
		dep := types.Deployment{}
		dep.Unmarshal(d)
		deps.Items = append(deps.Items, dep)
	}

	return keys, &deps, proof, nil
}

func (a *deploymentAdapter) KeyFor(address base.Bytes) base.Bytes {
	return append([]byte(DeploymentPath), address...)
}

func (a *deploymentAdapter) SequenceFor(address base.Bytes) Sequence {
	path := append([]byte(DeploymentSequencePath), address...)
	return NewSequence(a.db, path)
}
