package loan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/keeper"
	"loan/x/loan/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the loan
	for _, elem := range genState.LoanList {
		k.SetLoan(ctx, elem)
	}

	// Set loan count
	k.SetLoanCount(ctx, genState.LoanCount)
	// Set all the txHistory
	for _, elem := range genState.TxHistoryList {
		k.SetTxHistory(ctx, elem)
	}

	// Set txHistory count
	k.SetTxHistoryCount(ctx, genState.TxHistoryCount)
	// Set all the depositEarned
	for _, elem := range genState.DepositEarnedList {
		k.SetDepositEarned(ctx, elem)
	}

	// Set depositEarned count
	k.SetDepositEarnedCount(ctx, genState.DepositEarnedCount)
	// Set all the borrowAccured
	for _, elem := range genState.BorrowAccuredList {
		k.SetBorrowAccured(ctx, elem)
	}

	// Set borrowAccured count
	k.SetBorrowAccuredCount(ctx, genState.BorrowAccuredCount)
	// Set all the deposit
	for _, elem := range genState.DepositList {
		k.SetDeposit(ctx, elem)
	}

	// Set deposit count
	k.SetDepositCount(ctx, genState.DepositCount)
	// Set all the borrow
	for _, elem := range genState.BorrowList {
		k.SetBorrow(ctx, elem)
	}

	// Set borrow count
	k.SetBorrowCount(ctx, genState.BorrowCount)
	// Set all the apr
	for _, elem := range genState.AprList {
		k.SetApr(ctx, elem)
	}

	// Set apr count
	k.SetAprCount(ctx, genState.AprCount)
	// Set all the user
	for _, elem := range genState.UserList {
		k.SetUser(ctx, elem)
	}

	// Set user count
	k.SetUserCount(ctx, genState.UserCount)
	// Set all the pool
	for _, elem := range genState.PoolList {
		k.SetPool(ctx, elem)
	}

	// Set pool count
	k.SetPoolCount(ctx, genState.PoolCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LoanList = k.GetAllLoan(ctx)
	genesis.LoanCount = k.GetLoanCount(ctx)
	genesis.TxHistoryList = k.GetAllTxHistory(ctx)
	genesis.TxHistoryCount = k.GetTxHistoryCount(ctx)
	genesis.DepositEarnedList = k.GetAllDepositEarned(ctx)
	genesis.DepositEarnedCount = k.GetDepositEarnedCount(ctx)
	genesis.BorrowAccuredList = k.GetAllBorrowAccured(ctx)
	genesis.BorrowAccuredCount = k.GetBorrowAccuredCount(ctx)
	genesis.DepositList = k.GetAllDeposit(ctx)
	genesis.DepositCount = k.GetDepositCount(ctx)
	genesis.BorrowList = k.GetAllBorrow(ctx)
	genesis.BorrowCount = k.GetBorrowCount(ctx)
	genesis.AprList = k.GetAllApr(ctx)
	genesis.AprCount = k.GetAprCount(ctx)
	genesis.UserList = k.GetAllUser(ctx)
	genesis.UserCount = k.GetUserCount(ctx)
	genesis.PoolList = k.GetAllPool(ctx)
	genesis.PoolCount = k.GetPoolCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
