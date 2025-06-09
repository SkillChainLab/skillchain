package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RevenueRecordKeyPrefix is the prefix to retrieve all RevenueRecord
	RevenueRecordKeyPrefix = "RevenueRecord/value/"
)

// RevenueRecordKey returns the store key to retrieve a RevenueRecord from the index fields
func RevenueRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
