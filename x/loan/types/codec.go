package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestLoan{}, "loan/RequestLoan", nil)
	cdc.RegisterConcrete(&MsgApproveLoan{}, "loan/ApproveLoan", nil)
	cdc.RegisterConcrete(&MsgRepayLoan{}, "loan/RepayLoan", nil)
	cdc.RegisterConcrete(&MsgLiquidateLoan{}, "loan/LiquidateLoan", nil)
	cdc.RegisterConcrete(&MsgCancelLoan{}, "loan/CancelLoan", nil)
	cdc.RegisterConcrete(&MsgCreateTxHistory{}, "loan/CreateTxHistory", nil)
	cdc.RegisterConcrete(&MsgUpdateTxHistory{}, "loan/UpdateTxHistory", nil)
	cdc.RegisterConcrete(&MsgDeleteTxHistory{}, "loan/DeleteTxHistory", nil)
	cdc.RegisterConcrete(&MsgCreateDepositEarned{}, "loan/CreateDepositEarned", nil)
	cdc.RegisterConcrete(&MsgUpdateDepositEarned{}, "loan/UpdateDepositEarned", nil)
	cdc.RegisterConcrete(&MsgDeleteDepositEarned{}, "loan/DeleteDepositEarned", nil)
	cdc.RegisterConcrete(&MsgCreateBorrowAccured{}, "loan/CreateBorrowAccured", nil)
	cdc.RegisterConcrete(&MsgUpdateBorrowAccured{}, "loan/UpdateBorrowAccured", nil)
	cdc.RegisterConcrete(&MsgDeleteBorrowAccured{}, "loan/DeleteBorrowAccured", nil)
	cdc.RegisterConcrete(&MsgCreateDeposit{}, "loan/CreateDeposit", nil)
	cdc.RegisterConcrete(&MsgUpdateDeposit{}, "loan/UpdateDeposit", nil)
	cdc.RegisterConcrete(&MsgDeleteDeposit{}, "loan/DeleteDeposit", nil)
	cdc.RegisterConcrete(&MsgCreateBorrow{}, "loan/CreateBorrow", nil)
	cdc.RegisterConcrete(&MsgUpdateBorrow{}, "loan/UpdateBorrow", nil)
	cdc.RegisterConcrete(&MsgDeleteBorrow{}, "loan/DeleteBorrow", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRepayLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLiquidateLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTxHistory{},
		&MsgUpdateTxHistory{},
		&MsgDeleteTxHistory{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDepositEarned{},
		&MsgUpdateDepositEarned{},
		&MsgDeleteDepositEarned{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBorrowAccured{},
		&MsgUpdateBorrowAccured{},
		&MsgDeleteBorrowAccured{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDeposit{},
		&MsgUpdateDeposit{},
		&MsgDeleteDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBorrow{},
		&MsgUpdateBorrow{},
		&MsgDeleteBorrow{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
