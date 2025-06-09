package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// JobPostingKeyPrefix is the prefix to retrieve all JobPosting
	JobPostingKeyPrefix = "JobPosting/value/"
)

// JobPostingKey returns the store key to retrieve a JobPosting from the index fields
func JobPostingKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
