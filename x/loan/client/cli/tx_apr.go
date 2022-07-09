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

func CmdCreateApr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-apr [block-height] [deposit-apy] [borrow-apy]",
		Short: "Create a new apr",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBlockHeight, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argDepositApy, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argBorrowApy, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateApr(clientCtx.GetFromAddress().String(), argBlockHeight, argDepositApy, argBorrowApy)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateApr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-apr [id] [block-height] [deposit-apy] [borrow-apy]",
		Short: "Update a apr",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argBlockHeight, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			argDepositApy, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			argBorrowApy, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateApr(clientCtx.GetFromAddress().String(), id, argBlockHeight, argDepositApy, argBorrowApy)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteApr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-apr [id]",
		Short: "Delete a apr by id",
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

			msg := types.NewMsgDeleteApr(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
