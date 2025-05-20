package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/SkillChainLab/skillchain/x/verification/types"
	profiletypes "github.com/SkillChainLab/skillchain/x/profile/types"
	jobtypes "github.com/SkillChainLab/skillchain/x/job/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		profileKeeper profiletypes.ProfileKeeper
		jobKeeper     jobtypes.JobKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	profileKeeper profiletypes.ProfileKeeper,
	jobKeeper jobtypes.JobKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		profileKeeper: profileKeeper,
		jobKeeper:     jobKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetVerifiedInstitution stores a VerifiedInstitution by address
func (k Keeper) SetVerifiedInstitution(ctx sdk.Context, institution types.VerifiedInstitution) {
	store := k.storeService.OpenKVStore(ctx)
	b := k.cdc.MustMarshal(&institution)
	store.Set(append(types.VerifiedInstitutionKeyPrefix, []byte(institution.Address)...), b)
}

// GetVerifiedInstitution retrieves a VerifiedInstitution by address
func (k Keeper) GetVerifiedInstitution(ctx sdk.Context, address string) (types.VerifiedInstitution, bool) {
	store := k.storeService.OpenKVStore(ctx)
	b, err := store.Get(append(types.VerifiedInstitutionKeyPrefix, []byte(address)...))
	if err != nil || b == nil {
		return types.VerifiedInstitution{}, false
	}
	var institution types.VerifiedInstitution
	k.cdc.MustUnmarshal(b, &institution)
	return institution, true
}

// ListVerifiedInstitutions lists all VerifiedInstitutions
func (k Keeper) ListVerifiedInstitutions(ctx sdk.Context) []types.VerifiedInstitution {
	store := k.storeService.OpenKVStore(ctx)
	iterator, err := store.Iterator(types.VerifiedInstitutionKeyPrefix, nil)
	if err != nil {
		return nil
	}
	defer iterator.Close()
	var institutions []types.VerifiedInstitution
	for ; iterator.Valid(); iterator.Next() {
		var institution types.VerifiedInstitution
		k.cdc.MustUnmarshal(iterator.Value(), &institution)
		institutions = append(institutions, institution)
	}
	return institutions
}

// SetVerificationRequest stores a VerificationRequest by request_id
func (k Keeper) SetVerificationRequest(ctx sdk.Context, req types.VerificationRequest) {
	store := k.storeService.OpenKVStore(ctx)
	b := k.cdc.MustMarshal(&req)
	store.Set(append(types.VerificationRequestKeyPrefix, []byte(req.RequestId)...), b)
}

// GetVerificationRequest retrieves a VerificationRequest by request_id
func (k Keeper) GetVerificationRequest(ctx sdk.Context, requestId string) (types.VerificationRequest, bool) {
	store := k.storeService.OpenKVStore(ctx)
	b, err := store.Get(append(types.VerificationRequestKeyPrefix, []byte(requestId)...))
	if err != nil || b == nil {
		return types.VerificationRequest{}, false
	}
	var req types.VerificationRequest
	k.cdc.MustUnmarshal(b, &req)
	return req, true
}

// ListVerificationRequests lists all VerificationRequests
func (k Keeper) ListVerificationRequests(ctx sdk.Context) []types.VerificationRequest {
	store := k.storeService.OpenKVStore(ctx)
	iterator, err := store.Iterator(types.VerificationRequestKeyPrefix, nil)
	if err != nil {
		return nil
	}
	defer iterator.Close()
	var requests []types.VerificationRequest
	for ; iterator.Valid(); iterator.Next() {
		var req types.VerificationRequest
		k.cdc.MustUnmarshal(iterator.Value(), &req)
		requests = append(requests, req)
	}
	return requests
}
