package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"loan/x/loan/types"
)

func TestDepositEarnedMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateDepositEarned(ctx, &types.MsgCreateDepositEarned{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestDepositEarnedMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDepositEarned
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateDepositEarned{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateDepositEarned{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateDepositEarned{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateDepositEarned(ctx, &types.MsgCreateDepositEarned{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateDepositEarned(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDepositEarnedMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDepositEarned
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteDepositEarned{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteDepositEarned{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteDepositEarned{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateDepositEarned(ctx, &types.MsgCreateDepositEarned{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteDepositEarned(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
