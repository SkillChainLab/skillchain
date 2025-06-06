package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateJob{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApplyJob{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReviewApplication{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateJob{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMarkNotificationAsRead{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
