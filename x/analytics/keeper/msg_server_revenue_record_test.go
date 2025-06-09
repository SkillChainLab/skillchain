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

func TestRevenueRecordMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AnalyticsKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateRevenueRecord{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateRevenueRecord(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetRevenueRecord(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestRevenueRecordMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateRevenueRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateRevenueRecord{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateRevenueRecord{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateRevenueRecord{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AnalyticsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateRevenueRecord{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateRevenueRecord(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateRevenueRecord(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetRevenueRecord(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestRevenueRecordMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteRevenueRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteRevenueRecord{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteRevenueRecord{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteRevenueRecord{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AnalyticsKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateRevenueRecord(ctx, &types.MsgCreateRevenueRecord{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteRevenueRecord(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetRevenueRecord(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
