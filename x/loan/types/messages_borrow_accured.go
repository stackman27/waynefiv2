package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBorrowAccured = "create_borrow_accured"
	TypeMsgUpdateBorrowAccured = "update_borrow_accured"
	TypeMsgDeleteBorrowAccured = "delete_borrow_accured"
)

var _ sdk.Msg = &MsgCreateBorrowAccured{}

func NewMsgCreateBorrowAccured(creator string, blockHeight int32, asset string, amount int32, denom string) *MsgCreateBorrowAccured {
	return &MsgCreateBorrowAccured{
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgCreateBorrowAccured) Route() string {
	return RouterKey
}

func (msg *MsgCreateBorrowAccured) Type() string {
	return TypeMsgCreateBorrowAccured
}

func (msg *MsgCreateBorrowAccured) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBorrowAccured) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBorrowAccured) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBorrowAccured{}

func NewMsgUpdateBorrowAccured(creator string, id uint64, blockHeight int32, asset string, amount int32, denom string) *MsgUpdateBorrowAccured {
	return &MsgUpdateBorrowAccured{
		Id:          id,
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgUpdateBorrowAccured) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBorrowAccured) Type() string {
	return TypeMsgUpdateBorrowAccured
}

func (msg *MsgUpdateBorrowAccured) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBorrowAccured) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBorrowAccured) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBorrowAccured{}

func NewMsgDeleteBorrowAccured(creator string, id uint64) *MsgDeleteBorrowAccured {
	return &MsgDeleteBorrowAccured{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteBorrowAccured) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBorrowAccured) Type() string {
	return TypeMsgDeleteBorrowAccured
}

func (msg *MsgDeleteBorrowAccured) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBorrowAccured) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBorrowAccured) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
