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

func (k Keeper) NotificationAll(ctx context.Context, req *types.QueryAllNotificationRequest) (*types.QueryAllNotificationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var notifications []types.Notification

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	notificationStore := prefix.NewStore(store, types.KeyPrefix(types.NotificationKeyPrefix))

	pageRes, err := query.Paginate(notificationStore, req.Pagination, func(key []byte, value []byte) error {
		var notification types.Notification
		if err := k.cdc.Unmarshal(value, &notification); err != nil {
			return err
		}

		notifications = append(notifications, notification)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNotificationResponse{Notification: notifications, Pagination: pageRes}, nil
}

func (k Keeper) Notification(ctx context.Context, req *types.QueryGetNotificationRequest) (*types.QueryGetNotificationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetNotification(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNotificationResponse{Notification: val}, nil
}
