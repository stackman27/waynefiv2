package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"loan/x/loan/types"
)

func (k msgServer) CreateBorrowAccured(goCtx context.Context, msg *types.MsgCreateBorrowAccured) (*types.MsgCreateBorrowAccuredResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var borrowAccured = types.BorrowAccured{
		Creator:     msg.Creator,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	id := k.AppendBorrowAccured(
		ctx,
		borrowAccured,
	)

	return &types.MsgCreateBorrowAccuredResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBorrowAccured(goCtx context.Context, msg *types.MsgUpdateBorrowAccured) (*types.MsgUpdateBorrowAccuredResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var borrowAccured = types.BorrowAccured{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	val, found := k.GetBorrowAccured(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBorrowAccured(ctx, borrowAccured)

	return &types.MsgUpdateBorrowAccuredResponse{}, nil
}

func (k msgServer) DeleteBorrowAccured(goCtx context.Context, msg *types.MsgDeleteBorrowAccured) (*types.MsgDeleteBorrowAccuredResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetBorrowAccured(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBorrowAccured(ctx, msg.Id)

	return &types.MsgDeleteBorrowAccuredResponse{}, nil
}
