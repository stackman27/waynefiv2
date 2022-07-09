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

func createNBorrowAccured(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BorrowAccured {
	items := make([]types.BorrowAccured, n)
	for i := range items {
		items[i].Id = keeper.AppendBorrowAccured(ctx, items[i])
	}
	return items
}

func TestBorrowAccuredGet(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNBorrowAccured(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetBorrowAccured(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBorrowAccuredRemove(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNBorrowAccured(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBorrowAccured(ctx, item.Id)
		_, found := keeper.GetBorrowAccured(ctx, item.Id)
		require.False(t, found)
	}
}

func TestBorrowAccuredGetAll(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNBorrowAccured(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBorrowAccured(ctx)),
	)
}

func TestBorrowAccuredCount(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNBorrowAccured(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetBorrowAccuredCount(ctx))
}
