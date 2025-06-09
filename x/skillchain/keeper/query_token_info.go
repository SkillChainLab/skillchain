package keeper

import (
	"context"

	"skillchain/x/skillchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenInfo(goCtx context.Context, req *types.QueryTokenInfoRequest) (*types.QueryTokenInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get skillchain parameters
	params := k.GetParams(ctx)

	// Get total supply from bank keeper
	totalSupply := k.bankKeeper.GetSupply(ctx, "uskill")

	// Calculate circulating supply (total supply for now, can be modified for locked tokens)
	circulatingSupply := totalSupply.Amount

	return &types.QueryTokenInfoResponse{
		Name:              params.TokenName,
		Symbol:            params.TokenSymbol,
		Decimals:          params.TokenDecimals,
		Description:       params.TokenDescription,
		TotalSupply:       totalSupply.Amount.String(),
		CirculatingSupply: circulatingSupply.String(),
		BurnedAmount:      "0", // TODO: implement burned amount tracking
		MaxSupply:         params.MaxSupply,
		BurnEnabled:       params.BurnEnabled,
		ChainDescription:  params.ChainDescription,
		WebsiteUrl:        params.WebsiteUrl,
	}, nil
}
