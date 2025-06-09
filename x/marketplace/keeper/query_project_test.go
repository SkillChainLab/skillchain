package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	"skillchain/x/marketplace/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestProjectQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	msgs := createNProject(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetProjectRequest
		response *types.QueryGetProjectResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProjectRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetProjectResponse{Project: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProjectRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetProjectResponse{Project: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProjectRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Project(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestProjectQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	msgs := createNProject(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProjectRequest {
		return &types.QueryAllProjectRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProjectAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Project), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Project),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProjectAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Project), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Project),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ProjectAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Project),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ProjectAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
