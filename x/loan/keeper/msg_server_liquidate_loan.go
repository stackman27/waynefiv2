package keeper

import (
	"context"
	"errors"
	"strconv"

	"loan/x/loan/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LiquidateLoan(goCtx context.Context, msg *types.MsgLiquidateLoan) (*types.MsgLiquidateLoanResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// liquidate loan
	loan, ok := k.GetLoan(ctx, msg.Id)
	if !ok {
		return nil, errors.New("Error getting loan")
	}

	// liquidate-loan must be executed by the lender
	if msg.Creator != loan.Lender {
		return nil, errors.New("Not the lender of this loan")
	}

	// the status of the loan must be approved
	if loan.State != "approved" {
		return nil, errors.New("The loan myst be approved to be liquidated")
	}

	deadline, err := strconv.ParseInt(loan.Deadline, 10, 64)
	if err != nil {
		panic(err)
	}

	// the deadline block height must have passed
	if deadline > ctx.BlockHeight() {
		return nil, errors.New("The deadline hasn't passed ")
	}

	collateralAmt, _ := sdk.ParseCoinsNormalized(loan.Collateral)

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(loan.Lender), collateralAmt)

	loan.State = "liquidated"

	k.SetLoan(ctx, loan)

	return &types.MsgLiquidateLoanResponse{}, nil
}
