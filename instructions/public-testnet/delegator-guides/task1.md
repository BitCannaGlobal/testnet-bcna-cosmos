# Task 1 Creation of the KEPLR wallet & connecting to BitCanna web wallet

Keplr is a browser extension supported by both the Brave and Google Chrome browsers. This extension provides an easy and secure way of accessing your BitCanna wallet and is widely adopted by the Cosmos community. 

You can also connect to the BitCanna web wallet in alternative ways, such as a ledger device directly, or a ledger to Keplr. You can also continue exploring the wallet without an address, or with someone else's address (only view data). In this guide, we will connect to the web wallet with Keplr.

1.  Download the Keplr browser extension on https://chrome.google.com/webstore/detail/keplr/dmkamcknogkgcdfhhbddcghachkejeap 

2. Once Keplr is downloaded, launch the app and choose `Create new account`. Save the seed phrase of your account, and store it safely. **Losing the seedphrase means you lose access to your wallet (unrecoverable).** It doesn't matter what you fill out in the `Account name` field, it's simply a descriptive name.

3. Click `Next`, and fill out your seed phrase. Keplr asks this to confirm that you correctly stored your seed phrase. Next, click `Register`. Your account has now been registered.

4. Now head over to the [BitCanna web wallet](https://testnet-wallet.bitcanna.io/), and click `Keplr Browser Extension` as the login method. Keplr will now ask to add the BitCanna testnet chain to your Keplr wallet. Click `Approve`. The Keplr wallet will now automatically login to your account on the BitCanna web wallet. If it doesn't, make sure to select the correct network, as shown in the image below.

![](https://i.imgur.com/wknY7rf.png)

![](https://i.imgur.com/L0HETyz.png)

5. Now it's time to claim your testnet coins. Copy your `bcnaxxxxxxxx` address, which you can find in both the Keplr wallet, as in the web wallet. 

![](https://imgur.com/OXddLCt.png)

6. Head over to the BitCanna Discord, and find the `testnet-faucet` channel. Write the following command:

`!claim your_address`

For example:
```
!claim bcna14shzreglay98us0hep44hhhuy7dm43snv38plr
```

All done! Time to head to task 2 and perform transactions :) 
