package query

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ovrclk/akash/x/deployment/types"
)

const (
	deploymentsPath = "deployments"
	deploymentPath  = "deployment"
	groupPath       = "group"
)

// getDeploymentsPath returns deployments path for queries
func getDeploymentsPath() string {
	return deploymentsPath
}

// DeploymentPath return deployment path of given deployment id for queries
func DeploymentPath(id types.DeploymentID) string {
	return fmt.Sprintf("%s/%s/%v", deploymentPath, id.Owner, id.DSeq)
}

// getGroupPath return group path of given group id for queries
func getGroupPath(id types.GroupID) string {
	return fmt.Sprintf("%s/%s/%v/%v", groupPath, id.Owner, id.DSeq, id.GSeq)
}

// parseDeploymentPath returns DeploymentID details with provided queries, and return
// error if occured due to wrong query
func parseDeploymentPath(parts []string) (types.DeploymentID, error) {
	if len(parts) < 2 {
		return types.DeploymentID{}, fmt.Errorf("invalid path")
	}

	owner, err := sdk.AccAddressFromBech32(parts[0])
	if err != nil {
		return types.DeploymentID{}, err
	}

	dseq, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return types.DeploymentID{}, err
	}

	return types.DeploymentID{
		Owner: owner,
		DSeq:  dseq,
	}, nil
}

// parseGroupPath returns GroupID details with provided queries, and return
// error if occured due to wrong query
func parseGroupPath(parts []string) (types.GroupID, error) {
	if len(parts) < 3 {
		return types.GroupID{}, fmt.Errorf("invalid path")
	}

	did, err := parseDeploymentPath(parts[0:2])
	if err != nil {
		return types.GroupID{}, err
	}

	gseq, err := strconv.ParseUint(parts[2], 10, 32)

	return types.MakeGroupID(did, uint32(gseq)), nil
}
