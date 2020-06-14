package permission

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sweeneyf/dm-ledger/x/permission/types"
)

// NewHandler creates an sdk.Handler for all the permission type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreatePermission:
			return HandleMsgCreatePermission(ctx, k, msg)
		case MsgAccessRequest:
			return handleMsgAccessRequest(ctx, k, msg)
		case MsgDeletePermission:
			return HandleMsgDeletePermission(ctx, k, msg)
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

// HandleMsgCreatePermission - Handler for Cretaing a permission
func HandleMsgCreatePermission(ctx sdk.Context, k Keeper, msg MsgCreatePermission) (*sdk.Result, error) {

	//this will create or overwrite the permission registered with the data controller for data subject identified
	key := msg.Subject.String() + msg.Controller.String()
	permission := NewPermission(msg.Subject, msg.Controller, msg.DataPointer, msg.DataHash)
	// save the permission to the permission store
	k.SetPermission(ctx, key, &permission)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreatePermission),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Subject.String()),
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
			sdk.NewAttribute(types.AttributeDataPointer, msg.DataPointer),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgDeletePermission - Handler for deleting a permission
func HandleMsgDeletePermission(ctx sdk.Context, k Keeper, msg MsgDeletePermission) (*sdk.Result, error) {

	key := msg.Subject.String() + msg.Controller.String()

	/*
		permission, _ := k.GetPermission(ctx, key)
		if permission == nil {
			return nil, sdkerrors.Wrap(types.ErrpermissionDoesNotExist, key)
		}
		// now remove that dataset if it exists
		if err := permission.RemoveDataset(msg.Location, "TODO Change this to encrKey"); err != nil {
			return nil, sdkerrors.Wrap(types.ErrpermissionDoesNotExist, msg.Location)
		}

		// if there are no more datasets then delete the whole permission from the store
		if len(permission.Datasets) == 0 {
			k.DeletePermission(ctx, key)
		} */
	k.DeletePermission(ctx, key)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeDeletePermission),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
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
			sdk.NewAttribute(types.AttributeDataPointer, msg.Location),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
