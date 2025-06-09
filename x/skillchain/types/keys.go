package types

const (
	// ModuleName defines the module name
	ModuleName = "skillchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_skillchain"
)

var (
	ParamsKey = []byte("p_skillchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
