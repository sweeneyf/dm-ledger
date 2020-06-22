package grant

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sweeneyf/dm-ledger/util"
	"github.com/sweeneyf/dm-ledger/x/grant/types"
)

// NewHandler creates an sdk.Handler for all the grant type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgRequestAccess:
			return HandleMsgRequestAccess(ctx, k, msg)
		case types.MsgValidateToken:
			return HandleMsgValidateToken(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// HandleMsgRequestAccess handles the access request
func HandleMsgRequestAccess(ctx sdk.Context, k Keeper, msg MsgRequestAccess) (*sdk.Result, error) {

	key := msg.Subject.String() + msg.Controller.String()
	permission, err := k.PermissionKeeper.GetPermission(ctx, key)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrGrantDoesNotExist, "Cannot locate grant ") // If not, throw an error
	}

	token := util.GenerateUUID()
	//reqTime := ctx.BlockHeader().Time

	grant := Grant{
		Token:   token,
		Created: time.Now(),
		Status:  0,
		Expires: time.Now().Add(3600 * time.Second),
		Create:  permission.Policy.FindAccInACL(permission.Policy.Create, msg.Processor) > 0,
		Read:    permission.Policy.FindAccInACL(permission.Policy.Read, msg.Processor) > 0,
		Update:  permission.Policy.FindAccInACL(permission.Policy.Update, msg.Processor) > 0,
		Delete:  permission.Policy.FindAccInACL(permission.Policy.Delete, msg.Processor) > 0,
	}

	k.SetGrant(ctx, grant.Token, &grant)

	grantBz := types.ModuleCdc.MustMarshalBinaryLengthPrefixed(grant)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRequestAccess),
			sdk.NewAttribute(types.AttributeProcessor, msg.Processor.String()),
			sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
		),
	)
	return &sdk.Result{Data: grantBz, Events: ctx.EventManager().Events()}, nil
}

// HandleMsgValidateToken handles the access request
func HandleMsgValidateToken(ctx sdk.Context, k Keeper, msg MsgValidateToken) (*sdk.Result, error) {

	key := msg.Token
	grant, err := k.GetGrant(ctx, key)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrGrantDoesNotExist, "Cannot locate grant ") // If not, throw an error
	}

	if grant.Expires.Before(time.Now()) {
		return nil, sdkerrors.Wrap(types.ErrGrantExpired, "Grant has expired ")
	}

	grantBz := types.ModuleCdc.MustMarshalBinaryLengthPrefixed(grant)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeValidateToken),
			sdk.NewAttribute(types.AttributeProcessor, msg.Token),
		),
	)
	return &sdk.Result{Data: grantBz, Events: ctx.EventManager().Events()}, nil
}
