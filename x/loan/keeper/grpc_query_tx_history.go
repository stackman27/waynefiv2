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

func (k Keeper) TxHistoryAll(c context.Context, req *types.QueryAllTxHistoryRequest) (*types.QueryAllTxHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var txHistorys []types.TxHistory
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	txHistoryStore := prefix.NewStore(store, types.KeyPrefix(types.TxHistoryKey))

	pageRes, err := query.Paginate(txHistoryStore, req.Pagination, func(key []byte, value []byte) error {
		var txHistory types.TxHistory
		if err := k.cdc.Unmarshal(value, &txHistory); err != nil {
			return err
		}

		txHistorys = append(txHistorys, txHistory)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTxHistoryResponse{TxHistory: txHistorys, Pagination: pageRes}, nil
}

func (k Keeper) TxHistory(c context.Context, req *types.QueryGetTxHistoryRequest) (*types.QueryGetTxHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	txHistory, found := k.GetTxHistory(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTxHistoryResponse{TxHistory: txHistory}, nil
}
