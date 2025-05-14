package types

import "encoding/binary"

const (
	ModuleName = "job"

	StoreKey = ModuleName

	MemStoreKey = "mem_job"

	JobKeyPrefix     = "Job/value/"
	JobCountKey      = "Job/count"
)

var ParamsKey = []byte("p_job")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func JobKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, id)
	return append([]byte(JobKeyPrefix), b...)
}