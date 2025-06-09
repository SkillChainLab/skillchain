package types

const (
	// ModuleName defines the module name
	ModuleName = "profile"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_profile"
)

var (
	ParamsKey = []byte("p_profile")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
