package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateJobPosting{},
		&MsgUpdateJobPosting{},
		&MsgDeleteJobPosting{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProposal{},
		&MsgUpdateProposal{},
		&MsgDeleteProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProject{},
		&MsgUpdateProject{},
		&MsgDeleteProject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMilestone{},
		&MsgUpdateMilestone{},
		&MsgDeleteMilestone{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAcceptProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCompleteMilestone{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReleasePayment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDisputeProject{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
