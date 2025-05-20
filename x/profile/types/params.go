package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	KeyMaxProfileLength = []byte("MaxProfileLength")
)

// ParamTable for profile module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(maxProfileLength uint64) Params {
	return Params{
		MaxProfileLength: maxProfileLength,
	}
}

// DefaultParams returns the default parameters for the profile module
func DefaultParams() Params {
	return NewParams(1000) // Default max profile length
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// of profile module's parameters.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxProfileLength, &p.MaxProfileLength, validateMaxProfileLength),
	}
}

// Validate performs basic validation on profile parameters
func (p Params) Validate() error {
	if err := validateMaxProfileLength(p.MaxProfileLength); err != nil {
		return err
	}
	return nil
}

func validateMaxProfileLength(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max profile length must be positive: %d", v)
	}

	return nil
}
