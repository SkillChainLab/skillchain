package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProposalKeyPrefix is the prefix to retrieve all Proposal
	ProposalKeyPrefix = "Proposal/value/"
)

// ProposalKey returns the store key to retrieve a Proposal from the index fields
func ProposalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
