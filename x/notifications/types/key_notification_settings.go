package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NotificationSettingsKeyPrefix is the prefix to retrieve all NotificationSettings
	NotificationSettingsKeyPrefix = "NotificationSettings/value/"
)

// NotificationSettingsKey returns the store key to retrieve a NotificationSettings from the index fields
func NotificationSettingsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
