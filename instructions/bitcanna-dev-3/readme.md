# Devnet-3: Setup up your validator and join *bitcanna-dev-3*
> IMPORTANT NOTE: If you participated in the previous BitCanna testnets delete the bitcanna folder and start a new fresh installation:
> `rm -rf ~/.bcna`

## Target of this DevNet
### We are going to work in two new testnets: 
* bitcanna-dev-3 (with `bcnad` current version v.1.2 that we will update later)
* osmosis-dev-1 (with `osmosisd` current version binary) YOU DON'T NEED TO RUN THIS CHAIN.

**These chains will be interconnected, creating new IBC channels and testing the recovering of closed channels and stucked clients
Also we will test how to upgrade IBC channels when the remote chain upgrade their chain-id.**

### BitCanna Team is going to run a relayer to make the "official" test, but anyone can run their own relayer and create his/her own channels:
* BINARY: https://github.com/informalsystems/ibc-rs/releases/tag/v0.10.0
* INSTRUCTIONS: https://hermes.informal.systems/

### Data for config relayer:
```
omosis-dev-1 remote host to make/maintain IBC connections:

host: 188.166.126.81
rpc:  188.166.126.81:36657
grpc: 188.166.126.81:9093
lcd:  188.166.126.81:3317 (api-lcd)

bitcanna-dev-3 remote host to make/maintain IBC connections:

host: 188.166.126.81
rpc:  188.166.126.81:26657
grpc: 188.166.126.81:9090
lcd:  188.166.126.81:1317 (api-lcd)
```

## Running a validator on **bitcanna-dev-3**

You can run the validator software using the binary or compiling it by yourself, you can choose between _Step 0a_ or _Step 0b_ and continue at _Step 1_.

* Before you start, you might want to ensure your system is updated. You can also install a utility named `jq` to read and navigate in JSON files and outputs (other utilities could be installed too)

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
* Download the latest version (v1.2) from Github:
    ```
    cd $HOME
    wget https://github.com/BitCannaGlobal/bcna/releases/download/v1.2/bcnad
    chmod +x bcnad
    sudo mv bcnad /usr/local/bin/
    ```
* Check for the right version (v1.2): 
    ```
    bcnad version
       >>> response: 1.2
    ```
**NOTE:** If you have downloaded the binary avoid _Step 0B_ and go to _Step_ 1
## Step 0B - Run a fullnode / validator compiling source code (not recommended for new users)

