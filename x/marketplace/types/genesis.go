package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		JobPostingList: []JobPosting{},
		ProposalList:   []Proposal{},
		ProjectList:    []Project{},
		MilestoneList:  []Milestone{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in jobPosting
	jobPostingIndexMap := make(map[string]struct{})

	for _, elem := range gs.JobPostingList {
		index := string(JobPostingKey(elem.Index))
		if _, ok := jobPostingIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for jobPosting")
		}
		jobPostingIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in proposal
	proposalIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposalList {
		index := string(ProposalKey(elem.Index))
		if _, ok := proposalIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposal")
		}
		proposalIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in project
	projectIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProjectList {
		index := string(ProjectKey(elem.Index))
		if _, ok := projectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for project")
		}
		projectIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in milestone
	milestoneIndexMap := make(map[string]struct{})

	for _, elem := range gs.MilestoneList {
		index := string(MilestoneKey(elem.Index))
		if _, ok := milestoneIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for milestone")
		}
		milestoneIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
