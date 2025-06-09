package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "skillchain/testutil/keeper"
	"skillchain/x/analytics/keeper"
	"skillchain/x/analytics/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUserActivityMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AnalyticsKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateUserActivity{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateUserActivity(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetUserActivity(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestUserActivityMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateUserActivity
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateUserActivity{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateUserActivity{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateUserActivity{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AnalyticsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateUserActivity{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateUserActivity(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateUserActivity(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetUserActivity(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestUserActivityMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteUserActivity
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteUserActivity{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteUserActivity{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteUserActivity{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AnalyticsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateUserActivity(ctx, &types.MsgCreateUserActivity{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteUserActivity(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetUserActivity(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
