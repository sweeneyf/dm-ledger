package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrGrantDoesNotExist = sdkerrors.Register(ModuleName, 1, "grant does not exist")
	ErrGrantExpired      = sdkerrors.Register(ModuleName, 2, "grant has expired")
)
