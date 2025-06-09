package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MilestoneKeyPrefix is the prefix to retrieve all Milestone
	MilestoneKeyPrefix = "Milestone/value/"
)

// MilestoneKey returns the store key to retrieve a Milestone from the index fields
func MilestoneKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
