package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"loan/x/loan/types"
)

func TestBorrowAccuredMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateBorrowAccured(ctx, &types.MsgCreateBorrowAccured{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestBorrowAccuredMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateBorrowAccured
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateBorrowAccured{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBorrowAccured{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBorrowAccured{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateBorrowAccured(ctx, &types.MsgCreateBorrowAccured{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateBorrowAccured(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBorrowAccuredMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteBorrowAccured
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteBorrowAccured{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteBorrowAccured{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteBorrowAccured{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateBorrowAccured(ctx, &types.MsgCreateBorrowAccured{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteBorrowAccured(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
