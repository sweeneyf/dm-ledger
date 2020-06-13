package rest

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers permission-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	registerQueryRoutes(cliCtx, r, moduleName)
	registerTxRoutes(cliCtx, r, moduleName)

	//tx routes
	r.HandleFunc(fmt.Sprintf("/%s/register", moduleName), registerHandler(cliCtx)).Methods("Post")

	//query routes
	//r.HandleFunc(fmt.Sprintf("/%s/register", moduleName), registerHandler(cliCtx)).Methods("Post")

}
