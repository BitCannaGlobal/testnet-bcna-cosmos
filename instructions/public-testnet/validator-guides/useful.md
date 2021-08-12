# BitCanna Public Testnet. Useful commands for CLI environment

## State, address and basic transfers
#### Know the state of your validator, and the state of the sync process
You can query using the CLI (command line interface) or request a HTTP query at the RPC server:
    ```
    bcnad status
    ```

#### The current blockheight (asking the RPC):
```
    curl -s "localhost:26657/status" | jq .result.sync_info.latest_block_height
``` 
#### List your address(es)
```
bcnad keys list
  name: yournamewallet
  type: local
  address: bcna1g3fd6rslryv498tjqmmjcnq5dlr0r6udm2rxjk
  pubkey: ...
  mnemonic: ""
  threshold: 0
  pubkeys: []
  ```
#### Check your balance
```
bcnad query bank balances bcna14zvcj7uq5wnwe0hus9rl0yfd3sdckq40tuy6mu --output json | jq
```

#### Send coins from your wallet to an address
Usage:
`bcnacli tx bank send [from_key_or_address] [to_address] [amount] [flags]`
```
bcnad tx bank send bcna14zvcj7uq5wnwe0hus9rl0yfd3sdckq40tuy6mu bcna1r7f6445p8u569ssumjcvsl6tnys65zgah5zgvr 1000000000000ubcna   --chain-id bitcanna-testnet-6 -y --gas auto --gas-adjustment 1.5 --gas-prices 0.01ubcna
```
*Send from bcna14..(self) to bcna16.. 1 million BCNAs (6 decimals without dot)*

> When you don't remember your address you can use subqueries with `$(bcnad keys show Mywallet -a)`:
```
bcnad tx bank send $(bcnad keys show Mywallet -a) bcna1r7f6445p8u569ssumjcvsl6tnys65zgah5zgvr 1000000000000ubcna   --chain-id bitcanna-testnet-6 -y --gas auto --gas-adjustment 1.5 --gas-prices 0.01ubcna
```
> NOTE: We can avoid the `--chain-id bitcanna-testnet-6` flag configuring one time with command: 
```
bcnad config chain-id bitcanna-testnet-6
```

#### Check the HASH of one TX

When you send a TX to the blockchain it returns a unique hash which identifies it and shows a result. You can check this hash anytime with command line:
```
bcnad query tx 496F272F7A978387A085115E8F48C60FD6D1BF9722A6889805C2CED0B64429FC --output json | jq
```
In addition, you can check the hash from the browser.

