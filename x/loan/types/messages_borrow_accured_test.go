package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"loan/testutil/sample"
)

func TestMsgCreateBorrowAccured_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateBorrowAccured
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateBorrowAccured{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateBorrowAccured{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateBorrowAccured_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateBorrowAccured
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateBorrowAccured{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateBorrowAccured{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteBorrowAccured_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteBorrowAccured
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteBorrowAccured{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteBorrowAccured{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
