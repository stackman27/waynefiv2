package keeper

import (
	"context"

	"loan/x/loan/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserAll(c context.Context, req *types.QueryAllUserRequest) (*types.QueryAllUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []types.User
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))

	pageRes, err := query.Paginate(userStore, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return err
		}

		users = append(users, user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserResponse{User: users, Pagination: pageRes}, nil
}

func (k Keeper) User(c context.Context, req *types.QueryGetUserRequest) (*types.QueryGetUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	user, found := k.GetUser(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetUserResponse{User: user}, nil
}

func (k Keeper) UserLoad(c context.Context, req *types.QueryLoadUserRequest) (*types.QueryLoadUserResponse, error) {

	return &types.QueryLoadUserResponse{}, nil
}

func (k Keeper) UserBalance(c context.Context, req *types.QueryBalanceUserRequest) (*types.QueryBalanceUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	//TODO: omit keynotfound error when user has not been found
	userAddress, err := sdk.AccAddressFromBech32(req.Id)
	if err != nil {
		return nil, err
	}

	assetBalances := k.bankKeeper.GetAllBalances(ctx, userAddress)[1:]

	var userBalance types.UserAssetBalances
	userBalance.Uatom = assetBalances[1].Amount.Int64()
	userBalance.Ucro = assetBalances[2].Amount.Int64()
	userBalance.Udvpn = assetBalances[3].Amount.Int64()
	userBalance.Uiris = assetBalances[4].Amount.Int64()

	return &types.QueryBalanceUserResponse{UserAssetBalances: &userBalance}, nil
}
