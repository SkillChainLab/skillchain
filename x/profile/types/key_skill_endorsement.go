package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SkillEndorsementKeyPrefix is the prefix to retrieve all SkillEndorsement
	SkillEndorsementKeyPrefix = "SkillEndorsement/value/"
)

// SkillEndorsementKey returns the store key to retrieve a SkillEndorsement from the index fields
func SkillEndorsementKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
