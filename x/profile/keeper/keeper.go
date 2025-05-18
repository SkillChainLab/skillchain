package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	"github.com/SkillChainLab/skillchain/x/profile/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	storetypes "cosmossdk.io/store/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger
		storeKey     storetypes.StoreKey
		memKey       storetypes.StoreKey

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	storeKey,
	memKey storetypes.StoreKey,
	authority string,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		storeKey:     storeKey,
		memKey:       memKey,
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

func (k Keeper) getStore(ctx context.Context) prefix.Store {
	raw := k.storeService.OpenKVStore(ctx)
	store := runtime.KVStoreAdapter(raw)
	return prefix.NewStore(store, []byte(types.ProfileKeyPrefix))
}

func (k Keeper) SetProfile(ctx context.Context, profile types.Profile) {
	store := k.getStore(ctx)
	bz, err := k.cdc.Marshal(&profile)
	if err != nil {
		panic(err)
	}
	store.Set(types.ProfileKey(profile.Username), bz)
}

func (k Keeper) GetProfile(ctx context.Context, username string) (types.Profile, bool) {
	store := k.getStore(ctx)
	bz := store.Get(types.ProfileKey(username))
	if bz == nil {
		return types.Profile{}, false
	}

	var profile types.Profile
	k.cdc.MustUnmarshal(bz, &profile)
	return profile, true
}

func (k Keeper) RemoveProfile(ctx context.Context, username string) {
	store := k.getStore(ctx)
	store.Delete(types.ProfileKey(username))
}

// GetAllProfile returns all profiles
func (k Keeper) GetAllProfile(ctx context.Context) (list []types.Profile) {
	store := k.getStore(ctx)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Profile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// DeleteProfile deletes a profile
func (k Keeper) DeleteProfile(ctx context.Context, username string) {
	store := k.getStore(ctx)
	store.Delete(types.ProfileKey(username))
}

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// Set all the profile
	for _, profile := range genState.ProfileList {
		k.SetProfile(ctx, profile)
	}
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// Get all profile
	profileList := k.GetAllProfile(ctx)
	for _, profile := range profileList {
		genesis.ProfileList = append(genesis.ProfileList, profile)
	}

	return genesis
}

// RegisterInvariants registers the profile module invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	// Add any invariants here
}

// BeginBlocker is called at the beginning of every block
func BeginBlocker(ctx sdk.Context, k Keeper) {
	// Add any begin block logic here
}

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k Keeper) {
	// Add any end block logic here
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &QueryServerImpl{
		Keeper: keeper,
	}
}
