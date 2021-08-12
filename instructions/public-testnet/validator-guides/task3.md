# Task 3 Perform all types of transactions & commands

In this task, we urge you to play around in the cli environment by performing different kinds of transactions & commands. You might encounter a handful of commands day-to-day as a validator or delegator working in a cli-environment. To perform them, you'll need to write add specific parameters to make the transaction go through. 

We listed some of the most common transactions in this document called [useful commands](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/blob/main/instructions/public-testnet/validator-guides/useful.md) 

To navigate through the command-line, you can use: `bcnad --help` or more specifically `bcnad tx --help`.

You should be able to figure out other commands and transactions not listed in the document linked above on your own, since they will basically have the same structure. You can always write `--help` behind a command in the cli environment for additional instructions.

Example of a delegation transaction: (replace 'bcnavaloperxxxx' address to your target validator, and replace 'walletname' to your wallet name)

```
bcnad tx staking delegate bcnavaloper1dpa0fkmjf3taqnc4fu43l2lgektt3caea8pfwx 100000000ubcna --from walletname --gas-adjustment 1.5 --gas auto --gas-prices 0.01ubcna --chain-id bitcanna-testnet-6
```

Good luck!
