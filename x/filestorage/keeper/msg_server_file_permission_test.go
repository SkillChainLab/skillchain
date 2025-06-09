package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "skillchain/testutil/keeper"
	"skillchain/x/filestorage/keeper"
	"skillchain/x/filestorage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFilePermissionMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.FilestorageKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFilePermission{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateFilePermission(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetFilePermission(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestFilePermissionMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateFilePermission
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateFilePermission{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateFilePermission{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateFilePermission{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.FilestorageKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateFilePermission{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateFilePermission(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateFilePermission(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetFilePermission(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestFilePermissionMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteFilePermission
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteFilePermission{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteFilePermission{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteFilePermission{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.FilestorageKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateFilePermission(ctx, &types.MsgCreateFilePermission{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteFilePermission(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetFilePermission(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
