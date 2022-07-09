package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"loan/x/loan/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRequestLoan())
	cmd.AddCommand(CmdApproveLoan())
	cmd.AddCommand(CmdRepayLoan())
	cmd.AddCommand(CmdLiquidateLoan())
	cmd.AddCommand(CmdCancelLoan())
	cmd.AddCommand(CmdCreateTxHistory())
	cmd.AddCommand(CmdUpdateTxHistory())
	cmd.AddCommand(CmdDeleteTxHistory())
	cmd.AddCommand(CmdCreateDepositEarned())
	cmd.AddCommand(CmdUpdateDepositEarned())
	cmd.AddCommand(CmdDeleteDepositEarned())
	cmd.AddCommand(CmdCreateBorrowAccured())
	cmd.AddCommand(CmdUpdateBorrowAccured())
	cmd.AddCommand(CmdDeleteBorrowAccured())
	cmd.AddCommand(CmdCreateDeposit())
	cmd.AddCommand(CmdUpdateDeposit())
	cmd.AddCommand(CmdDeleteDeposit())
	cmd.AddCommand(CmdCreateBorrow())
	cmd.AddCommand(CmdUpdateBorrow())
	cmd.AddCommand(CmdDeleteBorrow())
	cmd.AddCommand(CmdCreateApr())
	cmd.AddCommand(CmdUpdateApr())
	cmd.AddCommand(CmdDeleteApr())
	cmd.AddCommand(CmdCreateUser())
	cmd.AddCommand(CmdDeleteUser())
	// this line is used by starport scaffolding # 1

	return cmd
}
