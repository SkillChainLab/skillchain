package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserSkillKeyPrefix is the prefix to retrieve all UserSkill
	UserSkillKeyPrefix = "UserSkill/value/"
)

// UserSkillKey returns the store key to retrieve a UserSkill from the index fields
func UserSkillKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
