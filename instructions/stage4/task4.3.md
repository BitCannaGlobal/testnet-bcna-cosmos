# ALTERNATIVE: STAGE 4. Task 3. Cast your vote.

In this task we are going to test the Voting feature with the BitCanna Web Wallet with Keplr extension, and with CLI (command line interface)

We will vote for a proposal to upgrade the chain. The proposal will point us to agree a block height to stop the chain and swap the binary file by consensus. You will first vote "NO" with KEPLR in the web wallet for testing purposes, and directly after you will vote "YES" through CLI to overwrite your vote, because we need the proposal to pass.

## Casting a vote using the Web Wallet with Keplr extension.

1. Go to the BitCanna web wallet at: https://testnet-wallet.bitcanna.io/welcome
2. Login using Keplr.
3. Go to the Proposal Section (https://testnet-wallet.bitcanna.io/proposals/2) on BitCanna Web Wallet, click on the Proposal with ID-2 (ID-1 won't work) and click "VOTE", and then click Vote "NO". In the Memo field in Keplr write: `voted NO with KEPLR` If you fail to vote "NO" with KEPLR you will be awarded with 0 points.

![](https://i.imgur.com/S3WAK4z.png)

![](https://i.imgur.com/DkWawvC.png)


4. Submit the TX HASH from your voting TX through Keplr in the Task Center, in the "proof" field for task 4.3.

Example: `https://testnet-explorer.bitcanna.io/transactions/F03C65E1F3A788FDA5D034D3B2634F2E8BAA165C0645F60E8B941F3BBC72DA04`

5. At the BitCanna Explorer, you can view the impact of your decision on the proposal (optional) https://testnet-explorer.bitcanna.io/proposals/2
![](https://i.imgur.com/y4iW44G.png)

6. Now it's time to vote again, but from CLI, and this time you must vote "YES" as shown in the example below. If you fail to vote "YES" with CLI you will be awarded with 0 points.
The syntax is:
```
bcnad tx gov vote [id_proposal] [option] --from walletname --gas auto --gas-adjustment 1.4 --gas-prices 0.01ubcna   --chain-id bitcanna-testnet-4
```
Example:
```
bcnad tx gov vote 2 yes --gas auto --gas-adjustment 1.4 --gas-prices 0.01ubcna  --from walletname --chain-id bitcanna-testnet-4
```
7. Submit the TX HASH from your voting TX through CLI in the Task Center, in the "proof" field for task 4.3.
Example: `https://testnet-explorer.bitcanna.io/transactions/59B441F7199FB5DFEF5D976661DA7A8804A05D2F2682FF7CD09B5F6FAB556590`

In the BitCanna Explorer you can view the impact of your new decision on the proposal (optional) https://testnet-explorer.bitcanna.io/proposals/2

![](https://i.imgur.com/Y3qlWKi.png)

Once the proposal has passed the network will HALT at the 5th of July, around 14:00 UTC, at block 58576. The network upgrade will then be implemented. Stay tuned for more details!
