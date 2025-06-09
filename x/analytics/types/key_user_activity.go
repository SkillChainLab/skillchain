package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserActivityKeyPrefix is the prefix to retrieve all UserActivity
	UserActivityKeyPrefix = "UserActivity/value/"
)

// UserActivityKey returns the store key to retrieve a UserActivity from the index fields
func UserActivityKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
