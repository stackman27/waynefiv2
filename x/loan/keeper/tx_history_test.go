package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "loan/testutil/keeper"
	"loan/testutil/nullify"
	"loan/x/loan/keeper"
	"loan/x/loan/types"
)

func createNTxHistory(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TxHistory {
	items := make([]types.TxHistory, n)
	for i := range items {
		items[i].Id = keeper.AppendTxHistory(ctx, items[i])
	}
	return items
}

func TestTxHistoryGet(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNTxHistory(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTxHistory(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTxHistoryRemove(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNTxHistory(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTxHistory(ctx, item.Id)
		_, found := keeper.GetTxHistory(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTxHistoryGetAll(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNTxHistory(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTxHistory(ctx)),
	)
}

func TestTxHistoryCount(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNTxHistory(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTxHistoryCount(ctx))
}
