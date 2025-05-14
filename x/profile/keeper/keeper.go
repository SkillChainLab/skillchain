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
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
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
		authority:    authority,
		logger:       logger,
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
	store := runtime.KVStoreAdapter(raw) // ✅ Doğru adapter
	return prefix.NewStore(store, []byte(types.ProfileKeyPrefix))
}

func (k Keeper) SetProfile(ctx context.Context, profile types.Profile) {
	store := k.getStore(ctx)
	bz, err := k.cdc.Marshal(&profile)
	if err != nil {
		panic(err)
	}
	store.Set(types.ProfileKey(profile.Username), bz) // ✅ burada key doğru
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
	store.Delete(types.ProfileKey(username)) // ✅
}
