package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"loan/x/loan/types"
)

func (k Keeper) BorrowAccuredAll(c context.Context, req *types.QueryAllBorrowAccuredRequest) (*types.QueryAllBorrowAccuredResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var borrowAccureds []types.BorrowAccured
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	borrowAccuredStore := prefix.NewStore(store, types.KeyPrefix(types.BorrowAccuredKey))

	pageRes, err := query.Paginate(borrowAccuredStore, req.Pagination, func(key []byte, value []byte) error {
		var borrowAccured types.BorrowAccured
		if err := k.cdc.Unmarshal(value, &borrowAccured); err != nil {
			return err
		}

		borrowAccureds = append(borrowAccureds, borrowAccured)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBorrowAccuredResponse{BorrowAccured: borrowAccureds, Pagination: pageRes}, nil
}

func (k Keeper) BorrowAccured(c context.Context, req *types.QueryGetBorrowAccuredRequest) (*types.QueryGetBorrowAccuredResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	borrowAccured, found := k.GetBorrowAccured(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetBorrowAccuredResponse{BorrowAccured: borrowAccured}, nil
}
