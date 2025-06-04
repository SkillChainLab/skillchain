package types

const (
	// ModuleName defines the module name
	ModuleName = "skilltoken"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_skilltoken"

	// TokenKeyPrefix is the prefix for storing tokens
	TokenKeyPrefix = "Token/value/"

	// TokenDenomPrefix is the prefix for token denoms
	TokenDenomPrefix = "skill/"
)

var (
	ParamsKey = []byte("p_skilltoken")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// TokenKey returns the key for storing a token
func TokenKey(symbol string) []byte {
	return []byte(TokenKeyPrefix + symbol)
}

// GetTokenDenom returns the denom for a token
func GetTokenDenom(symbol string) string {
	return TokenDenomPrefix + symbol
}
