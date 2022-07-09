package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"loan/testutil/sample"
)

func TestMsgCreateTxHistory_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTxHistory
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTxHistory{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTxHistory{
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

func TestMsgUpdateTxHistory_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTxHistory
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTxHistory{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTxHistory{
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

func TestMsgDeleteTxHistory_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTxHistory
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTxHistory{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTxHistory{
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
