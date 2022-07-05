package keeper

import (
	"context"
	"errors"

	"loan/x/loan/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RepayLoan(goCtx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, ok := k.GetLoan(ctx, msg.Id)
	if !ok {
		return nil, errors.New("Couldn't get the requested loan")
	}

	if loan.State != "approved" {
		return nil, errors.New("Loan not approved cannot repay")
	}

	// Check if the creator of this message is actuallly the borrower
	if msg.Creator != loan.Borrower {
		return nil, errors.New("The creator of the message is not the borrower")
	}

	// return the collateral back to the borrower
	collateralAmt, _ := sdk.ParseCoinsNormalized(loan.Collateral)

	// lender recieves the fees
	fees, _ := sdk.ParseCoinsNormalized(loan.Fee)

	// the amount borrowed
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)

	// send the collateral and fees back
	k.bankKeeper.SendCoins(ctx, sdk.AccAddress(loan.Borrower), sdk.AccAddress(loan.Lender), amount)

	k.bankKeeper.SendCoins(ctx, sdk.AccAddress(loan.Borrower), sdk.AccAddress(loan.Lender), fees)
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(loan.Borrower), collateralAmt)

	loan.State = "repayed"
	k.SetLoan(ctx, loan)

	return &types.MsgRepayLoanResponse{}, nil
}
