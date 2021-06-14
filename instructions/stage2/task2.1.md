# STAGE 2. Task 1. - Create & submit your GenTX.

In stage 2 (bitcanna-testnet-3) We are going to simulate a coordinated start, which we'll also use in the BitCanna Mainnet, also known as a "Dry run"

The main difference with the previous `bitcanna-testnet-2` is that the blockchain will not be live yet. We will set a specific date and time to start and we must achieve 2/3 of consensus before this specified date and time to launch the blockchain.

The addresses of all participating validators will be engraved in the Genesis file. These validators will have an initial balance and will be pre-activated.

After a successful launch, other validators can still be added, and get coins after sending the GenTX. 

## Targets to achieve

* Coordinate and test the Mainnet launch (dry-run).
* **Achieve 2/3** of the consensus to be able to start the network.
* Start the chain on **17th Thursday 14.00h UTC**

## Steps to coordinate the launch.


### Step 1. Prepare the new environment.

1. **Stop your daemon** if it is running:
```
sudo systemctl stop bcnad
```
2. **Delete the old** genesis file and previous GenTX (if it exists)
```
rm -rf $HOME/.bcna/config/genesis.json
rm -rf $HOME/.bcna/config/gentx/ 
```
3. Get the new **Genesis (temporary)**:

```
curl -s https://raw.githubusercontent.com/BitCannaGlobal/testnet-bcna-cosmos/main/instructions/stage2/genesis.json > ~/.bcna/config/genesis.json
```

4. Prepare your **GenTX.**
For this step you need to know your address. The following command will list all your addresses and wallet's name. If you have multiple wallets, please locate the wallet which you funded in the current testnet.
    ```
    bcnad keys list 
    bcnad keys show YourWalletName -a
         outputs > bcnaInexamplec7j0r5c6x68aa9utspejsemr3356lmn
    ```
We will now add your address to the genesis file. Only change the name of the wallet by your own wallet name, because a change in the coin amount will invalidate the GenTX. Failing to submit the correct GentTX (by changing the amount of bcna for example) will get you disqualified(!).
```
#copy and paste the following complete line, and only change your YourWalletName, failing to do so will get you disqualified
bcnad add-genesis-account $(bcnad keys show YourWalletName -a) 100000000000ubcna
```

**NOTE:** It will ask you by your wallet-passphrase. If no response, no error ;)
### Step 2. Create & submit your GenTX

> Time interval: 12th Saturday 10.00h UTC to 15th Tuesday 10.00h UTC 
(You MUST submit your GenTX only between these dates)

1. Generate the TX to become Genesis validator: **change the values between brackets < > ""** by yours
```
bcnad gentx <YourWallet> 90000000000ubcna --moniker "YourValidatorName" --chain-id bitcanna-testnet-3 -y
```
In example: 
`bcnad gentx Myvalidator 90000000000ubcna --moniker "BitCanna Power" --chain-id bitcanna-testnet-3 -y`

If you did everything right, you will see the following message:

```
Genesis transaction written to "/home/linux/.bcna/config/gentx/gentx-7c030b0172dcc19fad261dac5601c1282282e747.json"
```

2. You must submit this file to your Testnet Dashboard (task 2.1).
**The name of the file matters, it must be the same, so there's no need to change it! (example: gentx-yournodeID.json) **

**IMPORTANT NOTE:** We repeat, users that submit files with the wrong file name, a changed amounts in coins, or malformed json structure will be disqualified from this Stage.

**After completing the task, you will notice that you will stop signing blocks and this is completely normal. Please leave your node disabled. It will be jailed and it doesn't matter. Simply wait for task 2!**
