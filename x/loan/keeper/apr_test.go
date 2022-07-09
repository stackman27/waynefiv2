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

func createNApr(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Apr {
	items := make([]types.Apr, n)
	for i := range items {
		items[i].Id = keeper.AppendApr(ctx, items[i])
	}
	return items
}

func TestAprGet(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNApr(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetApr(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAprRemove(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNApr(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApr(ctx, item.Id)
		_, found := keeper.GetApr(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAprGetAll(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNApr(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllApr(ctx)),
	)
}

func TestAprCount(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNApr(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAprCount(ctx))
}
