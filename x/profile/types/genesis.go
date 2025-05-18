package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:      DefaultParams(),
		ProfileList: []Profile{},
	}
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs GenesisState) Validate() error {
	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Validate profiles
	for _, profile := range gs.ProfileList {
		if profile.Username == "" {
			return fmt.Errorf("profile username cannot be empty")
		}
		if profile.Bio == "" {
			return fmt.Errorf("profile bio cannot be empty")
		}
		if profile.Creator == "" {
			return fmt.Errorf("profile creator cannot be empty")
		}
	}

	return nil
} 