package rest

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers permission-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {

	//tx routes
	r.HandleFunc(fmt.Sprintf("/permissions"), registerHandler(cliCtx)).Methods("POST")

	//query routes
	// Get a list of all permissions
	r.HandleFunc(fmt.Sprintf("/permissions"), registerHandler(cliCtx)).Methods("GET")

}
