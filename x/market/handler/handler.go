package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ovrclk/akash/x/market/types"
)

// NewHandler returns a handler for "market" type messages
func NewHandler(keepers Keepers) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgCreateBid:
			return handleMsgCreateBid(ctx, keepers, msg)
		case types.MsgCloseBid:
			return handleMsgCloseBid(ctx, keepers, msg)
		case types.MsgCloseOrder:
			return handleMsgCloseOrder(ctx, keepers, msg)
		default:
			return nil, sdkerrors.ErrUnknownRequest
		}
	}
}

func handleMsgCreateBid(ctx sdk.Context, keepers Keepers, msg types.MsgCreateBid) (*sdk.Result, error) {
	if _, err := keepers.Market.CreateBid(ctx, msg.Order, msg.Provider, msg.Price); err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgCloseBid(ctx sdk.Context, keepers Keepers, msg types.MsgCloseBid) (*sdk.Result, error) {
	bid, ok := keepers.Market.GetBid(ctx, msg.BidID)
	if !ok {
		panic(types.ErrUnknownBid)
	}

	lease, ok := keepers.Market.GetLease(ctx, types.LeaseID(msg.BidID))
	if !ok {
		panic(types.ErrUnknownLeaseForBid)
	}

	order, ok := keepers.Market.GetOrder(ctx, msg.OrderID())
	if !ok {
		panic(types.ErrUnknownOrderForBid)
	}

	if lease.State != types.LeaseActive {
		panic(types.ErrLeaseNotActive)
	}

	if bid.State != types.BidMatched {
		panic(types.ErrBidNotMatched)
	}

	keepers.Market.OnBidClosed(ctx, bid)
	keepers.Market.OnLeaseClosed(ctx, lease)
	keepers.Market.OnOrderClosed(ctx, order)
	keepers.Deployment.OnLeaseClosed(ctx, order.GroupID())

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgCloseOrder(ctx sdk.Context, keepers Keepers, msg types.MsgCloseOrder) (*sdk.Result, error) {
	order, ok := keepers.Market.GetOrder(ctx, msg.OrderID)
	if !ok {
		panic(types.ErrUnknownOrder)
	}

	lease, ok := keepers.Market.LeaseForOrder(ctx, order.ID())
	if !ok {
		panic(types.ErrNoLeaseForOrder)
	}
	keepers.Market.OnOrderClosed(ctx, order)
	keepers.Market.OnLeaseClosed(ctx, lease)
	keepers.Deployment.OnLeaseClosed(ctx, order.GroupID())
	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}
