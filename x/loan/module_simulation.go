package loan

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"loan/testutil/sample"
	loansimulation "loan/x/loan/simulation"
	"loan/x/loan/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = loansimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestLoan = "op_weight_msg_request_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestLoan int = 100

	opWeightMsgApproveLoan = "op_weight_msg_approve_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveLoan int = 100

	opWeightMsgRepayLoan = "op_weight_msg_repay_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayLoan int = 100

	opWeightMsgLiquidateLoan = "op_weight_msg_liquidate_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLiquidateLoan int = 100

	opWeightMsgCancelLoan = "op_weight_msg_cancel_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelLoan int = 100

	opWeightMsgCreateTxHistory = "op_weight_msg_tx_history"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTxHistory int = 100

	opWeightMsgUpdateTxHistory = "op_weight_msg_tx_history"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTxHistory int = 100

	opWeightMsgDeleteTxHistory = "op_weight_msg_tx_history"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTxHistory int = 100

	opWeightMsgCreateDepositEarned = "op_weight_msg_deposit_earned"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDepositEarned int = 100

	opWeightMsgUpdateDepositEarned = "op_weight_msg_deposit_earned"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDepositEarned int = 100

	opWeightMsgDeleteDepositEarned = "op_weight_msg_deposit_earned"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDepositEarned int = 100

	opWeightMsgCreateBorrowAccured = "op_weight_msg_borrow_accured"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBorrowAccured int = 100

	opWeightMsgUpdateBorrowAccured = "op_weight_msg_borrow_accured"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBorrowAccured int = 100

	opWeightMsgDeleteBorrowAccured = "op_weight_msg_borrow_accured"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteBorrowAccured int = 100

	opWeightMsgCreateDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDeposit int = 100

	opWeightMsgUpdateDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDeposit int = 100

	opWeightMsgDeleteDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDeposit int = 100

	opWeightMsgCreateBorrow = "op_weight_msg_borrow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBorrow int = 100

	opWeightMsgUpdateBorrow = "op_weight_msg_borrow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBorrow int = 100

	opWeightMsgDeleteBorrow = "op_weight_msg_borrow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteBorrow int = 100

	opWeightMsgCreateApr = "op_weight_msg_apr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateApr int = 100

	opWeightMsgUpdateApr = "op_weight_msg_apr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateApr int = 100

	opWeightMsgDeleteApr = "op_weight_msg_apr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteApr int = 100

	opWeightMsgCreateUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUser int = 100

	opWeightMsgUpdateUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUser int = 100

	opWeightMsgDeleteUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUser int = 100

	opWeightMsgCreatePool = "op_weight_msg_create_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePool int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	loanGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TxHistoryList: []types.TxHistory{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TxHistoryCount: 2,
		DepositEarnedList: []types.DepositEarned{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		DepositEarnedCount: 2,
		BorrowAccuredList: []types.BorrowAccured{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		BorrowAccuredCount: 2,
		DepositList: []types.Deposit{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		DepositCount: 2,
		BorrowList: []types.Borrow{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		BorrowCount: 2,
		AprList: []types.Apr{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AprCount: 2,
		UserList: []types.User{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		UserCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&loanGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRequestLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestLoan, &weightMsgRequestLoan, nil,
		func(_ *rand.Rand) {
			weightMsgRequestLoan = defaultWeightMsgRequestLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestLoan,
		loansimulation.SimulateMsgRequestLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveLoan, &weightMsgApproveLoan, nil,
		func(_ *rand.Rand) {
			weightMsgApproveLoan = defaultWeightMsgApproveLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveLoan,
		loansimulation.SimulateMsgApproveLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRepayLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRepayLoan, &weightMsgRepayLoan, nil,
		func(_ *rand.Rand) {
			weightMsgRepayLoan = defaultWeightMsgRepayLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayLoan,
		loansimulation.SimulateMsgRepayLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLiquidateLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLiquidateLoan, &weightMsgLiquidateLoan, nil,
		func(_ *rand.Rand) {
			weightMsgLiquidateLoan = defaultWeightMsgLiquidateLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiquidateLoan,
		loansimulation.SimulateMsgLiquidateLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelLoan, &weightMsgCancelLoan, nil,
		func(_ *rand.Rand) {
			weightMsgCancelLoan = defaultWeightMsgCancelLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelLoan,
		loansimulation.SimulateMsgCancelLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTxHistory int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTxHistory, &weightMsgCreateTxHistory, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTxHistory = defaultWeightMsgCreateTxHistory
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTxHistory,
		loansimulation.SimulateMsgCreateTxHistory(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTxHistory int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTxHistory, &weightMsgUpdateTxHistory, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTxHistory = defaultWeightMsgUpdateTxHistory
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTxHistory,
		loansimulation.SimulateMsgUpdateTxHistory(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTxHistory int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTxHistory, &weightMsgDeleteTxHistory, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTxHistory = defaultWeightMsgDeleteTxHistory
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTxHistory,
		loansimulation.SimulateMsgDeleteTxHistory(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateDepositEarned int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDepositEarned, &weightMsgCreateDepositEarned, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDepositEarned = defaultWeightMsgCreateDepositEarned
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDepositEarned,
		loansimulation.SimulateMsgCreateDepositEarned(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDepositEarned int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDepositEarned, &weightMsgUpdateDepositEarned, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDepositEarned = defaultWeightMsgUpdateDepositEarned
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDepositEarned,
		loansimulation.SimulateMsgUpdateDepositEarned(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDepositEarned int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDepositEarned, &weightMsgDeleteDepositEarned, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDepositEarned = defaultWeightMsgDeleteDepositEarned
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDepositEarned,
		loansimulation.SimulateMsgDeleteDepositEarned(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateBorrowAccured int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateBorrowAccured, &weightMsgCreateBorrowAccured, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBorrowAccured = defaultWeightMsgCreateBorrowAccured
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBorrowAccured,
		loansimulation.SimulateMsgCreateBorrowAccured(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBorrowAccured int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBorrowAccured, &weightMsgUpdateBorrowAccured, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBorrowAccured = defaultWeightMsgUpdateBorrowAccured
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBorrowAccured,
		loansimulation.SimulateMsgUpdateBorrowAccured(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteBorrowAccured int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteBorrowAccured, &weightMsgDeleteBorrowAccured, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteBorrowAccured = defaultWeightMsgDeleteBorrowAccured
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteBorrowAccured,
		loansimulation.SimulateMsgDeleteBorrowAccured(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDeposit, &weightMsgCreateDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDeposit = defaultWeightMsgCreateDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDeposit,
		loansimulation.SimulateMsgCreateDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDeposit, &weightMsgUpdateDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDeposit = defaultWeightMsgUpdateDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDeposit,
		loansimulation.SimulateMsgUpdateDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDeposit, &weightMsgDeleteDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDeposit = defaultWeightMsgDeleteDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDeposit,
		loansimulation.SimulateMsgDeleteDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateBorrow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateBorrow, &weightMsgCreateBorrow, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBorrow = defaultWeightMsgCreateBorrow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBorrow,
		loansimulation.SimulateMsgCreateBorrow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBorrow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBorrow, &weightMsgUpdateBorrow, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBorrow = defaultWeightMsgUpdateBorrow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBorrow,
		loansimulation.SimulateMsgUpdateBorrow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteBorrow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteBorrow, &weightMsgDeleteBorrow, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteBorrow = defaultWeightMsgDeleteBorrow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteBorrow,
		loansimulation.SimulateMsgDeleteBorrow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateApr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateApr, &weightMsgCreateApr, nil,
		func(_ *rand.Rand) {
			weightMsgCreateApr = defaultWeightMsgCreateApr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateApr,
		loansimulation.SimulateMsgCreateApr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateApr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateApr, &weightMsgUpdateApr, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateApr = defaultWeightMsgUpdateApr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateApr,
		loansimulation.SimulateMsgUpdateApr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteApr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteApr, &weightMsgDeleteApr, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteApr = defaultWeightMsgDeleteApr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteApr,
		loansimulation.SimulateMsgDeleteApr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateUser int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateUser, &weightMsgCreateUser, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUser = defaultWeightMsgCreateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUser,
		loansimulation.SimulateMsgCreateUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUser int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateUser, &weightMsgUpdateUser, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUser = defaultWeightMsgUpdateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUser,
		loansimulation.SimulateMsgUpdateUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUser int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteUser, &weightMsgDeleteUser, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUser = defaultWeightMsgDeleteUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUser,
		loansimulation.SimulateMsgDeleteUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePool, &weightMsgCreatePool, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePool = defaultWeightMsgCreatePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePool,
		loansimulation.SimulateMsgCreatePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
