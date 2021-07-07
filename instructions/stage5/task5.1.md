# STAGE 5. Task 1. Join `bitcanna-testnet-5`

In Stage 5 we are going to test certain functions related to IBC.

For your convenience we have built a new clean chain which is already live. You'll be able to join this new chain by following the instructions below.

## Step 1. Setup the environment.
Remember:  
* Cosmovisor is used for service related operations.
* `bcnad` binary (in path) is used to type commands 
* Last version of `bcnad` and `cosmovisor` are [here](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/releases/tag/v0.testnet7) (v0.testnet7).

1. **Stop your daemon** if it is running:
```
sudo systemctl stop cosmovisor
```
2. **Fetch new genesis.json** into bcna's config directory, it will automatically overwrite the previous.
```
curl -s https://raw.githubusercontent.com/BitCannaGlobal/testnet-bcna-cosmos/main/instructions/stage5/genesis.json > ~/.bcna/config/genesis.json
```

* Ensure you have right genesis with sha256sum: `3f5e49af60c9defbc5d484812e8174cc52c29aa281914f792de90ec20b2770b5` typing: 

    ```
    sha256sum  .bcna/config/genesis.json 
    ```



3. Check your seeds / persistent_peers in config.toml and replace them if necessary
```
seeds = "d6aa4c9f3ccecb0cc52109a95962b4618d69dd3f@seed1.bitcanna.io:26656,41d373d03f93a3dc883ba4c1b9b7a781ead53d76@seed2.bitcanna.io:16656,8e241ba2e8db2e83bb5d80473b4fd4d901043dda@178.128.247.173:26656"

persistent_peers = "acc177d5af9fad3064c1831bba41718d5f0ef2ce@167.71.0.53:26656,dcdc83e240eb046faabef62e4daf1cfcecfa93a2@159.65.198.245:26656"
```

4. Start the daemon and sync the network.

```
bcnad unsafe-reset-all
sudo systemctl start cosmovisor
sudo journalctl -u cosmovisor -f
```


5. **Check the sync:** while `catching_up = true` the node is syncing. Also you can compare your current block with the last synced block of another node or at our [Explorer](https://testnet-explorer.bitcanna.io):
    ```
    curl -s localhost:26657/status  | jq .result.sync_info.catching_up
    #true output is syncing - false is synced

    curl -s localhost:26657/status | jq .result.sync_info.latest_block_height
    #this output is your last block synced

    curl -s "http://seed1.bitcanna.io:26657/status?"  | jq .result.sync_info.latest_block_height
    #this output the public node last block synced
    ```
    
## Step 2. Become validator.
To become a validator you need to perform additional steps. 
Your node must be fully synced in order to send the TX of validator creation and start to validate the network.

1. **Get test-coins from Discord Faucet:**
Go to #testnet-faucet channel and claim (one time) coins to your address using this syntax: 
    `!claim your_adress`
    In example: 
    `!claim bcna14shzreglay98us0hep44hhhuy7dm43snv38plr`

Make sure to keep around 1K BCNA liquid for this Stage to make TXs.

2. **Send the _Create validator_ TX:**

> Read the FAQ manual to understand all parameters:
> https://hackmd.io/_R2KtQzAS02QXdwpJdqmnw

When you have your node synced and your wallet funded with coins (see [_Some useful commands_](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/blob/main/instructions/stage1/useful.md)) send the TX to become _validator_ (change _wallet_name_ and _moniker_):
> If you copy/paste commands, consider copying to a plain text editor to avoid errors.

> You can use quotes to include spaces and more than two words.
`--from "Royal Queen Seed"`

```
bcnad tx staking create-validator \
    --amount 1000000ubcna \
    --commission-max-change-rate 0.10 \
    --commission-max-rate 0.2 \
    --commission-rate 0.1 \
    --from WALLET \
    --min-self-delegation 1 \
    --moniker MONIKER \
    --pubkey $(bcnad tendermint show-validator) \
    --chain-id bitcanna-testnet-5 \
    --gas auto \
    --gas-adjustment 1.5 \
    --gas-prices 0.01ubcna
```

You can check the list of validators (also in [Explorer](https://testnet-explorer.bitcanna.io/validators)):

```
    bcnad query staking validators --output json| jq
```

3. Another **IMPORTANT** but **optional** action is backup your Validator_priv_key (avoid if is done in other Stage):

    ```
    tar -czvf validator_key.tar.gz .bcna/config/*_key.json 
    gpg -o validator_key.tar.gz.gpg -ca validator_key.tar.gz
    rm validator_key.tar.gz
    ```
    This will create a GPG encrypted file with both key files.
