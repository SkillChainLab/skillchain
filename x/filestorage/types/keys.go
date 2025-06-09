package types

const (
	// ModuleName defines the module name
	ModuleName = "filestorage"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_filestorage"
)

var (
	ParamsKey = []byte("p_filestorage")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
