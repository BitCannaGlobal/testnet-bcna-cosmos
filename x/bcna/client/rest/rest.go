package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers bcna-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/bcna/bitcannaids/{id}", getBitcannaidHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/bcna/bitcannaids", listBitcannaidHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/bcna/bitcannaids", createBitcannaidHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bcna/bitcannaids/{id}", updateBitcannaidHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bcna/bitcannaids/{id}", deleteBitcannaidHandler(clientCtx)).Methods("POST")

}
