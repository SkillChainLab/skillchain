package types

const (
	// ModuleName defines the module name
	ModuleName = "notifications"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_notifications"
)

var (
	ParamsKey = []byte("p_notifications")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
