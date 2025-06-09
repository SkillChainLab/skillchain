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

func TestUserProfileMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ProfileKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateUserProfile{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateUserProfile(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetUserProfile(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestUserProfileMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateUserProfile
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateUserProfile{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateUserProfile{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateUserProfile{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProfileKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateUserProfile{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateUserProfile(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateUserProfile(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetUserProfile(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestUserProfileMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteUserProfile
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteUserProfile{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteUserProfile{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteUserProfile{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProfileKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateUserProfile(ctx, &types.MsgCreateUserProfile{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteUserProfile(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetUserProfile(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
