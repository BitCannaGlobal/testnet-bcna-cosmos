package bcna

import (
	"github.com/BitCannaGlobal/bcna/x/bcna/keeper"
	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the bitcannaid
	for _, elem := range genState.BitcannaidList {
		k.SetBitcannaid(ctx, *elem)
	}

	// Set bitcannaid count
	k.SetBitcannaidCount(ctx, uint64(len(genState.BitcannaidList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all bitcannaid
	bitcannaidList := k.GetAllBitcannaid(ctx)
	for _, elem := range bitcannaidList {
		elem := elem
		genesis.BitcannaidList = append(genesis.BitcannaidList, &elem)
	}

	return genesis
}
