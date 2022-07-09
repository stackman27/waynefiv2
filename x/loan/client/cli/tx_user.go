package cli

import (
	"strconv"

	"loan/x/loan/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user",
		Short: "Create a new user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			initialAtomDeposit := types.Deposit{Asset: "ATOM", Denom: "uatom"}
			initialCroDeposit := types.Deposit{Asset: "CRO", Denom: "ucro"}
			initialDvpnDeposit := types.Deposit{Asset: "DVPN", Denom: "udvpn"}
			initialIrisDeposit := types.Deposit{Asset: "IRIS", Denom: "uiris"}

			initialAtomBorrow := types.Borrow{Asset: "ATOM", Denom: "uatom"}
			initialCroBorrow := types.Borrow{Asset: "CRO", Denom: "ucro"}
			initialDvpnBorrow := types.Borrow{Asset: "DVPN", Denom: "udvpn"}
			initialIrisBorrow := types.Borrow{Asset: "IRIS", Denom: "uiris"}

			initialAtomDepositEarned := types.DepositEarned{Asset: "ATOM", Denom: "uatom"}
			initialCroDepositEarned := types.DepositEarned{Asset: "CRO", Denom: "ucro"}
			initialDvpnDepositEarned := types.DepositEarned{Asset: "IRIS", Denom: "uiris"}
			initialIrisDepositEarned := types.DepositEarned{Asset: "IRIS", Denom: "uiris"}

			initialAtomBorrowEarned := types.BorrowAccured{Asset: "ATOM", Denom: "uatom"}
			initialCroBorrowEarned := types.BorrowAccured{Asset: "CRO", Denom: "ucro"}
			initialDvpnBorrowEarned := types.BorrowAccured{Asset: "IRIS", Denom: "uiris"}
			initialIrisBorrowEarned := types.BorrowAccured{Asset: "IRIS", Denom: "uiris"}

			argsCollateral := []bool{false, false, false, false}
			argsDeposit := []*types.Deposit{&initialAtomDeposit, &initialCroDeposit, &initialDvpnDeposit, &initialIrisDeposit}
			argsBorrow := []*types.Borrow{&initialAtomBorrow, &initialCroBorrow, &initialDvpnBorrow, &initialIrisBorrow}
			argsAssetBalances := []int32{2000, 1000, 1000, 1000}
			argsDepositEarned := []*types.DepositEarned{&initialAtomDepositEarned, &initialCroDepositEarned, &initialDvpnDepositEarned, &initialIrisDepositEarned}
			argsBorrowAccured := []*types.BorrowAccured{&initialAtomBorrowEarned, &initialCroBorrowEarned, &initialDvpnBorrowEarned, &initialIrisBorrowEarned}

			argsTxHistories := []*types.TxHistory{}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUser(clientCtx.GetFromAddress().String(), argsCollateral, argsDeposit, argsBorrow, argsAssetBalances, argsDepositEarned, argsBorrowAccured, argsTxHistories)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-user [id]",
		Short: "Delete a user by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteUser(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
