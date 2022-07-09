package loan_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "loan/testutil/keeper"
	"loan/testutil/nullify"
	"loan/x/loan"
	"loan/x/loan/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LoanList: []types.Loan{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LoanCount: 2,
		TxHistoryList: []types.TxHistory{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TxHistoryCount: 2,
		DepositEarnedList: []types.DepositEarned{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DepositEarnedCount: 2,
		BorrowAccuredList: []types.BorrowAccured{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		BorrowAccuredCount: 2,
		DepositList: []types.Deposit{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DepositCount: 2,
		BorrowList: []types.Borrow{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		BorrowCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoanKeeper(t)
	loan.InitGenesis(ctx, *k, genesisState)
	got := loan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LoanList, got.LoanList)
	require.Equal(t, genesisState.LoanCount, got.LoanCount)
	require.ElementsMatch(t, genesisState.TxHistoryList, got.TxHistoryList)
	require.Equal(t, genesisState.TxHistoryCount, got.TxHistoryCount)
	require.ElementsMatch(t, genesisState.DepositEarnedList, got.DepositEarnedList)
	require.Equal(t, genesisState.DepositEarnedCount, got.DepositEarnedCount)
	require.ElementsMatch(t, genesisState.BorrowAccuredList, got.BorrowAccuredList)
	require.Equal(t, genesisState.BorrowAccuredCount, got.BorrowAccuredCount)
	require.ElementsMatch(t, genesisState.DepositList, got.DepositList)
	require.Equal(t, genesisState.DepositCount, got.DepositCount)
	require.ElementsMatch(t, genesisState.BorrowList, got.BorrowList)
	require.Equal(t, genesisState.BorrowCount, got.BorrowCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
