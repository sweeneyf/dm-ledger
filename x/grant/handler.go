package grant

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/sdk-tutorials/scavenge/x/scavenge/types"
)

// NewHandler creates an sdk.Handler for all the grant type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgAccessRequest:
			return handleMsgAccessRequest(ctx, k, msg)
		case MsgCreateGrant:
			return handleMsgCreateGrant(ctx, k, msg)
		// TODO: Define your msg cases
		//
		//Example:
		// case Msg<Action>:
		// 	return handleMsg<Action>(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgAccessRequest handles the access request
func handleMsgCreateGrant(ctx sdk.Context, k Keeper, msg MsgCreateGrant) (*sdk.Result, error) {

	grant := AccessControlGrant{}

	k.SetAccessControlRecord(ctx, msg.Subject.String()+msg.Controller.String()+msg.Processor.String(), grant)
	//	k.SetScavenge(ctx, scavenge)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeAccessRequest),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Processor.String()),
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
			sdk.NewAttribute(types.AttributeLocation, msg.Location),
			sdk.NewAttribute(types.AttributeAccessType, msg.AccessType),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handleMsgAccessRequest handles the access request
func handleMsgAccessRequest(ctx sdk.Context, k Keeper, msg MsgAccessRequest) (*sdk.Result, error) {

	//	k.SetScavenge(ctx, scavenge)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeAccessRequest),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Requestor.String()),
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Owner.String()),
			sdk.NewAttribute(types.AttributeLocation, msg.Location),
			sdk.NewAttribute(types.AttributeAccessType, msg.AccessType),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
