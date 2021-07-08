package cluster

import (
	"context"
	"github.com/ovrclk/akash/provider/event"
	"github.com/ovrclk/akash/pubsub"
	dtypes "github.com/ovrclk/akash/x/deployment/types"
	mtypes "github.com/ovrclk/akash/x/market/types"
	"time"
)

type transferHostnamesRequest struct {
	hostnames []string
	destination dtypes.DeploymentID
}

func (s *service) transferHostnames(req transferHostnamesRequest) {
	ownerAddr, err := req.destination.GetOwnerAddress()
	if err != nil {
		return
	}

	var dst *deploymentManager
	toPurge := make(map[mtypes.LeaseID]*deploymentManager)
	for managerLeaseID, mgr := range s.managers {
		leaseOwnerAddr, err := managerLeaseID.DeploymentID().GetOwnerAddress()
		if err != nil {
			return
		}
		if req.destination.Equals(managerLeaseID.DeploymentID()) {
			dst = mgr
		} else if ownerAddr.Equals(leaseOwnerAddr) {
			toPurge[managerLeaseID] = mgr
		}
	}

	log := s.log.With("owner", ownerAddr.String())

	// Check that the destination is running
	if dst == nil {
		log.Error("destination deployment does not exist", "dseq", req.destination.DSeq)
		return
	}

	// TODO - use an exclusive lock to prevent double running this
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <- s.lc.ShuttingDown():
			cancel()
		case <- ctx.Done():
			return
		}
	}()

	var ready chan struct{}
	if len(toPurge) != 0 {
		sub, err := s.bus.Subscribe()
		if err != nil {
			cancel()
			return
		}
		// Wait for each one to complete
		toCheck := make(chan mtypes.LeaseID, len(toPurge))
		go func() {
			defer sub.Close()
			// TODO - query the cluster to determine what deployments need to
			// actually change
			for k, mgr := range toPurge {
				err = mgr.force()
				if err != nil {
					log.Error("failed forcing purge", "dseq", k.DSeq)
					cancel()
					return
				}

			}

			for {
				sent := make(map[mtypes.LeaseID]struct{}, len(toPurge))
				var ev pubsub.Event
				select {
				case ev = <-sub.Events():

				case <-ctx.Done():
					return
				}

				deploymentEvent, ok := ev.(event.ClusterDeployment)
				if !ok {
					continue
				}

				if deploymentEvent.Status != event.ClusterDeploymentUpdated {
					continue
				}

				_, isWaitingOn := toPurge[deploymentEvent.LeaseID]
				if !isWaitingOn {
					continue
				}

				_, alreadySent := sent[deploymentEvent.LeaseID]
				if !alreadySent {
					log.Info("Got deployment update event for", "lease", deploymentEvent.LeaseID)
					toCheck <- deploymentEvent.LeaseID
					sent[deploymentEvent.LeaseID] = struct{}{}
				}
			}
		}()

		ready = make(chan struct{}, 1)
		go func () {
			checking := make([]mtypes.LeaseID, 0, len(toPurge))
			releasedHostnames := make(map[mtypes.LeaseID]struct{})
			polling := time.NewTicker(10 * time.Second)
			for {
				select {
				case lID := <-toCheck:
					checking = append(checking, lID)
					continue
				case <-polling.C:

				case <-ctx.Done():
					return
				}

				stillChecking := false
				for _, lID := range checking {
					_, done := releasedHostnames[lID]
					if done {
						continue
					}

					log.Info("Checking hostnames of", "lease", lID)
					hostnames, err := s.client.LeaseHostnames(ctx, lID)
					if err != nil {
						log.Error("failed checking hostnames of", "lease", lID, "err", err)
						cancel()
						return
					}

					inUse := checkForStringIntersection(req.hostnames, hostnames)
					if inUse {
						stillChecking = true
					} else {
						releasedHostnames[lID] = struct{}{}
					}

				}

				if !stillChecking {
					break
				}
			}
			select {
				case ready <- struct{}{}:
				default:
			}
		}()
	}

	go func (){
		if ready != nil {
			// Wait for everything to be ready or for something to fail
			select {
			case <-ready:
			case <-ctx.Done():
				return
			}
		}

		cancel() // Cancel everything
		log.Info("running final update to takeover hostnames")
		// Run the destination so it can pickup the hostnames
		err = dst.force()
		if err != nil {
			log.Error("failed final update")
		}
	}()
}

func checkForStringIntersection(lhs, rhs []string) bool{
	m := make(map[string]struct{})
	for _, v := range lhs {
		m[v] = struct{}{}
	}

	for _, v := range rhs {
		_, ok := m[v]
		if ok {
			return true
		}
	}

	return false
}