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

func TestJobPostingMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.MarketplaceKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateJobPosting{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateJobPosting(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetJobPosting(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestJobPostingMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateJobPosting
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateJobPosting{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateJobPosting{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateJobPosting{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MarketplaceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateJobPosting{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateJobPosting(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateJobPosting(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetJobPosting(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestJobPostingMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteJobPosting
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteJobPosting{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteJobPosting{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteJobPosting{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MarketplaceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateJobPosting(ctx, &types.MsgCreateJobPosting{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteJobPosting(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetJobPosting(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
