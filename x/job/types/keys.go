package types

import (
	"fmt"
	"encoding/binary"
)

const (
	ModuleName = "job"

	StoreKey = ModuleName

	MemStoreKey = "mem_job"

	JobKeyPrefix = "Job/value/"
	JobCountKey  = "Job/count"
	ApplicationKeyPrefix   = "Application/value/"
	ApplicationCountKey    = "Application/count"
)


var ParamsKey = []byte("p_job")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func JobKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(JobKeyPrefix), bz...)
}

func ApplicationKey(jobID uint64, applicant string) []byte {
	return []byte(fmt.Sprintf("application:%d:%s", jobID, applicant))
}