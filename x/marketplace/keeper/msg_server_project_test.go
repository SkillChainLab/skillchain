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

func TestProjectMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.MarketplaceKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateProject{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateProject(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetProject(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestProjectMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateProject
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateProject{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateProject{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateProject{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MarketplaceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateProject{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateProject(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateProject(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetProject(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestProjectMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteProject
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteProject{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteProject{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteProject{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MarketplaceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateProject(ctx, &types.MsgCreateProject{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteProject(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetProject(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
