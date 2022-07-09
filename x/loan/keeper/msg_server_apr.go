package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"loan/x/loan/types"
)

func (k msgServer) CreateApr(goCtx context.Context, msg *types.MsgCreateApr) (*types.MsgCreateAprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var apr = types.Apr{
		Creator:     msg.Creator,
		BlockHeight: msg.BlockHeight,
		DepositApy:  msg.DepositApy,
		BorrowApy:   msg.BorrowApy,
	}

	id := k.AppendApr(
		ctx,
		apr,
	)

	return &types.MsgCreateAprResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateApr(goCtx context.Context, msg *types.MsgUpdateApr) (*types.MsgUpdateAprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var apr = types.Apr{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		DepositApy:  msg.DepositApy,
		BorrowApy:   msg.BorrowApy,
	}

	// Checks that the element exists
	val, found := k.GetApr(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetApr(ctx, apr)

	return &types.MsgUpdateAprResponse{}, nil
}

func (k msgServer) DeleteApr(goCtx context.Context, msg *types.MsgDeleteApr) (*types.MsgDeleteAprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetApr(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveApr(ctx, msg.Id)

	return &types.MsgDeleteAprResponse{}, nil
}
