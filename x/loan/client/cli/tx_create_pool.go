package cli

import (
	"strconv"

	"loan/x/loan/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [asset] [denom] [collatoralFactor] [depositBalance] [borrowBalance]",
		Short: "Broadcast message create-pool",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argsAsset, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsDenom, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsCollatoralFactor, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argsDepositBalance, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argsBorrowBalance, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			argsUsers := []*types.User{}

			argsAprs := []*types.Apr{}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), argsAsset, argsDenom, argsCollatoralFactor, argsDepositBalance, argsBorrowBalance, argsUsers, argsAprs)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
