package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FilePermissionKeyPrefix is the prefix to retrieve all FilePermission
	FilePermissionKeyPrefix = "FilePermission/value/"
)

// FilePermissionKey returns the store key to retrieve a FilePermission from the index fields
func FilePermissionKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
