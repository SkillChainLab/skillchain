package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NotificationKeyPrefix is the prefix to retrieve all Notification
	NotificationKeyPrefix = "Notification/value/"
)

// NotificationKey returns the store key to retrieve a Notification from the index fields
func NotificationKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
