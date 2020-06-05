package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrGrantDoesNotExist         = sdkerrors.Register(ModuleName, 1, "grant does not exist")
	ErrGrantLocationDoesNotExist = sdkerrors.Register(ModuleName, 2, "grant location does not exist")
)
