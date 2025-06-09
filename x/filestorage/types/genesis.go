package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FileRecordList:     []FileRecord{},
		FilePermissionList: []FilePermission{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in fileRecord
	fileRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.FileRecordList {
		index := string(FileRecordKey(elem.Index))
		if _, ok := fileRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for fileRecord")
		}
		fileRecordIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in filePermission
	filePermissionIndexMap := make(map[string]struct{})

	for _, elem := range gs.FilePermissionList {
		index := string(FilePermissionKey(elem.Index))
		if _, ok := filePermissionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for filePermission")
		}
		filePermissionIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
