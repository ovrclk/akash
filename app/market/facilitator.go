package market

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ovrclk/photon/app/deploymentOrder"
	"github.com/ovrclk/photon/state"
	"github.com/ovrclk/photon/txutil"
	"github.com/ovrclk/photon/types/base"
	tmtypes "github.com/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/consensus/types"
	"github.com/tendermint/tendermint/rpc/core"
	tmtmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tmlibs/log"
)

const (
	subscriber = "photon-market"
)

type Facilitator interface {
	OnBeginBlock(req tmtypes.RequestBeginBlock) error
	OnCommit(state state.State) error
}

type facilitator struct {
	log       log.Logger
	validator tmtmtypes.PrivValidatorFS

	block *tmtypes.RequestBeginBlock
	rs    *ctypes.RoundState

	mtx *sync.Mutex
}

func (f *facilitator) OnBeginBlock(req tmtypes.RequestBeginBlock) error {
	f.block = &req
	return nil
}

func (f *facilitator) buildTx(signer tmtmtypes.PrivValidatorFS, nonce uint64, payload interface{}) ([]byte, error) {
	txb, err := txutil.NewTxBuilder(nonce, payload)
	if err != nil {
		return nil, err
	}
	sig, err := signer.Sign(txb.SignBytes())
	if err != nil {
		return nil, err
	}
	pubkey := signer.PubKey
	txb.Sign(pubkey, sig)
	return txb.TxBytes()
}

func (f *facilitator) sendTx(tx []byte) error {
	result, err := core.BroadcastTxCommit(tx)
	if err != nil {
		data, _ := json.MarshalIndent(err, "", "  ")
		println("err:\n" + string(data))
		f.log.Error("failed to send tx's", err)
		return err
	}
	data, _ := json.MarshalIndent(result, "", "  ")
	println("result:\n" + string(data))
	return nil
}

func (f *facilitator) getNonce(state state.State) (uint64, error) {
	address := new(base.Bytes)
	address.Unmarshal(f.validator.Address)
	account, err := state.Account().Get(*address)
	if err != nil {
		return 0, err
	}
	if account == nil {
		return uint64(1), nil
	}
	return account.Nonce, nil
}

func (f *facilitator) OnCommit(state state.State) error {
	if !f.checkCommit(state) {
		f.log.Info("NOT MY TURN TO DO MARKET STUFF.")
		return nil
	}
	f.log.Info("I SHOULD MAKE THE MARKET STUFF HAPPEN.")

	// create deployment order tx
	createDeploymentOrderTxs, err := deploymentOrder.CreateDeploymentOrderTxs(state)
	if err != nil {
		f.log.Error("Failed to generate createDeploymentOrder transactions", err)
	}

	nonce, err := f.getNonce(state)
	if err != nil {
		f.log.Error("failed to get validator nonce", err)
	}

	// todo: increment nonce for account
	for _, createDeploymentOrder := range createDeploymentOrderTxs {
		nonce += 1
		tx, err := f.buildTx(f.validator, nonce, createDeploymentOrder)
		if err != nil {
			f.log.Error("failed to build tx", err)
		}
		go f.sendTx(tx)
	}

	return nil
}

func (f *facilitator) checkCommit(state state.State) bool {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	if f.block == nil || f.rs == nil {
		return false
	}

	if uint64(f.rs.Height) != state.Version() {
		return false
	}

	if bytes.Compare(f.validator.GetAddress(), f.rs.Validators.GetProposer().PubKey.Address()) != 0 {
		return false
	}

	if bytes.Compare(f.block.Hash, f.rs.ProposalBlock.Header.Hash()) != 0 {
		return false
	}

	return true
}

func (f *facilitator) onProposalComplete(rs *ctypes.RoundState) {
	f.log.Info("proposal complete",
		"height", rs.Height,
		"proposer", hex.EncodeToString(rs.Validators.GetProposer().PubKey.Address()),
		"block-hash", rs.ProposalBlock.Header.Hash())
	f.mtx.Lock()
	defer f.mtx.Unlock()
	f.rs = rs
}

func (f *facilitator) onEvent(evt interface{}) {

	tmed, ok := evt.(tmtmtypes.TMEventData)
	if !ok {
		f.log.Error("bad event type", "type", fmt.Sprintf("%T", evt))
		return
	}

	edrs, ok := tmed.Unwrap().(tmtmtypes.EventDataRoundState)
	if !ok {
		f.log.Error("bad event data type", "type", fmt.Sprintf("%T", tmed))
		return
	}

	if edrs.RoundState == nil {
		f.log.Error("nil round state")
		return
	}

	rs, ok := edrs.RoundState.(*ctypes.RoundState)
	if !ok {
		f.log.Error("bad round state type", "type", fmt.Sprintf("%T", edrs.RoundState))
		return
	}

	f.onProposalComplete(rs)
}

func NewFacilitator(log log.Logger, validator tmtmtypes.PrivValidatorFS, bus *tmtmtypes.EventBus) (Facilitator, error) {

	f := &facilitator{
		log:       log,
		validator: validator,
		mtx:       new(sync.Mutex),
	}

	ch := make(chan interface{})

	if err := bus.Subscribe(context.Background(), "photon-market", tmtmtypes.EventQueryCompleteProposal, ch); err != nil {
		return nil, err
	}

	go func(ch chan interface{}) {
		for evt := range ch {
			f.onEvent(evt)
		}
	}(ch)

	return f, nil
}
