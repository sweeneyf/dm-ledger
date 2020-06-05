package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAccessRequest{}, "grant/AccessRequest", nil)
	cdc.RegisterConcrete(MsgCreateGrant{}, "grant/CreateGrant", nil)
	cdc.RegisterConcrete(MsgDeleteGrant{}, "grant/DeleteGrant", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
