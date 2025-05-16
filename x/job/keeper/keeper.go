package keeper

import (
	"context"
	"encoding/binary"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

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
	return prefix.NewStore(store, []byte{})
}

func (k Keeper) AppendJob(ctx context.Context, title, desc, budget, creator string) uint64 {
	count := k.GetJobCount(ctx)

	store := k.getStore(ctx)
	key := types.JobKey(count)

	job := types.Job{
		Id:          count,
		Creator:     creator,
		Title:       title,
		Description: desc,
		Budget:      budget,
	}

	bz := k.cdc.MustMarshal(&job)
	store.Set(key, bz)

	k.SetJobCount(ctx, count+1)

	return count
}

func (k Keeper) GetJobCount(ctx context.Context) uint64 {
	store := k.getBaseStore(ctx)
	bz := store.Get([]byte(types.JobCountKey))
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetJobCount(ctx context.Context, count uint64) {
	store := k.getBaseStore(ctx)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set([]byte(types.JobCountKey), bz)
}

func (k Keeper) AppendApplication(ctx context.Context, jobID uint64, applicant string, coverLetter string) error {
	store := k.getStore(ctx)
	key := types.ApplicationKey(jobID, applicant)

	if store.Has(key) {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "already applied to job %d", jobID)
	}

	app := types.Application{
		JobId:       jobID,
		Applicant:   applicant,
		CoverLetter: coverLetter,
	}

	bz := k.cdc.MustMarshal(&app)
	store.Set(key, bz)

	return nil
}



func (k Keeper) GetApplicationCount(ctx sdk.Context) uint64 {
	store := k.getStore(ctx)
	bz := store.Get([]byte(types.ApplicationCountKey))
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetApplicationCount(ctx sdk.Context, count uint64) {
	store := k.getStore(ctx)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set([]byte(types.ApplicationCountKey), bz)
}