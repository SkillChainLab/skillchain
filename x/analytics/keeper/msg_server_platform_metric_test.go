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

func TestPlatformMetricMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AnalyticsKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreatePlatformMetric{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreatePlatformMetric(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetPlatformMetric(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestPlatformMetricMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdatePlatformMetric
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdatePlatformMetric{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdatePlatformMetric{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdatePlatformMetric{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AnalyticsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreatePlatformMetric{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreatePlatformMetric(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePlatformMetric(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPlatformMetric(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPlatformMetricMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePlatformMetric
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeletePlatformMetric{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeletePlatformMetric{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeletePlatformMetric{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AnalyticsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreatePlatformMetric(ctx, &types.MsgCreatePlatformMetric{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeletePlatformMetric(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPlatformMetric(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
