package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrPermissionDoesNotExist         = sdkerrors.Register(ModuleName, 1, "permission does not exist")
	ErrpermissionLocationDoesNotExist = sdkerrors.Register(ModuleName, 2, "permission location does not exist")
)
