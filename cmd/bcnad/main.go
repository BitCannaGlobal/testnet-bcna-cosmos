package main

import (
	"os"

	"github.com/BitCannaGlobal/testnet-bcna-cosmos/app"
	"github.com/BitCannaGlobal/testnet-bcna-cosmos/cmd/bcnad/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
