package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

const (
	// ModuleName defines the module name
	ModuleName = "job"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_job"

	JobKeyPrefix         = "Job/value/"
	ApplicationKeyPrefix = "Application/value/"
)

var (
	// ModuleCdc defines the module codec
	ModuleCdc = codec.NewLegacyAmino()
)

var ParamsKey = []byte("p_job")

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	cryptocodec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

// RegisterLegacyAminoCodec registers the necessary interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateJob{}, "job/CreateJob", nil)
	cdc.RegisterConcrete(&MsgApplyJob{}, "job/ApplyJob", nil)
	cdc.RegisterConcrete(&MsgReviewApplication{}, "job/ReviewApplication", nil)
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func JobKey(id uint64) []byte {
	return []byte(fmt.Sprintf("Job/value/%d", id))
}

func ApplicationKey(jobId uint64, applicant string) []byte {
	return []byte(fmt.Sprintf("Application/value/%d:%s", jobId, applicant))
}

func JobCountKey() []byte {
	return []byte("Job/count")
}

func ApplicationCountKey() []byte {
	return []byte("Application/count")
}
