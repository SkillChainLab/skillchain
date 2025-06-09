package keeper

import (
	"context"

	"skillchain/x/analytics/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePlatformMetric(goCtx context.Context, msg *types.MsgCreatePlatformMetric) (*types.MsgCreatePlatformMetricResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetPlatformMetric(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var platformMetric = types.PlatformMetric{
		Creator:     msg.Creator,
		Index:       msg.Index,
		MetricName:  msg.MetricName,
		MetricValue: msg.MetricValue,
		MetricType:  msg.MetricType,
		Period:      msg.Period,
		Timestamp:   msg.Timestamp,
		Metadata:    msg.Metadata,
	}

	k.SetPlatformMetric(
		ctx,
		platformMetric,
	)
	return &types.MsgCreatePlatformMetricResponse{}, nil
}

func (k msgServer) UpdatePlatformMetric(goCtx context.Context, msg *types.MsgUpdatePlatformMetric) (*types.MsgUpdatePlatformMetricResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPlatformMetric(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var platformMetric = types.PlatformMetric{
		Creator:     msg.Creator,
		Index:       msg.Index,
		MetricName:  msg.MetricName,
		MetricValue: msg.MetricValue,
		MetricType:  msg.MetricType,
		Period:      msg.Period,
		Timestamp:   msg.Timestamp,
		Metadata:    msg.Metadata,
	}

	k.SetPlatformMetric(ctx, platformMetric)

	return &types.MsgUpdatePlatformMetricResponse{}, nil
}

func (k msgServer) DeletePlatformMetric(goCtx context.Context, msg *types.MsgDeletePlatformMetric) (*types.MsgDeletePlatformMetricResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPlatformMetric(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePlatformMetric(
		ctx,
		msg.Index,
	)

	return &types.MsgDeletePlatformMetricResponse{}, nil
}
