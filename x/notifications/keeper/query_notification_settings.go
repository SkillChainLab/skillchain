package keeper

import (
	"context"

	"skillchain/x/notifications/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NotificationSettingsAll(ctx context.Context, req *types.QueryAllNotificationSettingsRequest) (*types.QueryAllNotificationSettingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var notificationSettingss []types.NotificationSettings

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	notificationSettingsStore := prefix.NewStore(store, types.KeyPrefix(types.NotificationSettingsKeyPrefix))

	pageRes, err := query.Paginate(notificationSettingsStore, req.Pagination, func(key []byte, value []byte) error {
		var notificationSettings types.NotificationSettings
		if err := k.cdc.Unmarshal(value, &notificationSettings); err != nil {
			return err
		}

		notificationSettingss = append(notificationSettingss, notificationSettings)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNotificationSettingsResponse{NotificationSettings: notificationSettingss, Pagination: pageRes}, nil
}

func (k Keeper) NotificationSettings(ctx context.Context, req *types.QueryGetNotificationSettingsRequest) (*types.QueryGetNotificationSettingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetNotificationSettings(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNotificationSettingsResponse{NotificationSettings: val}, nil
}
