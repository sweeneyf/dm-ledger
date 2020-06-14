package rest

// The packages below are commented out at first to prevent an error if this file isn't initially saved.
import (
	// "bytes"
	// "net/http"

	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/sweeneyf/dm-ledger/x/permission/types"
)

type registerReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Subject     string       `json:"subject"`
	Controller  string       `json:"controller"`
	DataPointer string       `json:"dataPointer"`
	DataHash    string       `json:"dataHash"`
}

func registerHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req registerReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		subjAddr, err := sdk.AccAddressFromBech32(req.Subject)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "invalid subject address")
			return
		}

		contrAddr, err := sdk.AccAddressFromBech32(req.Subject)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "invalid controller address")
			return
		}

		// create the message
		msg := types.NewMsgCreatePermission(subjAddr, contrAddr, req.DataPointer, req.DataHash)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
