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

func TestNotificationSettingsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NotificationsKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateNotificationSettings{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateNotificationSettings(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetNotificationSettings(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestNotificationSettingsMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateNotificationSettings
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateNotificationSettings{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateNotificationSettings{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateNotificationSettings{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NotificationsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateNotificationSettings{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateNotificationSettings(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateNotificationSettings(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetNotificationSettings(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestNotificationSettingsMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteNotificationSettings
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteNotificationSettings{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteNotificationSettings{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteNotificationSettings{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NotificationsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateNotificationSettings(ctx, &types.MsgCreateNotificationSettings{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteNotificationSettings(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetNotificationSettings(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
