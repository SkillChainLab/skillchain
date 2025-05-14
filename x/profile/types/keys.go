package types

const (
	// ModuleName defines the module name
	ModuleName = "profile"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_profile"

	// ProfileKeyPrefix defines the prefix for storing Profiles
	ProfileKeyPrefix = "Profile/value/"
)

var (
	ParamsKey = []byte("p_profile")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func ProfileKey(username string) []byte {
	return []byte(ProfileKeyPrefix + username)
}
