# STAGE 4. Task 4. The SATIVA Upgrade
Core Team will create a governance proposal on 1st July, 2021 at ~14:30 UTC, to upgrade the network to a new release, `v0.testnet7`.

__New to upgrades?__ Read [this](https://docs.cosmos.network/master/modules/gov)

**Upgrade Schedule:**
- Proposal: 1st Jul, 2021 ~14:30 UTC
- Voting Period: 1st Jul-03rd Jul, ~14:30 UTC
- Upgrade Height: `59576` (Close to 05th Jul, ~1400 UTC)

**What should validators do?**
- Review the software upgrade proposal and cast your vote before voting period endtime (you should have already voted in task #3, otherwise go cast your vote)
- Upgrade your nodes by performing Option 2

---

# The SATIVA Upgrade


## How to upgrade


There are two ways to perform the UPGRADE when the time (`block-height=59576`) comes:
* Option 1 Manually, after the network upgrade (0 points)
* Option 2 Automated, with Cosmovisor, ahead of the network upgrade **(RECOMMENDED)** (200 points)

Choose one of them (not both)

### Option 1 (0 points awarded): Manually, if you fail to update before the network upgrade (normally recommended for Wallet nodes, Explorer nodes, etc)

1) Stop your bcnad service **after you see this** in your logs ` ERR UPGRADE "sativa" NEEDED at height: 59576`

```
sudo service bcnad stop
```
2) Download the binary or compile it from the source
```
wget -nc https://github.com/BitCannaGlobal/testnet-bcna-cosmos/releases/download/v0.testnet7/bcnad
```
3) Check the sha256sum. It must return: `d36d6df2a8155a92f4c6a9696ac38e4878ec750d5dff9ad8a5c5e3fadbea6edb`

To check it:
```
sha256sum ./bcnad
```

4) Verify that the version is `v0.testnet7`
```
chmod +x bcnad
./bcnad version
```
5) Move file to exec path and overwrite the previous version
```
sudo mv bcnad $(which bcnad)   #copy&paste don't replace anything
```
Alternatively you can run (if you know the destination): 
```
sudo mv bcnad /usr/local/bin/
```
6) Start again the service bcnad
```
sudo service bcnad start
```
7) Ensure with the logs that everything is going ok
```
sudo journalctl -u bcnad -f
```

---
## Option 2: Setting up Cosmovisor (recommended) (200 points awarded)
With Cosmovisor you can schedule updates ahead of time by setting the location of the old binary and the new binary.

Cosmovisor is a **service that listens to the LOGs output and waits** for the correct block to update.

 - ### Step 1. Creating directories for cosmovisor
 ```
mkdir -p ${HOME}/.bcna/cosmovisor/genesis/bin
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/sativa/bin
 ```
 
 - ### Step 2. Get Cosmovisor 
 You can build from the source or download from our github
 ```
 wget https://github.com/BitCannaGlobal/testnet-bcna-cosmos/releases/download/v0.testnet7/cosmovisor
 chmod +x cosmovisor
 sudo mv cosmovisor /usr/local/bin
 ```
 
 - ### Step 3. Get the new binary `v0.testnet-7` from BitCanna or build it
 
1) If you want to build the binary, instructions to build is on the [README](https://github.com/BitCannaGlobal/testnet-bcna-cosmos#readme) of Github


2) If you want download the binary:
```
wget -nc https://github.com/BitCannaGlobal/testnet-bcna-cosmos/releases/download/v0.testnet7/bcnad
```
3) Check the sha256sum. It must return: `d36d6df2a8155a92f4c6a9696ac38e4878ec750d5dff9ad8a5c5e3fadbea6edb`

    To check it:
    ```
    sha256sum ./bcnad
    ```

4) Verify that version is `v0.testnet7`
```
chmod +x bcnad
./bcnad version
```
 
5) Move the newly built binary to it's directory
 ```
mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/sativa/bin/bcnad
 ```
 
- ### Step 4. Finishing Cosmovisor setup. 

1) Copy the current (.testnet6) binary to Cosmovision genesis (if is the first update) folder:
 ```
#copy and paste literally, do not change anything.
cp $(which bcnad) ${HOME}/.bcna/cosmovisor/genesis/bin/
 ```
 
2) Setup cosmovisor current version link
```
ln -s -T ${HOME}/.bcna/cosmovisor/genesis ${HOME}/.bcna/cosmovisor/current
```
3) You can check that everything is OK: 
```
ls .bcna/cosmovisor/ -lh

total 8.0K
lrwxrwxrwx 1 raul raul   35 Apr 27 14:16 current -> /home/raul/.bcna/cosmovisor/genesis
drwxrwxr-x 3 raul raul 4.0K Apr 27 14:09 genesis
drwxrwxr-x 4 raul raul 4.0K Apr 27 14:15 upgrades
```
4) Setup cosmovisor systemd service (copy and paste all to create the file service)
```
echo "[Unit]
Description=Cosmovisor BitCanna Service
After=network-online.target
[Service]
User=${USER}
Environment=DAEMON_NAME=bcnad
Environment=DAEMON_RESTART_AFTER_UPGRADE=true
Environment=DAEMON_HOME=${HOME}/.bcna
ExecStart=$(which cosmovisor) start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```
5) Let's change the BCNAD service by COSMOVISOR service
```
sudo mv cosmovisor.service /lib/systemd/system/
sudo systemctl daemon-reload
sudo systemctl stop bcnad.service && sudo systemctl disable bcnad.service 
sudo systemctl enable cosmovisor.service && sudo systemctl start cosmovisor.service
```

Check logs

```
sudo journalctl -u cosmovisor -f
``` 

### Step 5. Finishing the work
If everything is OK, you must delete the old binary from path and let the control of the binaries/service to Cosmovisor. In place of bcnad you can use `cosmosvisor` , in example `cosmovisor status`


To do it make the next changes:

1. Add this variables to then end of .profile file

```
nano $HOME/.profile

#add to the end of the file:
export DAEMON_NAME=bcnad
export DAEMON_RESTART_AFTER_UPGRADE=true
export DAEMON_HOME=${HOME}/.bcna
PATH="${HOME}/.bcna/cosmovisor/current/bin:$PATH" #we do this to continue use bcnad as binary for send commands, totally optional.
```

2. Reload the config of .profile with:
```
source .profile
```

3. Try cosmovisor
```
cosmovisor version #show the current  BCNAD version v0.testnet6 # Will be v0.testnet7 after the upgrade block- height
bcnad version #show the same version
cosmovisor status  #must show the sync info
```
4. If previous test was right, delete old binary
```
sudo rm /usr/local/bin/bcnad  #or wherever you had it
```
### Reminder

Done! In the future, you must use the cosmovisor command instead of bcnad command if you want to perform service related commands. In example: 

* for stop the service:
```
sudo service cosmovisor stop 
```
* start the service or reload:
```
sudo service cosmovisor start

sudo service cosmovisor restart
```
* to check the logs

```
sudo journalctl -u cosmovisor -f
```

