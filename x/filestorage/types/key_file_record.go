package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FileRecordKeyPrefix is the prefix to retrieve all FileRecord
	FileRecordKeyPrefix = "FileRecord/value/"
)

// FileRecordKey returns the store key to retrieve a FileRecord from the index fields
func FileRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
