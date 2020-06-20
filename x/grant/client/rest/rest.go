package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers grant-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {

	//tx routes
	//r.HandleFunc(fmt.Sprintf("/grants"), registerHandler(cliCtx)).Methods("POST")

	//query routes
	// Get a list of all grants
	//r.HandleFunc(fmt.Sprintf("/grants"), registerHandler(cliCtx)).Methods("GET")

}
