package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"loan/x/loan/types"
)

func CmdCreateBorrowAccured() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-borrow-accured [block-height] [asset] [amount] [denom]",
		Short: "Create a new borrowAccured",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBlockHeight, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argAsset := args[1]
			argAmount, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argDenom := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBorrowAccured(clientCtx.GetFromAddress().String(), argBlockHeight, argAsset, argAmount, argDenom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBorrowAccured() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-borrow-accured [id] [block-height] [asset] [amount] [denom]",
		Short: "Update a borrowAccured",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argBlockHeight, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			argAsset := args[2]

			argAmount, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			argDenom := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBorrowAccured(clientCtx.GetFromAddress().String(), id, argBlockHeight, argAsset, argAmount, argDenom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBorrowAccured() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-borrow-accured [id]",
		Short: "Delete a borrowAccured by id",
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

			msg := types.NewMsgDeleteBorrowAccured(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
