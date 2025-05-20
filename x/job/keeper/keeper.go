package keeper

import (
	"context"
	"encoding/binary"
	"fmt"

	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

// Ensure Keeper implements the JobKeeper interface
var _ types.JobKeeper = Keeper{}

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger
		authority    string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		logger:       logger,
		authority:    authority,
	}
}

func (k Keeper) GetAuthority() string {
	return k.authority
}

func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) getStore(ctx context.Context) prefix.Store {
	raw := k.storeService.OpenKVStore(ctx)
	store := runtime.KVStoreAdapter(raw)
	return prefix.NewStore(store, []byte(types.JobKeyPrefix))
}

func (k Keeper) getBaseStore(ctx context.Context) prefix.Store {
	raw := k.storeService.OpenKVStore(ctx)
	store := runtime.KVStoreAdapter(raw)
	return prefix.NewStore(store, []byte(types.StoreKey))
}

func (k Keeper) GetJob(ctx context.Context, id uint64) (types.Job, bool) {
	store := k.getStore(ctx)
	key := types.JobKey(id)
	fmt.Printf("DEBUG: Looking for job with key: %x\n", key)

	bz := store.Get(key)
	if bz == nil {
		fmt.Printf("DEBUG: No job found for ID: %d\n", id)
		return types.Job{}, false
	}

	var job types.Job
	k.cdc.MustUnmarshal(bz, &job)

	// Ensure ID is set
	job.Id = id

	fmt.Printf("DEBUG: Found job: %+v\n", job)
	return job, true
}

func (k Keeper) AppendJob(ctx context.Context, title, desc, budget, creator string) uint64 {
	count := k.GetJobCount(ctx)
	jobID := count + 1 // Start from 1 instead of 0

	store := k.getStore(ctx)
	key := types.JobKey(jobID)

	job := types.Job{
		Id:          jobID,
		Creator:     creator,
		Title:       title,
		Description: desc,
		Budget:      budget,
	}

	bz := k.cdc.MustMarshal(&job)
	store.Set(key, bz)

	k.SetJobCount(ctx, jobID)

	return jobID
}

func (k Keeper) GetJobCount(ctx context.Context) uint64 {
	store := k.getBaseStore(ctx)
	bz := store.Get(types.JobCountKey())
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetJobCount(ctx context.Context, count uint64) {
	store := k.getBaseStore(ctx)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(types.JobCountKey(), bz)
}

func (k Keeper) AppendApplication(ctx context.Context, jobID uint64, applicant string, coverLetter string) error {
	// First verify that the job exists
	job, found := k.GetJob(ctx, jobID)
	if !found {
		return errorsmod.Wrapf(types.ErrJobNotFound, "job %d not found", jobID)
	}

	store := k.getStore(ctx)
	key := []byte(fmt.Sprintf("Application/value/%d:%s", jobID, applicant))

	if store.Has(key) {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "already applied to job %d", jobID)
	}

	app := types.Application{
		JobId:          jobID,
		Applicant:      applicant,
		CoverLetter:    coverLetter,
		Status:         "PENDING",
		JobTitle:       job.Title,
		JobDescription: job.Description,
		JobBudget:      job.Budget,
	}

	bz := k.cdc.MustMarshal(&app)
	store.Set(key, bz)

	return nil
}

func (k Keeper) GetApplicationCount(ctx context.Context) uint64 {
	store := k.getBaseStore(ctx)
	bz := store.Get(types.ApplicationCountKey())
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetApplicationCount(ctx context.Context, count uint64) {
	store := k.getBaseStore(ctx)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(types.ApplicationCountKey(), bz)
}

func (k Keeper) SetJob(ctx context.Context, job types.Job) {
	store := k.getStore(ctx)
	key := types.JobKey(job.Id)
	bz := k.cdc.MustMarshal(&job)
	store.Set(key, bz)
}
