package grant

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sweeneyf/dm-ledger/x/grant/types"
)

// NewHandler creates an sdk.Handler for all the grant type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgAccessRequest:
			return handleMsgAccessRequest(ctx, k, msg)
		case MsgCreateGrant:
			return HandleMsgCreateGrant(ctx, k, msg)
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

// HandleMsgCreateGrant - Handler for Cretaing a grant
func HandleMsgCreateGrant(ctx sdk.Context, k Keeper, msg MsgCreateGrant) (*sdk.Result, error) {

	//this will create or overwrite the grant asscoiated with key generated
	key := msg.Subject.String() + msg.Controller.String() + msg.Processor.String()
	grant := &AccessControlGrant{
		Subject:    msg.Subject,
		Controller: msg.Controller,
		Processor:  msg.Processor,
		GDPRData: GDPRData{
			Location: msg.Location,
			EncrKey:  "TODO Change this to encrKey",
			Policy:   Policy{AccessType: msg.AccessType},
		},
	}
	// save the grant to the grant store
	k.SetAccessControlRecord(ctx, key, grant)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateGrant),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Subject.String()),
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
			sdk.NewAttribute(types.AttributeLocation, msg.Location),
			sdk.NewAttribute(types.AttributeAccessType, msg.AccessType),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgDeletGrant - Handler for deleting a grant
func HandleMsgDeletGrant(ctx sdk.Context, k Keeper, msg MsgDeleteGrant) (*sdk.Result, error) {

	key := msg.Subject.String() + msg.Controller.String() + msg.Processor.String()

	/*
		grant, _ := k.GetAccessControlGrant(ctx, key)
		if grant == nil {
			return nil, sdkerrors.Wrap(types.ErrGrantDoesNotExist, key)
		}
		// now remove that dataset if it exists
		if err := grant.RemoveDataset(msg.Location, "TODO Change this to encrKey"); err != nil {
			return nil, sdkerrors.Wrap(types.ErrGrantDoesNotExist, msg.Location)
		}

		// if there are no more datasets then delete the whole grant from the store
		if len(grant.Datasets) == 0 {
			k.DeleteAccessControlRecord(ctx, key)
		} */
	k.DeleteAccessControlRecord(ctx, key)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeDeleteGrant),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Processor.String()),
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
			sdk.NewAttribute(types.AttributeLocation, msg.Location),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events(), Log: fmt.Sprintf("key is %s", key)}, nil
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
