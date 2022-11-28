# Devnet-5: Setup up your validator and join *bitcanna-dev-5*

> IMPORTANT NOTE: If you participated in previous BitCanna testnets,  delete the bitcanna folders and start a fresh new installation:
> 
> ```sudo service bcnad stop``` 
> 
> `rm -rf ~/.bcna && sudo rm $(which bcnad)`

## Target of this DevNet (outdated...)
We created a new testnet that follows the exact upgrade cycle as the current Mainnet. In this devnet the participants can join the chain from version v.1.4.2 codename `strangeBuddheads`, using our statesync script. We will prepare to upgrade to the newest version of the software, v1.5.3-rc1 codename `TrichomeMonster-ICA`, which includes the ***InterChain Accounts module***.

### We are going to work in three new testnets: 
* bitcanna-dev-5 (with `bcnad` current version ~~v.1.4.2, which we will update later~~ v1.5.3-rc1 [upgrade here](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/tree/main/instructions/bitcanna-dev-5#upgrade-instructions-to-v151) )
* innuendo-1 (with `quicksilverd` current version binary) YOU DON'T NEED TO RUN THIS CHAIN.
* axelar-lisbon (with `axelard` current version binary) YOU DON'T NEED TO RUN THIS CHAIN.

## Running a validator on **bitcanna-dev-5**
* Before you start, you want to ensure your system is updated.  Besides other utilities you can install `jq` which is a utility to read and navigate JSON files and output.
    ```
    sudo apt update && sudo apt upgrade -y
    sudo apt install -y build-essential curl wget jq
    ```
* Increase the default open files limit. If we don't raise this value nodes will crash once the network grows large enough.
    ```
    sudo su -c "echo 'fs.file-max = 65536' >> /etc/sysctl.conf"
    sudo sysctl -p
    ```
## Step 1 - Download and run statesync script
By running the statesync script we download the latest binary (v1.5.3-rc1) and sync the chain to the latest block.
1. **Download the statesync script** for new peers from Github:
    ```
    cd ~
    wget https://raw.githubusercontent.com/BitCannaGlobal/cosmos-statesync_client/main/statesync_DEVNET-5_client_linux_new.sh
    ```
2. **Run the script**: 
    ```
    bash statesync_DEVNET-5_client_linux_new.sh
    ```
    Watch the output of the logs. When the chain is synced (wait 1 minute), press ctrl+C to stop the script and proceed with the next step.
    
    This is an example of the output you see when you can stop the script by pressing ctrl+C
    ```
    4:39PM INF indexed block height=2920 module=txindex
    4:39PM INF Timed out dur=4988.140195 height=2921 module=consensus round=0 step=1
    4:39PM INF commit is for a block we do not know about; set ProposalBlock=nil commit=3E75B8B4371324172A860BBBB4BE8B5C2A2C96A7FA5F5507BB8457D0B40F00D2 commit_round=0 height=2921 module=consensus proposal={}
    4:39PM INF received complete proposal block hash=3E75B8B4371324172A860BBBB4BE8B5C2A2C96A7FA5F5507BB8457D0B40F00D2 height=2921 module=consensus
    4:39PM INF finalizing commit of block hash={} height=2921 module=consensus num_txs=0 root=E8705846BEAAA45BC87474D9ACBFBA074447ED8A680FAB5AD53516E7E0B2C7C7
    4:39PM INF minted coins from module account amount=4836690ubcna from=mint module=x/bank
    4:39PM INF executed block height=2921 module=state num_invalid_txs=0 num_valid_txs=0
    4:39PM INF commit synced commit=436F6D6D697449447B5B36352037332032303120393220313733203133203537203930203138352036342035382031323520323230203133392031313620313730203336203932203535203131382031303920363520323037203138382037312031333520313236203234352031343820353520313837203235305D3A4236397D
    4:39PM INF committed state app_hash=4149C95CAD0D395AB9403A7DDC8B74AA245C37766D41CFBC47877EF59437BBFA height=2921 module=state num_txs=0
    4:39PM INF indexed block height=2921 module=txindex
    ```

3. **Move the new `bcnad` binary** to your machine's PATH.
    ```
    sudo mv bcnad /usr/local/bin/ 
    ```
## Step 2 - Prepare the node
To create a validator you need a funded wallet. Once the wallet is created, go to the **#devnet-faucet** channel on [Discord](https://discord.com/channels/805725188355260436/847019574662922260) and claim your devnet coins. For example: `!claim bcna14shzreglay98us0hep44hhhuy7dm43snv38plr`

1. **Create a wallet:**
You can create a wallet with one or more keys (addresses) using `bcnad`.  Replace **"MyFirstAddress"** with your desired name.
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

2. **Service creation**
With all configurations ready you can set up `systemd` to run the node daemon with auto-restart.
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

	Check the logs to see if everything is working correct:
   ```
    sudo journalctl -fu bcnad
   ```

## Step 3 - Create the validator
When your node is synced and your wallet funded it's time to send the TX to become validator:
(change **_WALLET_NAME_** and **_MONIKER_**)
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
	--chain-id bitcanna-dev-5 \
	--gas auto \
	--gas-adjustment 1.5 \
	--gas-prices 0.001ubcna
```

You can check the list of validators (also in [Explorer](https://testnets-cosmos.mintthemoon.xyz/bitcanna/staking)):

   ```
   bcnad query staking validators --output json| jq
   ```

Another **IMPORTANT** but **optional** action is backup your Validator_priv_key:

   ```
    tar -czvf validator_key.tar.gz .bcna/config/*_key.json 
    gpg -o validator_key.tar.gz.gpg -ca validator_key.tar.gz
    rm validator_key.tar.gz
   ```
   This will create a GPG encrypted file with both key files.
##

## UPGRADE INSTRUCTIONS FOR `wakeandbake46.6 v2.0.1-rc2` ON NOVEMBER 28TH

## Governance proposal: halt-height `1.032.049`

https://testnet.ping.pub/bitcanna/gov/10

## Attended (manual) upgrade.

This section of the guide shows how to perform a **manual** update of the binaries after a governance proposal has been approved for a chain update.
1) Stop your bcnad service **after you see this** in your logs `ERR UPGRADE "wakeandbake46.6" NEEDED at height: 1.032.049`
```
sudo service bcnad stop
```
2) Download the binary [or compile it from the source](#If-you-want-to-build-from-the-source)
```
cd ~
rm -f ./bcnad && rm -f ./bcna_linux_amd64.tar.gz # clean the previous downloads
wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v2.0.1-rc2/bcna_linux_amd64.tar.gz
```
3) Check the sha256sum. 
```
sha256sum bcna_linux_amd64.tar.gz
```
> It must return: `a424ad37b301578370bce58134c78891798c9c0b519ec09954054184a2a868e1`

4) Verify that the version is `v2.0.1-rc2`
```
tar zxvf  bcna_linux_amd64.tar.gz
rm bcna_linux_amd64.tar.gz
./bcnad version
```
5) Move the new binary to your machine's PATH and overwrite the previous version
```
sudo mv bcnad $(which bcnad)   #copy&paste don't replace anything
```
> If you know the exact destination you could also run: 
```
sudo mv bcnad /usr/local/bin/ #or wherever you have it
```
6) Start the bcnad service
```
sudo service bcnad start
```

Of course, if you are familiar with the Cosmos environment, you can keep the daemon running while you are compiling/downloading and later make the upgrade in one command: 
```
sudo service bcnad stop && sudo mv bcnad $(which bcnad) && sudo service bcnad start
```
7) Ensure that everything is OK by checking the logs 
```
sudo journalctl -fu bcnad -o cat
```

## Unattended (Cosmovisor) upgrade. 
This section of the guide shows how to perform a **automated** upgrade of the binaries after a governance proposal has been approved for a chain update.

For detailed instructions about setting up Cosmovisor from scratch, check this [guide](https://hackmd.io/jsJCqEyJSHKVOFKjScn3rw).

This guide shows how to download the binary. If you want to build the binary from the source, detailed instructions can be found in the [README](https://github.com/BitCannaGlobal/bcna/blob/main/README.md) of our GitHub (`git checkout v2.0.1-rc2`)

### Step 1. Setup Cosmovisor folder
1) Create new directory
```
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/wakeandbake46.6/bin
```

2) Download the current version `2.0.1-rc2`.
```
cd ~
rm -f bcna_linux_amd64.tar.gz
wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v2.0.1-rc2/bcna_linux_amd64.tar.gz
```
3) Check the sha256sum.
```
sha256sum ./bcna_linux_amd64.tar.gz
```
> It must return: `a424ad37b301578370bce58134c78891798c9c0b519ec09954054184a2a868e1`

4) Verify that the version is:`2.0.1-rc2`
```
rm -rf ./bcnad
tar zxvf  bcna_linux_amd64.tar.gz
rm bcna_linux_amd64.tar.gz
./bcnad version
```
5) Move the newly downloaded binary to the upgrades directory.
```
mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/wakeandbake46.6/bin/
```
> If you build the binary from the code source, move it to the same folder

6) If you want to know if Cosmovisor handles the correct binary file, exec:
```
sudo service cosmovisor status
```
And check the path of the binary file.

## If you want to build from the source

 If you want to build the binary from the source, detailed instructions can be found in the [README](https://github.com/BitCannaGlobal/bcna/blob/main/README.md) of our GitHub:
```
    git clone https://github.com/BitCannaGlobal/bcna.git
    cd bcna
    git checkout v2.0.1-rc2
    make build && make install 
```

 [_Some useful commands_](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/blob/main/instructions/public-testnet/validator-guides/useful.md) to help you navigate the `bcnad` CLI
