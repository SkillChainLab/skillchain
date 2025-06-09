package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "skillchain/testutil/keeper"
	"skillchain/x/marketplace/keeper"
	"skillchain/x/marketplace/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMilestoneMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.MarketplaceKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateMilestone{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateMilestone(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetMilestone(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestMilestoneMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateMilestone
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateMilestone{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateMilestone{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateMilestone{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MarketplaceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateMilestone{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateMilestone(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateMilestone(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetMilestone(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestMilestoneMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteMilestone
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteMilestone{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteMilestone{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteMilestone{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MarketplaceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateMilestone(ctx, &types.MsgCreateMilestone{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteMilestone(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetMilestone(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
