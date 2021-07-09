  # STAGE 5. Task #2 Create a Microtick wallet and perform IBC transactions between the chains

In this Task we are going to send BCNA test-tokens to a Microtick chain address.

Also we are going to do the way back.

The whole task is going to be possible using a public IBC relayer connected to Microtick Tesnet (microtick-testnet-rc8-1)

This is a brief scheme:

* From BC chain > Send ubcna > MT chain
* From MT chain > Send ubcna > BC chain (return)

Extra points awarded if you do:

* From MT chain > Send utick > BC chain
* From BC chain > Send utick > MT chain

## Extra info
Public BitCanna IBC relayer:
 * src-port: `transfer`
 * src-channel: `channel-0`
 * chain-id: `bitcanna-testnet-5`
 * rpc-server: http://seed2.bitcanna.io:16657

Public Microtick IBC relayer:
 * src-port: `transfer` 
 * src-channel: `channel-3`
* chain-id: `microtick-testnet-rc8-1`
* rpc-server: http://seed2.bitcanna.io:26657

Explorers:
* [Microtick Testnet](https://testnet-explorer.microtick.zone/)
* [BitCanna Testnet](https://testnet-explorer.bitcanna.io/)


## Step 0. Get a new Microtick address and fund it
> We are going to avoid making a Microtick validator for simplicity. We will use a remote public RPC node to send TXes.
> * BitCanna facilitates one at: `tcp://seed2.bitcanna.io:26656`


1. Get the Microtick testnet rc9 binary:
```
wget https://microtick.com/releases/testnet/stargate/mtm-v2-rc9-linux-x86_64.tar.gz
```
2. Sha256sum of this file must be: `0e8eaab9617051bdd8247b0b3f8e27537ca41bb9f7195cb757ad8cf22a7de523`

```
sha256sum mtm-v2-rc9-linux-x86_64.tar.gz
```
3. Decompress the executable and move to system path:
```
tar -xzvf mtm-v2-rc9-linux-x86_64.tar.gz
rm mtm-v2-rc9-linux-x86_64.tar.gz
sudo mv mtm /usr/local/bin/
mtm version
      --> output: mtm v2 (mtm-v2-rc9)

```
4. Get a Microtick address, set your own name
```
mtm keys add MicroWalletName 
```
This should be the output: 
```
- name: MicroWalletName
  type: local
  address: micro1v0hrrga4yulkquyfnd56cfx4003dzhgggyknkg
  pubkey: micropub1addwnpepqfqnehhejmgjnv5nquvmcdp42d65t05jakvqeay74gx02jd6q0rs6lg4ewu
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

digital rely message surprise  despair popular rabbit frog gravity clog culture destroy summer radar high erode chef spend situate time renew iron token venture
```

> Please store these seed words, you might need them in future tasks.



5. Fund your address with TICK testnet coins at our Discord channel testnet-faucet using the same bot you used to claim BCNA coins:
```
!claim micro1v0hrrga4yulkquyfnd56cfx4003dzhgggyknkg
```
6. We will use ENV to remember the addresses and parameters.
```
export MT_ADDR="$(mtm keys show MicroWalletName -a)"
```
After providing the password, write again:
```
echo $MT_ADDR  #it must show your MT address
export MT_RPC="http://seed2.bitcanna.io:26657"
export MT_chain="microtick-testnet-rc8-1"
export MT_channel="channel-3"
export MT_port="transfer"
```
To finish, we can also use a ENV for BitCanna params (change by yours):
```
export BCNA_ADDR="$(bcnad keys show BCwalletname -a)"
```
After providing the password, write again:
```
echo $BCNA_ADDR  #it must show your BC address
export BCNA_RPC="http://seed2.bitcanna.io:16657"
export BCNA_chain="bitcanna-testnet-5"
export BCNA_channel="channel-0"
export BCNA_port="transfer"
```


## Step 1. Send `ubcna` tokens to a Microtick address

**Important issue:**
It is possible that your initial IBC transactions will time-out, even though it will look like your transaction was successful (you will get a TX hash). We're currently investigating this issue. You can check if your coins have actually moved your balance by performing the query command or in the explorer. If your coins haven't left, simply submit your transaction AGAIN, because the bug only occurs in the initial transaction. Any information or suggestions regarding this issue is very welcome.

1. Before start you can check your balance at your BCNA wallet:
```
bcnad query bank balances $BCNA_ADDR 

balances:
- amount: "999996810"
  denom: ubcna
```

2. You can check the balance at your MT wallet as well (optional)
```
mtm query bank balances $MT_ADDR --node $MT_RPC
balances:
- amount: "1000000000"
  denom: utick
```

3. Send the tokens from BC to MT


Syntax:
```
Usage:
  bcnad tx ibc-transfer transfer [src-port] [src-channel] [receiver] [amount] [flags]
```
So, you must indicate your own Bitcanna source wallet:

```
bcnad tx ibc-transfer  transfer $BCNA_port $BCNA_channel $MT_ADDR 1000000ubcna --chain-id $BCNA_chain --from BCwalletname -y --gas auto --gas-prices 0.09ubcna --gas-adjustment 1.5  --packet-timeout-timestamp 6000000000000
```
> **You must provide the TX HASH as proof in the Task Center, please do so at the end of the task because you need to provide them all together in 1 task**

4. Check the balance again:
```
bcnad query bank balances $BCNA_ADDR 

balances:
- amount: "998886810"
  denom: ubcna

------------------------------------------------------------

mtm query bank balances $MT_ADDR --node $MT_RPC 

balances:
- amount: "1000000"
  denom: ibc/754B5F1A44FA09CD03CE503C75A4FB2CC37FE3C00D7DA044CF4832F87ADC2D12
- amount: "1000000000"
  denom: utick
```
> **YES! The `ibc/754B5F1A44FA09CD03CE503C75A4FB2CC37FE3C00D7DA044CF4832F87ADC2D12` is the BCNA token in the IBC Cosmos Ecosystem**

You can also check the balance in the explorer (use your address):
https://testnet-explorer.microtick.zone/account/micro1v0hrrga4yulkquyfnd56cfx4003dzhgggyknkg

## Step 2. Send back `ubcna` tokens from Microtick chain to BC chain.
Regarding the previous balances in Step 1.4, let's make the reserve TX, by using the `mtm` binary we will send the IBC token to our BC chain again.

**Important issue:**
It is possible that your initial IBC transactions will time-out, even though it will look like your transaction was successful (you will get a TX hash). We're currently investigating this issue. You can check if your coins have actually moved your balance by performing the query command or in the explorer. If your coins haven't left, simply submit your transaction AGAIN, because the bug only occurs in the initial transaction. Any information or suggestions regarding this issue is very welcome.

1. Send the IBC coin to BitCanna chain again: 
Syntax:
```
Usage:
  mtm tx ibc-transfer transfer [src-port] [src-channel] [receiver] [amount] [flags]
```
So, you must indicate your own Microtick source wallet:

```
mtm tx ibc-transfer  transfer $MT_port $MT_channel $BCNA_ADDR 1000000ibc/754B5F1A44FA09CD03CE503C75A4FB2CC37FE3C00D7DA044CF4832F87ADC2D12 --chain-id $MT_chain --from MicroWalletName -y --gas auto --gas-prices 0.09utick --gas-adjustment 1.5  --node $MT_RPC  --packet-timeout-timestamp 6000000000000
```
> **You must provide the TX HASH as proof in the Task Center**

2. Check balances again: The IBC coins should be 0.
```
mtm query bank balances $MT_ADDR --node $MT_RPC 
balances:
- amount: "0"
  denom: ibc/754B5F1A44FA09CD03CE503C75A4FB2CC37FE3C00D7DA044CF4832F87ADC2D12
- amount: "999995585"
  denom: utick
``` 
And in the BitCanna chain your amount should equal 1000000ubcna again.

```
bcnad query bank balances $BCNA_ADDR

balances:
- amount: "999996810"
  denom: ubcna
```



## Points awarded (100) if you perform the following TXes correctly without a guide:

From MT chain > Send utick > BC chain
From BC chain > Send utick > MT chain

No step-by-step instructoins will be provided.

You should be able to do this by doing some research, or working together with other participants. Good luck!
> Don't forget to note the TX HASHes in the Task Center.


From MT chain > Send utick > BC chain:
```

```

From BC chain > Send utick > MT chain in return:
```

```
