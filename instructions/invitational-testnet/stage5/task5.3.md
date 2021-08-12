# STAGE 5. Task 3. Setup your own IBC relayer and perform transactions

The following instructions are for transferring the tokens from `bitcanna-testnet-5` (source) chain to `microtick-testnet-rc8-1` (destination) chain and transfer them back to source chain.

In real life you can set up connections to Osmosis, Gravity, etc, etc. 

## Setting up the relayer

Useful links:
* [Node.js 14.16.1](https://www.google.com/search?q=install+nodejs+ubuntu+20.04) or later is required.
* [TypeScript Relayer](https://github.com/confio/ts-relayer)

### Install Relayer:

```
npm i -g @confio/relayer@main    #sometimes it need exec as root with sudo
```
This will generate two files: `ibc-setup` and `ibc-relayer`

### Usage
#### After installation, ibc-setup and ibc-relayer executables are available.

`ibc-setup`
Collection of commands to quickly setup a relayer and query IBC/chain data.

* run ibc-setup --help to print usage
* [ibc-setup spec](https://github.com/confio/ts-relayer/blob/main/spec/ibc-setup.md)

`ibc-relayer`
Reads the configuration and starts relaying packets.

* run ibc-relayer --help to print usage
* [ibc-relayer spec](https://github.com/confio/ts-relayer/blob/main/spec/ibc-relayer.md)

### Setup the relayer.
- **Step-1:** Init relayer config

```
ibc-setup init 
```
    
It will create the folders and config files, don't worry by the error messages:

```
Initialized home directory at /root/.ibc-setup
Pulled default registry.yaml from remote.
Exited early. Registry file downloaded to /root/.ibc-setup/registry.yaml. Please edit that file and add any chains you wish. Then complete the initialization by running ibc-setup init --src <chain-1> --dest <chain-2>.
```

- **Step-2:** We will write the Microtick and Bitcanna chain configs at the relayer config file by copy/pasting the text below

    Note: Please change gas-prices value based on your rpc-node. Also make sure to update the `bitcanna-testnet-5` RPC node if you are running this on a different machine.
    Copy&paste everything:
```
echo "  bitcanna:
    chain_id: bitcanna-testnet-5
    prefix: bcna
    gas_price: 0.01ubcna
    ics20_port: 'transfer'
    rpc:
      - http://localhost:26657
  microtick:
    chain_id: microtick-testnet-rc8-1
    prefix: micro
    gas_price: 0.01utick
    ics20_port: 'transfer'
    rpc:
      - http://seed2.bitcanna.io:26657" >> $HOME/.ibc-setup/registry.yaml
```
- **Step-3:** Generate the `app.yaml` config file with the account that relayer will use.

    ```bash
    ibc-setup init --src bitcanna --dest microtick

    >> will output something like:
    Saved configuration to /root/.ibc-setup/app.yaml
    Source address: bcna1qtenuvr6lm40lc4EXAMPLEvsa4nqdstdye3mgd
    Destination address: micro1qtenuvr6lm40lcEXAMPLEvsa4nqdstdv3uaet
    ```

You must fund the address (the output from your cli) using the Discord faucet:

```bash
!claim bcna1qtenuvr6lm40lc4EXAMPLEvsa4nqdstdye3mgd
!claim micro1qtenuvr6lm40lcEXAMPLEvsa4nqdstdv3uaet
```

**IMPORTANT: If addresses haven't got funds it won't work!**
You can check both address in the respective explorer.


- **Step-4** Initiate the connection:
    ```
    ibc-setup ics20
    ```
    You should see a similar output to this (can take around 2 minutes):
    ```
     info: Connection open init: 07-tendermint-2 => 07-tendermint-9
    info: Connection open try: 07-tendermint-9 => 07-tendermint-2 (connection-2)
    info: Connection open ack: connection-2 => connection-3
    info: Connection open confirm: connection-3
    Created connections [bitcanna-testnet-5, connection-2, 07-tendermint-2] <=> [microtick-testnet-rc8-1, connection-3, 07-tendermint-9]
    info: Create channel with sender bitcanna-testnet-5: transfer => transfer
    Created channel:
      bitcanna-testnet-5: transfer/channel-1 (connection-2)
      microtick-testnet-rc8-1: transfer/channel-4 (connection-3)

    ```

If you don't see errors, you can go ahead. Now read your log and find your custom channels as shown below. In the below example, the channels are 1 and 4.
```
    Created channel:
      bitcanna-testnet-5: transfer/channel-1 (connection-2)
      microtick-testnet-rc8-1: transfer/channel-4 (connection-3)
```
From the previous output you have to set your custom channels in the command below (replace the channel numbers):
```
export BCNA_port="transfer"
export BCNA_channel="channel-1"

export MT_port="transfer"
export MT_channel="channel-4"
export MT_RPC="http://45.79.187.79:26657"
```

> Save your channel numbers and provide them as proof in the task center. (Please only submit one task, so submit them later once you've completed the entire guide)
> Note: If you log out you will need to set this ENV again

- **Step-5** Start the relayer :)
For production you would need to use a "service". We will not go through this in this guide. So if u want to keep your relayer open, the relayer needs to keep runnning in your screen. So, please use screen sessions (experienced linux users) or open a new terminal (easy for beginners) to perform your transactions in Step 6.

The command below will allow start your own relayer in a [SCREEN](https://www.google.com/search?q=screen+linux) session or in a new terminal window, instead of running a service.
```
ibc-relayer start  --poll 10 -v --src bitcanna --dest microtick 

```
NOTE: We repeat, you must maintain the relayer running when you are performing the IBC transactions.


- **Step 6** Make some TX between chains using your relayer.
Now, in order to earn the task points, you will have to use your own relayer to transfer coins between the chains, the only difference related to Task 5.2, is that now you use another port/channel in the source chain and the destination chain. You could set your ENV values again (which you set in Task 5.2) or use direct values.

Task 5.2 variables can be found [here.](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/blob/main/instructions/stage5/task5.2.md)

You can use your old addresses, the ones you used in task 5.2 for example. The new addresses generated in this task are only for the relayer to properly work.

* Send `ubcna` tokens from BC to MT chain
* Send `utick` tokens from MT to BC chain

Please add a memo to your transactions, example: 
```
--memo "Send ubcna tokens from BC to MT chain"
```

You must submit a proof of each IBC denom hash as described in the task instructions on Gitbook, and also submit BitCanna and Microtick channel numbers in the Task Center.

