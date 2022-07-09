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

func (k Keeper) AprAll(c context.Context, req *types.QueryAllAprRequest) (*types.QueryAllAprResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var aprs []types.Apr
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	aprStore := prefix.NewStore(store, types.KeyPrefix(types.AprKey))

	pageRes, err := query.Paginate(aprStore, req.Pagination, func(key []byte, value []byte) error {
		var apr types.Apr
		if err := k.cdc.Unmarshal(value, &apr); err != nil {
			return err
		}

		aprs = append(aprs, apr)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAprResponse{Apr: aprs, Pagination: pageRes}, nil
}

func (k Keeper) Apr(c context.Context, req *types.QueryGetAprRequest) (*types.QueryGetAprResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	apr, found := k.GetApr(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAprResponse{Apr: apr}, nil
}
