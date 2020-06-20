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
		case MsgUpdatePermission:
			return HandleMsgUpdatePermission(ctx, k, msg)
		case MsgDeletePermission:
			return HandleMsgDeletePermission(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// HandleMsgCreatePermission - Handler for Creating(registering) a permission
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
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgUpdatePermission - Handler for grant a  a permission
func HandleMsgUpdatePermission(ctx sdk.Context, k Keeper, msg MsgUpdatePermission) (*sdk.Result, error) {

	//this will create or overwrite the permission registered with the data controller for data subject identified
	key := msg.Subject.String() + msg.Controller.String()
	permission, err := k.GetPermission(ctx, key)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrPermissionDoesNotExist, "Cannot locate permission ") // If not, throw an error
	}

	permission.Policy.UpdatePolicy(msg.Processor, msg.Create, msg.Read, msg.Update, msg.Delete)
	// save the permission to the permission store
	k.SetPermission(ctx, key, permission)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeUpdatePermission),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Subject.String()),
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgDeletePermission - Handler for deleting a permission
func HandleMsgDeletePermission(ctx sdk.Context, k Keeper, msg MsgDeletePermission) (*sdk.Result, error) {

	key := msg.Subject.String() + msg.Controller.String()
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
