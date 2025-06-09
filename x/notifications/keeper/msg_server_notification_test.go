package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "skillchain/testutil/keeper"
	"skillchain/x/notifications/keeper"
	"skillchain/x/notifications/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestNotificationMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NotificationsKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateNotification{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateNotification(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetNotification(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestNotificationMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateNotification
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateNotification{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateNotification{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateNotification{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NotificationsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateNotification{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateNotification(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateNotification(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetNotification(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestNotificationMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteNotification
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteNotification{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteNotification{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteNotification{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NotificationsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateNotification(ctx, &types.MsgCreateNotification{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteNotification(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetNotification(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
