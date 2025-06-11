package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		VusdEnabled:         true,
		VusdMockPrice:       "0.50", // 1 SKILL = $0.50 USD
		MinCollateralRatio:  "1.50", // 150% minimum collateral ratio
		PriceUpdateAuthority: "skill1rpu3l6rlvrxgp0drjquu9p5kr0dt7lrfzxac3l", // Alice as default authority
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