## Delegator corner
**As a delegator you are not obliged to run a fullnode or validator.**
There are two ways to delegate and transfer funds:
* Using a command line interface (*bcnad*)
* Log in into a [webwallet](https://testnet-wallet.bitcanna.io/) or mobile wallet (to be released)

Some useful commands:

* **Query delegations.**

*If you want to know if your address is delegating you can type:*

`bcnad query staking delegations YOUR_BCNA_ADDRESS  --output json | jq`

In example

```bcnad query staking delegations bcna14zvcj7uq5wnwe0hus9rl0yfd3sdckq40tuy6mu  --output json | jq```

*If you know you had a previous delegation to a validator:*

```
bcnad query staking delegationbcna1dpa0fkmjf3taqnc4fu43l2lgektt3caey6sf5j bcnavaloper1dpa0fkmjf3taqnc4fu43l2lgektt3caea8pfwx --output json | jq
```

* **Make a delegation**: you need your address and validator's address and indicate the amount of coins to delegate:

```
bcnad tx staking delegate  bcnavaloper1dpa0fkmjf3taqnc4fu43l2lgektt3caea8pfwx 100000000ubcna --from walletname --gas-adjustment 1.5 --gas auto --gas-prices 0.01ubcna --chain-id bitcanna-testnet-6
```
* **Redelegate illiquid tokens from one validator to another**: tx staking redelegate [src-validator-addr] [dst-validator-addr] [amount]
We can also move our staked tokens from one validator to another by:
```
bcnad tx staking redelegate bcnavaloper1km88duvfjhq4cda36cf75kgfrv89k8mcldtshg bcnavaloper1pyrnags7c0n74d2ffcnlyyjyr5454m08xprsnz 10000000000ubcna --from walletname --gas-adjustment 1.5 --gas auto --gas-prices 0.01ubcna --chain-id bitcanna-testnet-6
```
* **Unbond shares from a validator**: tx staking unbond [validator-addr] [amount]
Delegator can unbond their staked tokens by:
```
bcnad tx staking unbond bcnavaloper1dpa0fkmjf3taqnc4fu43l2lgektt3caea8pfwx 10000000000ubcna --from walletname --gas-adjustment 1.5 --gas auto --gas-prices 0.01ubcna --chain-id bitcanna-testnet-6
```
_Remark_: Note that funds will only be available after the `unbonding_time` has passed.
* Check the list of validators: 

```bcnad query staking validators  --output json | jq```

## Rewards as delegator

* Get and claim the rewards as delegator:

**Get info about the rewards from all validators where an address have delegated:**

```
bcnad query distribution rewards bcna1km88duvfjhq4cda36cf75kgfrv89k8mcxs6sdu --output json | jq
```

**Claim the reward:**

```
bcnad tx distribution withdraw-all-rewards --from MyWalletName --gas-adjustment 1.5 --gas auto --gas-prices 0.01ubcna --chain-id bitcanna-testnet-6
```
## Rewards as validator
Withdraw rewards from a given delegation address,
and optionally withdraw validator commission if the delegation address given is a validator operator (...bonded to)

Specifying the flag `--commission ` it withdraws the validator's commission in addition to the rewards.

* Query: 
```
bcnad query distribution commission bcnavaloper1r7f6445p8u569ssumjcvsl6tnys65zgawfngkh 
commission:
 -amount: "4152546.128239246736202461"
  denom: ubcna
```

* Claim:
```
bcnad tx distribution withdraw-rewards bcnavaloper1r7f6445p8u569ssumjcvsl6tnys65zgawfngkh --from MyWalletName --commission --gas-adjustment 1.5 --gas auto --gas-prices 0.01ubcna --chain-id bitcanna-testnet-6
```

## Edit a validator info
After the creation of a validator, we can edit it to display more information or fix errors. In example, we can complete/edit the website of a validator, their email contact, Keybase profile, etc.

```
Usage:
  bcnad tx staking edit-validator [flags]
```
In this example we are going to edit/add a Keybase profile and a website + email info:
```
bcnad tx staking edit-validator --identity "C3480D0FD85F60B9" --website "www.bitcanna.io" --security-contact "dev@bitcanna.io" --gas auto --gas-adjustment 1.4 --gas-prices 0.010ubcna --from ValidatorWallet
```

## Working with local Keys
### Listing (all) the wallets

``` 
bcnad keys list --output json |jq
Enter keyring passphrase:
[
  {
    "name": "resources",
    "type": "local",
    "address": "bcna14zvcj7uq5wnwe0hus9rl0yfd3sdckq40tuy6mu",
    "pubkey": "bcnapub1addwnpepqdf620ae6ekqt999sfhzjszc8d8w49c7umad0l4xrv62jhax63mkyazd8z5"
  },
  {
    "name": "validator",
    "type": "local",
    "address": "bcna1dpa0fkmjf3taqnc4fu43l2lgektt3caey6sf5j",
    "pubkey": "bcnapub1addwnpepq05hazfnerkart6v3dc2xg6usn7kkr679axjntkvem6zla49s238yj765ku"
  }
]
```
### List information from a specific wallet
`bcnad keys show <wallet_name>`

`bcnad keys show resources --output json | jq`

### Export the  privkey (wallet) 
`bcnad keys export <name_wallet>`

```bcnad keys export resources
Enter passphrase to decrypt your key:
Enter passphrase to encrypt the exported key:
Enter keyring passphrase:
-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: C1xxxx17C197152CD34XXX2BE68B6Dxxx
type: secp256k1

KNxxxxvyYjV6Yq2WuqbaRlM6OikkAG/8V/xju2Faiqxx+x+xCtYNhJN0QluFNTEJ
jv3M8AS4JgwEkPUWExrxxxxxxKVf+gVGcMxYYUMI=
=zXZO
-----END TENDERMINT PRIVATE KEY-----
```
 **Copy the content in a file named __keyfile_wallet_resources.key__ in example**

### Restore your wallet using the private key
`bcnad keys import <name> <keyfile>`

```
bcnad keys import wallet2 keyfile_wallet_resources.key
```

### Restore your wallet using the seedphrase
```bcnad keys add <name_wallet> --recover ```
