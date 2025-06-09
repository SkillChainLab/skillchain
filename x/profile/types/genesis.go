package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		UserProfileList:      []UserProfile{},
		UserSkillList:        []UserSkill{},
		SkillEndorsementList: []SkillEndorsement{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in userProfile
	userProfileIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserProfileList {
		index := string(UserProfileKey(elem.Index))
		if _, ok := userProfileIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userProfile")
		}
		userProfileIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in userSkill
	userSkillIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserSkillList {
		index := string(UserSkillKey(elem.Index))
		if _, ok := userSkillIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userSkill")
		}
		userSkillIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in skillEndorsement
	skillEndorsementIndexMap := make(map[string]struct{})

	for _, elem := range gs.SkillEndorsementList {
		index := string(SkillEndorsementKey(elem.Index))
		if _, ok := skillEndorsementIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for skillEndorsement")
		}
		skillEndorsementIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
