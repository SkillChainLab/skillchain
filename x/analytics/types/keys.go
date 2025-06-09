package types

const (
	// ModuleName defines the module name
	ModuleName = "analytics"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_analytics"
)

var (
	ParamsKey = []byte("p_analytics")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
