package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

const (
	JobApplicationKeyPrefix = "JobApplication/value/"
)

// SetJobApplication sets a job application in the store
func (k Keeper) SetJobApplication(ctx context.Context, application types.Application) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(JobApplicationKeyPrefix))

	b := k.cdc.MustMarshal(&application)
	store.Set(types.KeyPrefix(fmt.Sprintf("%d:%s", application.JobId, application.Applicant)), b)
}

// GetJobApplication gets a job application from the store
func (k Keeper) GetJobApplication(ctx context.Context, jobId uint64, applicant string) (types.Application, bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(JobApplicationKeyPrefix))

	var application types.Application
	b := store.Get(types.KeyPrefix(fmt.Sprintf("%d:%s", jobId, applicant)))
	if b == nil {
		return application, false
	}

	k.cdc.MustUnmarshal(b, &application)
	return application, true
}

// GetJobApplications gets all applications for a job
func (k Keeper) GetJobApplications(ctx context.Context, jobId uint64) []types.Application {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(JobApplicationKeyPrefix))

	var applications []types.Application
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var application types.Application
		k.cdc.MustUnmarshal(iterator.Value(), &application)
		if application.JobId == jobId {
			applications = append(applications, application)
		}
	}

	return applications
} 