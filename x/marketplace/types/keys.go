package types

const (
	// ModuleName defines the module name
	ModuleName = "marketplace"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_marketplace"
)

var (
	ParamsKey = []byte("p_marketplace")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
