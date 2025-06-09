package filestorage_test

import (
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	filestorage "skillchain/x/filestorage/module"
	"skillchain/x/filestorage/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FileRecordList: []types.FileRecord{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		FilePermissionList: []types.FilePermission{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FilestorageKeeper(t)
	filestorage.InitGenesis(ctx, k, genesisState)
	got := filestorage.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FileRecordList, got.FileRecordList)
	require.ElementsMatch(t, genesisState.FilePermissionList, got.FilePermissionList)
	// this line is used by starport scaffolding # genesis/test/assert
}
