# STAGE 1.Task 1. Setup up your validator and join *bitcanna-testnet-2*
**bcnad** is a blockchain application built using Cosmos SDK v.0.42.4 and Tendermint v.0.34.10.


You can run the validator  software using the binary or compiling it by yourself, choose between _Step 0a_ or _Step 0b_ and continue at _Step 1_.

* Before start you might want to ensure your system is updated. You can also install a utility named `jq` to read and navigate in JSON files and outputs (other utilities could be installed too)

    ```
    sudo apt-get update
    sudo apt-get upgrade -y
    sudo apt-get install -y build-essential curl wget jq
    ```
* Increasing the default open files limit.
If we don't raise this value nodes will crash once the network grows large enough
    ```
    sudo su -c "echo 'fs.file-max = 65536' >> /etc/sysctl.conf"
    sudo sysctl -p
    ```

## Step 0A - Run a fullnode / validator using the binaries
By downloading the binary we avoid compiling the source code.
* Download the latest version (Indica revision 2 - v0.testnet6) from Github:
    ```
    cd $HOME
    wget https://github.com/BitCannaGlobal/testnet-bcna-cosmos/releases/download/v0.testnet6/bcnad
    chmod +x bcnad
    sudo mv bcnad /usr/local/bin/
    ```
* Check for the right version (v0.testnet6): 
    ```
    bcnad version
       >>> response: 0.testnet6
    ```
**NOTE:** If you have downloaded the binary avoid _Step 0B_ and jump to _Step_ 1 now
## Step 0B - Run a fullnode / validator compiling source code

The updated instructions are always in our GitHub Readme page, click on this [link](https://github.com/BitCannaGlobal/testnet-bcna-cosmos) to go there. 

## Step 1 - Setting up the connection
Instructions for setting up the connection with TestNet Blockchain
1.  **Create a wallet**:
You may create a wallet with one or more keys (addresses) using `bcnad`; you can choose the name (we advice you use one word)
```
    bcnad keys add MyFirstAddress

      name: MyFirstAddress
      type: local
      address: bcna14shzreglay98us0hep44hhhuy7dm43snv38plr
      pubkey: bcnapub1addwnpepqvtpzyugupvcu773rzdcvhele6e22txy2zr235dn7uf8t2mlqcarcyx2gg9
      mnemonic: ""
      threshold: 0
      pubkeys: []


     Important write this mnemonic phrase in a safe place.
    It is the only way to recover your account if you ever forget your password.

    deposit daring slim glide spend dolphin expire shadow cluster vivid orphan work pond section client friend yellow west hamster torch settle island opinion gloom
```
   


your address is something like this: `bcna14shzreglay98us0hep44hhhuy7dm43snv38plr`


2. **Initialize the folders:** change **_Moniker_** by your validator name (use quotes for two or more separated words *"Royal Queen Seeds"*)
    ```
    bcnad init Moniker --chain-id bitcanna-testnet-2
    ```
    This will create a `$HOME/.bcna` folder
3. **Download the Genesis** `genesis.json` file
    ```
    cd $HOME
    wget https://raw.githubusercontent.com/BitCannaGlobal/testnet-bcna-cosmos/main/instructions/stage1/genesis.json 
    mv genesis.json $HOME/.bcna/config/
    ```
4. Add to _config.toml_ file the server SEEDs:

    ```
    sed -E -i 's/seeds = \".*\"/seeds = \"d6aa4c9f3ccecb0cc52109a95962b4618d69dd3f@seed1.bitcanna.io:26656,41d373d03f93a3dc883ba4c1b9b7a781ead53d76@seed2.bitcanna.io:16656\"/' $HOME/.bcna/config/config.toml
    ```

5. Add to _config.toml_ file the _persistent_peers_: (this list could grow up the next days)

    ```
    sed -E -i 's/persistent_peers = \".*\"/persistent_peers = \"d6aa4c9f3ccecb0cc52109a95962b4618d69dd3f@seed1.bitcanna.io:26656,41d373d03f93a3dc883ba4c1b9b7a781ead53d76@seed2.bitcanna.io:16656\"/' $HOME/.bcna/config/config.toml
    ```
  

6. You can set the minimum gas prices for transactions to be accepted into your node’s mempool. This sets a lower bound on gas prices, preventing spam.
    ``` 
    sed -E -i 's/minimum-gas-prices = \".*\"/minimum-gas-prices = \"0.01ubcna\"/' $HOME/.bcna/config/app.toml
    ```

7. Open the P2P port (26656 by default)

    ```sudo ufw allow 26656```

8. Test the connection (CTRL + C to stop)

    `bcnad start --log_level info`
    ```
    3:31PM INF Committed state appHash=77D16BED3F109A4A05A971C92602029569E049DFC1DC128CFF5CCAE3158F4B1B height=3886 module=state txs=0
    3:31PM INF Indexed block height=3886 module=txindex
    3:31PM INF minted coins from module account amount=1034628bcna from=mint module=x/bank
    3:31PM INF Executed block height=3887 invalidTxs=0 module=state validTxs=0
    3:31PM INF commit synced commit=436F6D6D697449447B5B38332031333820373720313731203135362032333220313431203435203137332037372031352031363020373120393720393520352031393020313836203733203131342034322031313620313230203536203338203230203337203437203231392032353220343920385D3A4632467D
    3:31PM INF Committed state appHash=538A4DAB9CE88D2DAD4D0FA047615F05BEBA49722A7478382614252FDBFC3108 height=3887 module=state txs=0
    ```

9. **Check the sync:** while `catching_up = true` the node is syncing. Also you can compare your current block with the last synced block of another node or at our [Explorer](https://testnet-explorer.bitcanna.io):
    ```
    curl -s localhost:26657/status  | jq .result.sync_info.catching_up
    #true output is syncing - false is synced

    curl -s localhost:26657/status | jq .result.sync_info.latest_block_height
    #this output is your last block synced

    curl -s "http://seed1.bitcanna.io:26657/status?"  | jq .result.sync_info.latest_block_height
    #this output the public node last block synced
    ```
10. **Service creation**
With all configurations ready, you can start your blockchain node with a single command (`bcnad start`). In this tutorial, however, you will find a simple way to set up `systemd` to run the node daemon with auto-restart.

Setup `bcnad` systemd service (copy and paste all to create the file service):
```
    cd $HOME
    echo "[Unit]
    Description=BitCanna Node
    After=network-online.target
    [Service]
    User=${USER}
    ExecStart=$(which bcnad) start
    Restart=always
    RestartSec=3
    LimitNOFILE=4096
    [Install]
    WantedBy=multi-user.target
    " >bcnad.service
```
    
Let's enable and activate the BCNAD service.

```
    sudo mv bcnad.service /lib/systemd/system/
    sudo systemctl enable bcnad.service && sudo systemctl start bcnad.service
```
Check the logs to see if it is working?
    ```
    sudo journalctl -u bcnad -f
    ``` 

## Step 2 - Become a validator
To become a validator you need to perform additional steps. 
Your node must be fully synced in order to send the TX of validator creation and start to validate the network.

1. **Get test-coins from Discord Faucet:**
Go to #testnet-faucet channel and claim (one time) coins to your address using this syntax: 
    `!claim your_adress`
    In example: 
    `!claim bcna14shzreglay98us0hep44hhhuy7dm43snv38plr`

2. **Send the _Create validator_ TX:**

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
    --chain-id bitcanna-testnet-2 \
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
