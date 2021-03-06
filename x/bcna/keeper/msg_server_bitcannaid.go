package keeper

import (
	"context"
	"fmt"

	"github.com/BitCannaGlobal/testnet-bcna-cosmos/x/bcna/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateBitcannaid(goCtx context.Context, msg *types.MsgCreateBitcannaid) (*types.MsgCreateBitcannaidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendBitcannaid(
		ctx,
		msg.Creator,
		msg.Bcnaid,
		msg.Address,
	)

	return &types.MsgCreateBitcannaidResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBitcannaid(goCtx context.Context, msg *types.MsgUpdateBitcannaid) (*types.MsgUpdateBitcannaidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var bitcannaid = types.Bitcannaid{
		Creator: msg.Creator,
		Id:      msg.Id,
		Bcnaid:  msg.Bcnaid,
		Address: msg.Address,
	}

	// Checks that the element exists
	if !k.HasBitcannaid(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetBitcannaidOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBitcannaid(ctx, bitcannaid)

	return &types.MsgUpdateBitcannaidResponse{}, nil
}

func (k msgServer) DeleteBitcannaid(goCtx context.Context, msg *types.MsgDeleteBitcannaid) (*types.MsgDeleteBitcannaidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasBitcannaid(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetBitcannaidOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBitcannaid(ctx, msg.Id)

	return &types.MsgDeleteBitcannaidResponse{}, nil
}
