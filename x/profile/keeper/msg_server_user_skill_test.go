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

func TestUserSkillMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ProfileKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateUserSkill{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateUserSkill(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetUserSkill(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestUserSkillMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateUserSkill
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateUserSkill{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateUserSkill{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateUserSkill{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProfileKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateUserSkill{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateUserSkill(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateUserSkill(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetUserSkill(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestUserSkillMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteUserSkill
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteUserSkill{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteUserSkill{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteUserSkill{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProfileKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateUserSkill(ctx, &types.MsgCreateUserSkill{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteUserSkill(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetUserSkill(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
