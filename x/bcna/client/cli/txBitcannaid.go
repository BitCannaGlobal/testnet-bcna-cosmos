package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/BitCannaGlobal/testnet-bcna-cosmos/x/bcna/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdCreateBitcannaid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bitcannaid [bcnaid] [address]",
		Short: "Creates a new bitcannaid",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBcnaid := string(args[0])
			argsAddress := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBitcannaid(clientCtx.GetFromAddress().String(), string(argsBcnaid), string(argsAddress))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBitcannaid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-bitcannaid [id] [bcnaid] [address]",
		Short: "Update a bitcannaid",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsBcnaid := string(args[1])
			argsAddress := string(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBitcannaid(clientCtx.GetFromAddress().String(), id, string(argsBcnaid), string(argsAddress))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBitcannaid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-bitcannaid [id] [bcnaid] [address]",
		Short: "Delete a bitcannaid by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteBitcannaid(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
