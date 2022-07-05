package keeper

import (
	"context"
	"errors"

	"loan/x/loan/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelLoan(goCtx context.Context, msg *types.MsgCancelLoan) (*types.MsgCancelLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get loan
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errors.New("loan not found")
	}

	if loan.Borrower != msg.Creator {
		return nil, errors.New("Only the borrower can cancel the loan")
	}

	if loan.State != "requested" {
		return nil, errors.New("Loan can only be cancelled during the requeted phase")
	}

	collateralAmt, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(loan.Borrower), collateralAmt)

	loan.State = "cancelled"

	k.SetLoan(ctx, loan)

	return &types.MsgCancelLoanResponse{}, nil
}
