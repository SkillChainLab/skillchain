package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserProfileKeyPrefix is the prefix to retrieve all UserProfile
	UserProfileKeyPrefix = "UserProfile/value/"
)

// UserProfileKey returns the store key to retrieve a UserProfile from the index fields
func UserProfileKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
