package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PlatformMetricKeyPrefix is the prefix to retrieve all PlatformMetric
	PlatformMetricKeyPrefix = "PlatformMetric/value/"
)

// PlatformMetricKey returns the store key to retrieve a PlatformMetric from the index fields
func PlatformMetricKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
