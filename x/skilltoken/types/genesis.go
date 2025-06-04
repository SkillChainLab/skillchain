package types

import (
	"fmt"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		Tokens: []*Token{
			{
				Symbol:       "SKILL",
				Name:         "SkillChain Token",
				TotalSupply:  "1000000000", // 1 billion tokens as string
				Decimals:     6,
				Owner:        "", // Will be set to the first validator
				Mintable:     true,
				Burnable:     true,
				Transferable: true,
			},
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Validate tokens
	tokenSymbols := make(map[string]bool)
	for _, token := range gs.Tokens {
		if tokenSymbols[token.Symbol] {
			return fmt.Errorf("duplicate token symbol: %s", token.Symbol)
		}
		tokenSymbols[token.Symbol] = true

		// Check if total supply is a valid integer and not negative
		if token.TotalSupply == "" {
			return fmt.Errorf("token %s total supply cannot be empty", token.Symbol)
		}
		// Optionally, parse to int and check negative
		// (skip for simplicity, or use strconv.ParseInt if needed)

		if token.Decimals > 18 {
			return fmt.Errorf("token %s decimals cannot be greater than 18", token.Symbol)
		}
	}

	return nil
}
