package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"loan/testutil/sample"
)

func TestMsgCreateDepositEarned_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDepositEarned
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDepositEarned{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDepositEarned{
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

func TestMsgUpdateDepositEarned_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDepositEarned
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDepositEarned{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDepositEarned{
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

func TestMsgDeleteDepositEarned_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteDepositEarned
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteDepositEarned{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteDepositEarned{
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
