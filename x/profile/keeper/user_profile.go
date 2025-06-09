package keeper

import (
	"context"
	"fmt"
	"time"

	"skillchain/x/profile/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetUserProfile set a specific userProfile in the store from its index
func (k Keeper) SetUserProfile(ctx context.Context, userProfile types.UserProfile) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserProfileKeyPrefix))
	b := k.cdc.MustMarshal(&userProfile)
	store.Set(types.UserProfileKey(
		userProfile.Index,
	), b)
}

// GetUserProfile returns a userProfile from its index
func (k Keeper) GetUserProfile(
	ctx context.Context,
	index string,

) (val types.UserProfile, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserProfileKeyPrefix))

	b := store.Get(types.UserProfileKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserProfile removes a userProfile from the store
func (k Keeper) RemoveUserProfile(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserProfileKeyPrefix))
	store.Delete(types.UserProfileKey(
		index,
	))
}

// GetAllUserProfile returns all userProfile
func (k Keeper) GetAllUserProfile(ctx context.Context) (list []types.UserProfile) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserProfileKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserProfile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CalculateUserReputation calculates user reputation based on endorsements received
func (k Keeper) CalculateUserReputation(ctx context.Context, userAddress string) uint64 {
	baseReputation := uint64(100) // Starting reputation
	endorsementBonus := uint64(0)
	personalStakingBonus := uint64(0)

	// Get all endorsements for this user
	allEndorsements := k.GetAllSkillEndorsement(ctx)

	for _, endorsement := range allEndorsements {
		if endorsement.TargetUser == userAddress {
			// Different weights for different endorsement types
			var baseWeight uint64
			switch endorsement.EndorsementType {
			case "strong":
				baseWeight = 20
			case "moderate":
				baseWeight = 10
			case "basic":
				baseWeight = 5
			default:
				baseWeight = 5
			}

			// Enhanced endorser reputation weighting system
			endorserProfile, found := k.GetUserProfile(ctx, endorsement.Endorser)
			var endorserMultiplier float64 = 1.0

			if found {
				endorserRep := endorserProfile.ReputationScore

				// Tiered reputation multiplier system
				switch {
				case endorserRep >= 500: // Master level
					endorserMultiplier = 2.5
				case endorserRep >= 300: // Expert level
					endorserMultiplier = 2.0
				case endorserRep >= 200: // Advanced level
					endorserMultiplier = 1.5
				case endorserRep >= 150: // Intermediate level
					endorserMultiplier = 1.2
				case endorserRep >= 100: // Beginner level
					endorserMultiplier = 1.0
				default: // Below baseline
					endorserMultiplier = 0.7
				}
			}

			// Add token staking bonus if tokens are staked
			stakingMultiplier := 1.0
			if endorsement.SkillTokensStaked > 0 {
				// Each staked SKILL token adds 10% bonus (max 100% bonus for 10+ tokens)
				stakingBonus := float64(endorsement.SkillTokensStaked) * 0.1
				if stakingBonus > 1.0 {
					stakingBonus = 1.0 // Cap at 100% bonus
				}
				stakingMultiplier = 1.0 + stakingBonus
			}

			// Calculate final weighted bonus
			finalWeight := float64(baseWeight) * endorserMultiplier * stakingMultiplier
			endorsementBonus += uint64(finalWeight)
		}

		// NEW: Calculate personal staking bonus for tokens this user has staked
		if endorsement.Endorser == userAddress && endorsement.SkillTokensStaked > 0 {
			// Personal staking bonus: 2 reputation points per staked token
			// This rewards users for putting their money where their mouth is
			personalStakingBonus += endorsement.SkillTokensStaked * 2
		}
	}

	totalReputation := baseReputation + endorsementBonus + personalStakingBonus

	// Cap reputation at reasonable maximum
	if totalReputation > 1000 {
		totalReputation = 1000
	}

	return totalReputation
}

// UpdateUserReputation updates a user's reputation score
func (k Keeper) UpdateUserReputation(ctx context.Context, userAddress string) error {
	profile, found := k.GetUserProfile(ctx, userAddress)
	if !found {
		return fmt.Errorf("user profile not found: %s", userAddress)
	}

	newReputation := k.CalculateUserReputation(ctx, userAddress)
	profile.ReputationScore = newReputation
	profile.UpdatedAt = uint64(time.Now().Unix())

	k.SetUserProfile(ctx, profile)
	return nil
}
