package keeper

import (
	"context"
	"errors"

	"loan/x/loan/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the loan actually exist
	loan, ok := k.GetLoan(ctx, msg.Id)
	if !ok {
		return nil, errors.New("Could not find the loan requested")
	}

	if loan.State != "requested" {
		return nil, errors.New("Loan couldn't be approved. Make sure it's in requested state")
	}

	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
		return nil, err
	}

	k.bankKeeper.SendCoins(ctx, lender, borrower, amount)

	loan.Lender = msg.Creator
	loan.State = "approved"

	k.SetLoan(ctx, loan)

	return &types.MsgApproveLoanResponse{}, nil
}
