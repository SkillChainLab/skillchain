package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "skillchain/testutil/keeper"
	"skillchain/x/profile/keeper"
	"skillchain/x/profile/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestSkillEndorsementMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ProfileKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateSkillEndorsement{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateSkillEndorsement(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetSkillEndorsement(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestSkillEndorsementMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateSkillEndorsement
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateSkillEndorsement{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateSkillEndorsement{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateSkillEndorsement{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProfileKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateSkillEndorsement{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateSkillEndorsement(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateSkillEndorsement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetSkillEndorsement(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestSkillEndorsementMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteSkillEndorsement
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteSkillEndorsement{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteSkillEndorsement{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteSkillEndorsement{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProfileKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateSkillEndorsement(ctx, &types.MsgCreateSkillEndorsement{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteSkillEndorsement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetSkillEndorsement(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
