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

func createNDepositEarned(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DepositEarned {
	items := make([]types.DepositEarned, n)
	for i := range items {
		items[i].Id = keeper.AppendDepositEarned(ctx, items[i])
	}
	return items
}

func TestDepositEarnedGet(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNDepositEarned(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetDepositEarned(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestDepositEarnedRemove(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNDepositEarned(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDepositEarned(ctx, item.Id)
		_, found := keeper.GetDepositEarned(ctx, item.Id)
		require.False(t, found)
	}
}

func TestDepositEarnedGetAll(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNDepositEarned(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDepositEarned(ctx)),
	)
}

func TestDepositEarnedCount(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNDepositEarned(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetDepositEarnedCount(ctx))
}