The updated instructions are always in our GitHub Readme page, click on this [link](https://github.com/BitCannaGlobal/bcna) to go there. 

## Step 1 - Setting up the connection
Instructions for setting up the connection with the BitCanna Public Testnet Blockchain
1. **Set the chain-id param** 
```
    bcnad config chain-id bitcanna-dev-3
```
2.  **Create a wallet**:
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

    deposit daring slim glide spend dolphin expire shadow cluster weed orphan work 420 section client friend yellow west hamster torch settle island opinion gloom
```
   


Your address will look something similar like this: `bcna14shzreglay98us0hep44hhhuy7dm43snv38plr`


3. **Initialize the folders:** change **_Moniker_** by your validator name (use quotes for two or more separated words *"Royal Queen Seeds"*)
    ```
    bcnad init Moniker --chain-id bitcanna-dev-3
    ```
    This will create a `$HOME/.bcna` folder
4. **Download the Genesis** `genesis.json` file
    ```
    cd $HOME
    curl -s https://raw.githubusercontent.com/BitCannaGlobal/testnet-bcna-cosmos/main/instructions/bitcanna-dev-3/genesis.json > ~/.bcna/config/genesis.json
    ```
   Ensure you have the correct file. Run the SHA256SUM test:
    ```
     sha256sum $HOME/.bcna/config/genesis.json
     <output> fcdc71e952da0ea591fb71c2587d25a6182bac69025e1753b86db2469a432aab
    ```
5. **Add to _config.toml_ file the server SEEDs:**

    ```
    sed -E -i 's/seeds = \".*\"/seeds = \"3ec664ad43c88b22ebc738c29ed7b6d4bd8da1d1@188.166.126.81:26656\"/' $HOME/.bcna/config/config.toml
    ```
6. You can **set the minimum gas prices** for transactions to be accepted into your node’s mempool. This sets a lower bound on gas prices, preventing spam.
    ``` 
    sed -E -i 's/minimum-gas-prices = \".*\"/minimum-gas-prices = \"0.001ubcna\"/' $HOME/.bcna/config/app.toml
    ```

7. **Open the P2P port (26656 by default)**

    ```sudo ufw allow 26656```

8. Test the connection (**CTRL + C to stop**)

    `bcnad start --log_level info`
    ```
    3:31PM INF Committed state appHash=77D16BED3F109A4A05A971C92602029569E049DFC1DC128CFF5CCAE3158F4B1B height=3886 module=state txs=0
    3:31PM INF Indexed block height=3886 module=txindex
    3:31PM INF minted coins from module account amount=1034628bcna from=mint module=x/bank
    3:31PM INF Executed block height=3887 invalidTxs=0 module=state validTxs=0
    3:31PM INF commit synced commit=436F6D6D697449447B5B38332031333820373720313731203135362032333220313431203435203137332037372031352031363020373120393720393520352031393020313836203733203131342034322031313620313230203536203338203230203337203437203231392032353220343920385D3A4632467D
    3:31PM INF Committed state appHash=538A4DAB9CE88D2DAD4D0FA047615F05BEBA49722A7478382614252FDBFC3108 height=3887 module=state txs=0
    ```

9. **Service creation**
**Ensure that you stopped** the previous test with CTRL+C.
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
    
Enable and activate the BCNAD service.

```
    sudo mv bcnad.service /lib/systemd/system/
    sudo systemctl enable bcnad.service && sudo systemctl start bcnad.service
```
Check the logs to see if it is working:
    ```
    sudo journalctl -u bcnad -f
    ``` 
    
11. **Check the sync:** while `catching_up = true` the node is syncing. Also you can compare your current block with the last synced block of another node or at our [Explorer](https://testnet-explorer.bitcanna.io):
    ```
    curl -s localhost:26657/status  | jq .result.sync_info.catching_up
    #true output is syncing - false is synced

    curl -s localhost:26657/status | jq .result.sync_info.latest_block_height
    #this output is your last block synced

    curl -s "https://rpc-testnet.bitcanna.io/status?"  | jq .result.sync_info.latest_block_height
    #this output the public node last block synced
    ```

## Step 2 - Become a validator
To become a validator you need to perform additional steps. 
Your node must be fully synced in order to send the TX of validator creation and start to validate the network. You can check if your node has fully synced by comparing your logs and the latest block in the explorer (https://wallet-testnet.bitcanna.io/)

1. **Get test-coins from Discord Faucet:**
Go to #testnet-faucet channel and claim your coins (one time) to your address using this syntax: 
    `!claim your_address`
    For example: 
    `!claim bcna14shzreglay98us0hep44hhhuy7dm43snv38plr`

2. **Send the _Create validator_ TX:**

> Read the FAQ manual to understand all parameters:
> https://hackmd.io/_R2KtQzAS02QXdwpJdqmnw

When you have your node synced and your wallet funded with coins (see [_Some useful commands_](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/blob/main/instructions/public-testnet/validator-guides/useful.md)) send the TX to become _validator_ (change _wallet_name_ and _moniker_):
> You can use quotes to include spaces and more than two words
`--from "Royal Queen Seeds"`

```
bcnad tx staking create-validator \
    --amount 1000000ubcna \
    --commission-max-change-rate 0.10 \
    --commission-max-rate 0.2 \
    --commission-rate 0.1 \
    --from WALLET_NAME \
    --min-self-delegation 1 \
    --moniker MONIKER \
    --pubkey $(bcnad tendermint show-validator) \
    --chain-id bitcanna-dev-3 \
    --gas auto \
    --gas-adjustment 1.5 \
    --gas-prices 0.001ubcna
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

# First upgrade: 25 January at 16:15h CET (UTC+1)
1. Setup the halt-height number to stop the network.
    ```
    nano ~/.bcna/config/app.toml
    halt-height = 177735   (save the file)
    sudo service bcnad restart
    ```
2. Download the new binary or compile it.
    ```
    wget https://github.com/BitCannaGlobal/bcna/releases/download/v.1.3.0-rc2/bcnad
    sha256sum bcnad 
        f5beb5613bdf2fc6b07e11b771349f23347536d244713d08eb0b7e10c2b7a6ed  bcnad
    chmod +x bcnad
    ./bcnad version
            .1.3.0-rc2-3-ge4da27a

    ```
3. Wait until the halt-heigh ~ 16:15h CET . The daemon and the blockchain will stop.
    ```
    sudo journalctl -u bcnad -f   #(you can see here the log)
    sudo mv bcnad $(which bcnad)  #(replace the binary)
    bcnad version
        .1.3.0-rc2-3-ge4da27a
    sudo service bcnad start      #(start the service again)
    ```
