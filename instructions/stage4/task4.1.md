# STAGE 4. Task 1. - Create & submit your GenTX for testnet-4.
> Time interval: From June 30th Wednesday 10.00h UTC to June 29th Wednesday 14.00h UTC

In Stage 4 (bitcanna-testnet-4) We are going to simulate a coordinated start again, which we'll also use in the BitCanna Mainnet, also known as a "Dry run"

We will set a specific date and time to start and we must achieve 2/3 of consensus before this specified date and time to launch the blockchain.

The addresses of all participating validators will be engraved in the Genesis file. These validators will have an initial balance and will be pre-activated.

After a successful launch, other validators can still be added, and get coins after sending the GenTX. 

## Targets to achieve

* Coordinate and test the Mainnet launch (Second dry-run).
* **Achieve 2/3** of the consensus to be able to start the network.
* Start the chain on **July 1st Thursday 14.00h UTC**

## Steps to coordinate the launch.


### Step 1. Prepare the new environment.

1. **Stop your daemon** if it is running:
```
sudo systemctl stop bcnad
```
2. **Keep a backup** of your old validator directory.
cp -a ~/.bcna ~/bcna_backup

3. **Delete the old** GenTX file (if it exists)
```
rm -rf $HOME/.bcna/config/gentx/ 
```
4. **Init the new chain locally.** Change ***Moniker*** by your validator name. The parameter `rewrite` generate a new genesis file locally.
```
bcnad init "Moniker" --chain-id bitcanna-testnet-4 --overwrite
```

5. Prepare your **GenTX.**
For this step you need to know your address/wallet's name. The following command will list all your addresses and wallet's name. If you have multiple wallets, please locate the wallet which you funded in the current testnet.
    ```
    bcnad keys list 
    bcnad keys show YourWalletName -a
         outputs > bcnaInexamplec7j0r5c6x68aa9utspejsemr3356lmn
    ```
You will now add your address to the genesis file. Only change the name of the wallet by your own wallet name, because a change in the coin amount will invalidate the GenTX. Failing to submit the correct GentTX (by changing the amount of bcna for example) will get you disqualified(!).
```
#copy and paste the following complete line, and only change your YourWalletName, failing to do so will get you disqualified
bcnad add-genesis-account $(bcnad keys show YourWalletName -a) 100000000000ubcna
```

**NOTE:** It will ask you by your wallet-passphrase. If no response, no error ;)
### Step 2. Create & submit your GenTX

> Time interval: From June 29th Tuesday 14.00h UTC to 30th Wednesday 14.00h UTC
(You MUST submit your GenTX only between these dates)

1. Generate the TX to become Genesis validator: **change the values between brackets < > ``""``** by yours
```
bcnad gentx <YourWalletName> 90000000000ubcna --moniker "YourValidatorName" --chain-id bitcanna-testnet-4 -y
```
In example: 
`bcnad gentx Myvalidator 90000000000ubcna --moniker "BitCanna Power" --chain-id bitcanna-testnet-4 -y`

**NOTE:** If you are a Cosmos expert you can add others flags to the GenTX. Please note that no points will be awarded to you if you submit an invalid GenTX.

If you did everything right, you will see the following message:

```
Genesis transaction written to "/home/linux/.bcna/config/gentx/gentx-7c030b0172dcc19fad261dac5601c1282282e747.json"
```

2. You must submit this file to your Testnet Dashboard (task 4.1).
**The name of the file uploaded to the task center should be the same as the file generated from Stage 2 Task 1. So don't change the name! (example: gentx-7c030b0172dcc19fad261dac5601c1282282e747.json)**

**IMPORTANT NOTE:** We repeat, users that submit files with the wrong file name (don't worry if knack dashboard modifies the filename when you upload it), a changed amounts in coins, or malformed json structure will be disqualified from this Stage.

**After completing the task, you will notice that you will stop signing blocks and this is completely normal. Please leave your node disabled. It will be jailed and it doesn't matter because `bitcanna-testnet-3` is going to disappear. Simply wait for task 2!** 
