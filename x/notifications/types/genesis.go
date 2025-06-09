package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NotificationList:         []Notification{},
		NotificationSettingsList: []NotificationSettings{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in notification
	notificationIndexMap := make(map[string]struct{})

	for _, elem := range gs.NotificationList {
		index := string(NotificationKey(elem.Index))
		if _, ok := notificationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for notification")
		}
		notificationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in notificationSettings
	notificationSettingsIndexMap := make(map[string]struct{})

	for _, elem := range gs.NotificationSettingsList {
		index := string(NotificationSettingsKey(elem.Index))
		if _, ok := notificationSettingsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for notificationSettings")
		}
		notificationSettingsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
