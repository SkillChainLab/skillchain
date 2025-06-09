package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PlatformMetricList: []PlatformMetric{},
		UserActivityList:   []UserActivity{},
		RevenueRecordList:  []RevenueRecord{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in platformMetric
	platformMetricIndexMap := make(map[string]struct{})

	for _, elem := range gs.PlatformMetricList {
		index := string(PlatformMetricKey(elem.Index))
		if _, ok := platformMetricIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for platformMetric")
		}
		platformMetricIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in userActivity
	userActivityIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserActivityList {
		index := string(UserActivityKey(elem.Index))
		if _, ok := userActivityIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userActivity")
		}
		userActivityIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in revenueRecord
	revenueRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.RevenueRecordList {
		index := string(RevenueRecordKey(elem.Index))
		if _, ok := revenueRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for revenueRecord")
		}
		revenueRecordIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
