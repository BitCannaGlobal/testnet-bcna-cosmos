# Task 5 Vote on a proposal through CLI to upgrade chain

In this task we are going to vote on a proposal through the command line interface to upgrade the chain.

We will vote for a proposal to upgrade the chain to get the latest update from Tendermint (0.34.11 to 0.34.12).
The proposal will point us to agree a block height to stop the chain and swap the binary file by consensus. You can vote whatever you want, as long as your last vote is “YES” (you can overwrite your vote). We want the proposal to pass, so please vote “YES”!


1. Vote "YES" for the proposal
The syntax is:
```
bcnad tx gov vote [id_proposal] [option] --from walletname --gas auto --gas-adjustment 1.4 --gas-prices 0.01ubcna   --chain-id bitcanna-testnet-6
```
Example:
```
bcnad tx gov vote 1 yes --gas auto --gas-adjustment 1.4 --gas-prices 0.01ubcna  --from walletname --chain-id bitcanna-testnet-6
```

In the BitCanna Explorer you can view the impact of your new decision on the proposal (optional) https://testnet-explorer.bitcanna.io/proposals/1

![](https://i.imgur.com/Y3qlWKi.png)

Once the proposal has passed the network will HALT at at block **194976**, August 23rd, around 11:00 UTC. The network upgrade will then be implemented. Tomorrow, August 20th, we will publish a guide to prepare your validator for the update.
