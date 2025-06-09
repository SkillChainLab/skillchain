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

func TestFileRecordMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.FilestorageKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFileRecord{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateFileRecord(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetFileRecord(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestFileRecordMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateFileRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateFileRecord{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateFileRecord{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateFileRecord{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.FilestorageKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateFileRecord{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateFileRecord(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateFileRecord(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetFileRecord(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestFileRecordMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteFileRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteFileRecord{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteFileRecord{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteFileRecord{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.FilestorageKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateFileRecord(ctx, &types.MsgCreateFileRecord{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteFileRecord(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetFileRecord(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
