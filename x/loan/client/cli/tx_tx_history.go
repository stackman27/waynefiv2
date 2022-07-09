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

func CmdCreateTxHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-tx-history [block-height] [tx] [asset] [amount] [denom]",
		Short: "Create a new txHistory",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBlockHeight, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argTx := args[1]
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

			msg := types.NewMsgCreateTxHistory(clientCtx.GetFromAddress().String(), argBlockHeight, argTx, argAsset, argAmount, argDenom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateTxHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-tx-history [id] [block-height] [tx] [asset] [amount] [denom]",
		Short: "Update a txHistory",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argBlockHeight, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			argTx := args[2]

			argAsset := args[3]

			argAmount, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			argDenom := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTxHistory(clientCtx.GetFromAddress().String(), id, argBlockHeight, argTx, argAsset, argAmount, argDenom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteTxHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-tx-history [id]",
		Short: "Delete a txHistory by id",
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

			msg := types.NewMsgDeleteTxHistory(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
