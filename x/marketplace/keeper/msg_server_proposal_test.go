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

func TestProposalMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.MarketplaceKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateProposal{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateProposal(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetProposal(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestProposalMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateProposal
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateProposal{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateProposal{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateProposal{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MarketplaceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateProposal{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateProposal(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateProposal(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetProposal(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestProposalMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteProposal
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteProposal{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteProposal{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteProposal{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MarketplaceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateProposal(ctx, &types.MsgCreateProposal{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteProposal(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetProposal(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
