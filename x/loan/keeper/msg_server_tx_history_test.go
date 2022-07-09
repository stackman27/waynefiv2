package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"loan/x/loan/types"
)

func TestTxHistoryMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateTxHistory(ctx, &types.MsgCreateTxHistory{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestTxHistoryMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateTxHistory
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTxHistory{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTxHistory{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTxHistory{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateTxHistory(ctx, &types.MsgCreateTxHistory{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateTxHistory(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTxHistoryMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTxHistory
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTxHistory{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTxHistory{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteTxHistory{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateTxHistory(ctx, &types.MsgCreateTxHistory{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteTxHistory(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
