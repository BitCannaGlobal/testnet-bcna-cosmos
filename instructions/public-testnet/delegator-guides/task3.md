 # Task 3 Vote on a proposal through web-wallet to upgrade chain

In this task we are going to test the Voting feature with the BitCanna web wallet.

We will vote for a proposal to upgrade the chain to get the latest update from Tendermint (0.34.11 to 0.34.12). 
The proposal will point us to agree a block height to stop the chain and swap the binary file by consensus. You can vote whatever you want, as long as your last vote is "YES" (you can overwrite your vote). We want the proposal to pass, so please vote "YES"!

## Casting a vote using the Web Wallet with Keplr extension.

1. Go to the BitCanna web wallet at: https://testnet-wallet.bitcanna.io/
2. Login using Keplr.
3. Go to the Proposal Section (https://testnet-wallet.bitcanna.io/proposals/1) in the BitCanna web wallet, click on the Proposal with ID-1 and click "VOTE", click "YES", and confirm your transaction.

![](https://i.imgur.com/S3WAK4z.png)

4. At the BitCanna Explorer, you can view the impact of your decision on the proposal (optional) https://testnet-explorer.bitcanna.io/proposals/1

Once the proposal has passed the network will HALT at at block 194976, August 23rd, around 11:00 UTC. The network upgrade will then be implemented. Tomorrow, August 20th, we will publish a guide for validators to prepare for the update. As a delegator, there's nothing you have to do except from monitoring your delegations. If a validator fails to upgrade and becomes jailed, you can get slashed!
