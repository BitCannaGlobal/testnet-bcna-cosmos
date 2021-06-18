# Alternative join to `bitcanna-testnet-3`, how become validator after 17-06-2021 14h UTC

New Testnet chain-id `bitcanna-testnet-3` was scheduled to be launched at 17th of June - 14:00h UTC.
If you aren't in the Genesis registered as validator, you can run a new one. This guide will describe exactly how.

## Few quick steps to join if your validator isn't in the Genesis file.
> Time interval: After 17th Thursday  14.00h UTC if YOU DON'T RUN A VALIDATOR ALREADY

## Step 1. Setup the environment.
Step 1 can be skipped if you already followed [the guide for Stage 2 - Task 2]([https://](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/blob/main/instructions/stage2/task2.2.md))

1. **Stop your daemon** if it is running:
```
sudo systemctl stop bcnad
```
2. **Fetch new genesis.json** into bcna's config directory, it will automatically overwrite the previous.
```
curl -s https://raw.githubusercontent.com/BitCannaGlobal/testnet-bcna-cosmos/main/instructions/stage2/genesis.json > ~/.bcna/config/genesis.json
```

3. **Ensure you retrieve the right file**. The sha256sum must be `5d5f635660d8634e1a18002cba3cc68d3fcf60f8b2c1499569f13c91c1e10387`

```
sha256sum ~/.bcna/config/genesis.json 

    > expected output:  5d5f635660d8634e1a18002cba3cc68d3fcf60f8b2c1499569f13c91c1e10387  /home/user/.bcna/config/genesis.json
``` 
4. Start the daemon and sync the network.

```
bcnad unsafe-reset-all
sudo systemctl start bcnad
sudo journalctl -u bcnad -f
```

## Step2. Become validator after 17-Jun 14h UTC 
1. **Check the sync:** while `catching_up = true` the node is syncing. Also you can compare your current block with the last synced block of another node or at our [Explorer](https://testnet-explorer.bitcanna.io):
    ```
    curl -s localhost:26657/status  | jq .result.sync_info.catching_up
    #true output is syncing - false is synced

    curl -s localhost:26657/status | jq .result.sync_info.latest_block_height
    #this output is your last block synced

    curl -s "http://seed1.bitcanna.io:26657/status?"  | jq .result.sync_info.latest_block_height
    #this output the public node last block synced
    ```
2. **Check if you have coins, if not ask an admin ;)**
To become a validator you need testnet coins. Check it at our [Explorer.](https://testnet-explorer.bitcanna.io)
You can also check from your cli: 
    ``` 
    bcnad query bank balances bcnaAddress
    ```
3. **Send the _Create validator_ TX:**

> Read the FAQ manual to understand all parameters:
> https://hackmd.io/_R2KtQzAS02QXdwpJdqmnw

When you have your node synced and your wallet funded with coins (see [_Some useful commands_](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/blob/main/instructions/stage1/useful.md)) send the TX to become _validator_ (change _wallet_name_ and _moniker_):
> You can use quotes to include spaces and more than two words
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
    --chain-id bitcanna-testnet-3 \
    --gas auto \
    --gas-adjustment 1.5 \
    --gas-prices 0.01ubcna
```

You can check the list of validators (also in [Explorer](https://testnet-explorer.bitcanna.io/validators)):

```
    bcnad query staking validators --output json| jq
```

3. Another **IMPORTANT** but **optional** action is backup your Validator_priv_key:

    ```
    tar -czvf validator_key.tar.gz .bcna/config/*_key.json 
    gpg -o validator_key.tar.gz.gpg -ca validator_key.tar.gz
    rm validator_key.tar.gz
    ```
    This will create a GPG encrypted file with both key files.
