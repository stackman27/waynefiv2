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

func (k Keeper) DepositEarnedAll(c context.Context, req *types.QueryAllDepositEarnedRequest) (*types.QueryAllDepositEarnedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var depositEarneds []types.DepositEarned
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	depositEarnedStore := prefix.NewStore(store, types.KeyPrefix(types.DepositEarnedKey))

	pageRes, err := query.Paginate(depositEarnedStore, req.Pagination, func(key []byte, value []byte) error {
		var depositEarned types.DepositEarned
		if err := k.cdc.Unmarshal(value, &depositEarned); err != nil {
			return err
		}

		depositEarneds = append(depositEarneds, depositEarned)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDepositEarnedResponse{DepositEarned: depositEarneds, Pagination: pageRes}, nil
}

func (k Keeper) DepositEarned(c context.Context, req *types.QueryGetDepositEarnedRequest) (*types.QueryGetDepositEarnedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	depositEarned, found := k.GetDepositEarned(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetDepositEarnedResponse{DepositEarned: depositEarned}, nil
}
