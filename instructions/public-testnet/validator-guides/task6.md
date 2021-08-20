# Task 6 Setting up Cosmovisor to automate updates.
We created a governance proposal on the 19th of August, 2021 at ~12:00 UTC, to upgrade the network to a new release, `v0.3-beta`.

__New to Cosmos-SDK blockchain upgrades?__ We suggest reading [this article.](https://docs.cosmos.network/master/modules/gov)

__New to Cosmovisor?__ We suggest reading [this article.](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor)

**Upgrade Schedule:**
- Proposal: 19th August, 2021 12:00 UTC
- Voting period: 19th August - 22nd August, 12:00 UTC
- Upgrade height: `194976` (Close to 23th August, ~11:00 UTC)

**What should validators do?**
- Review the software upgrade proposal and cast your vote before voting period endtime (you should have already voted in task #5, otherwise go cast your vote)
- Upgrade your nodes by performing Option 2



**What should previous invitational testnet participants do?**
> If you have participated in the last Testnet and you are running currently the `cosmovisor` service please go to end of the guide
- Review the software upgrade proposal and cast your vote before voting period endtime (you should have already voted in task #5, otherwise go cast your vote)

---

# The INDICA Upgrade


## How to upgrade


There are two ways to perform the UPGRADE when the time (`block-height=194976`) comes:
* Option 1 Manually, after the network upgrade.
* Option 2 Automated, with Cosmovisor, ahead of the network upgrade **(RECOMMENDED)**

> **Choose one of them (not both)**

### Option 1: Manually, not recommended. This option is for service providers (not relevant for public testnet) So please head to option 2.

1) Stop your bcnad service **after you see this** in your logs ` ERR UPGRADE "indica" NEEDED at height: 194976`

    ```
    sudo service bcnad stop
    ```
2) Download the binary or compile it from the source.
    ```
    wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v0.3-beta/bcnad
    ```
3) Check the sha256sum. It must return: `1b2f3763dd3f11fe63b0212dc21d20fa09a114760928e8a0656936632de2c1f8`

    To check it:
    ```
    sha256sum ./bcnad
    ```

4) Verify that the version is `v0.3-beta`
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
## Option 2: Setting up Cosmovisor (recommended) 
With Cosmovisor you can schedule updates ahead of time by setting the location of the old binary and the new binary.

Cosmovisor is a **service that listens to the LOGs output and waits** for the correct block to update.

 - ### Step 1. Creating directories for cosmovisor
     ```
    mkdir -p ${HOME}/.bcna/cosmovisor/genesis/bin
    mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/indica/bin
     ```
 
 - ### Step 2. Get Cosmovisor 
1) You can build from the source or download from our github  (not recomended because last version works with SDK v0.43.x):
     ```
     wget https://github.com/BitCannaGlobal/bcna/releases/download/v0.3-beta/cosmovisor
    ```
2) Check the sha256sum. It must return: `622e1f60a94bb93a3810bb7820b9e42c7822086d4996455342b76e1087576d0f`

    To check it:
    ```
    sha256sum ./cosmovisor
    ```
    If it matches then proceed to move to the system path:
    ```
     chmod +x cosmovisor
     sudo mv cosmovisor /usr/local/bin
     ```
 
 - ### Step 3. Get the new binary `v0.3-beta` from BitCanna or build it
 
1) 
    a) If you want to build the binary, instructions to build is on the [README](https://github.com/BitCannaGlobal/bcna#readme) of Github

    b) If you want download the binary:
    ```
    cd ~             #goes to the $HOME folder
    rm -f bcnad      #it deletes if exist the old binary there
    wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v0.3-beta/bcnad
    ```
2) Check the sha256sum. It must return: `1b2f3763dd3f11fe63b0212dc21d20fa09a114760928e8a0656936632de2c1f8`

    To check it:
    ```
    sha256sum ./bcnad
    ```

3) Verify that version is `v0.3-beta`
    ```
    chmod +x bcnad
    ./bcnad version
    ```
 
4) Move the newly built binary to it's directory
     ```
    mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/indica/bin/bcnad
     ```
 
- ### Step 4. Finishing Cosmovisor setup. 

1) Copy the current (`0.2-beta`) binary to Cosmovision genesis (if is the first update) folder:
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
cosmovisor version #show the current  BCNAD version 0.2-beta BUT Will be v0.3-beta after the upgrade block- height
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

## Instructions for validators that participated in the invitational testnet:

Your setup is almost done, you only must to create the new folder structure to hold the new binary under the `indica` release. You don't need to update or download a new version of Cosmovisor (new github releases are related to SDK v.0.43.x versions).

### Step 1. Create the folder infrastructure:
```
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/indica/bin
```
### Step 2. Get the new binary `v0.3-beta` from BitCanna or build it:
 
1) 
    a) If you want to build the binary, instructions to build is on the [README](https://github.com/BitCannaGlobal/bcna#readme) of Github

    b) If you want download the binary:
    ```
    cd ~             #goes to the $HOME folder
    rm -f bcnad      #it deletes if exist the old binary there
    wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v0.3-beta/bcnad
    ```
2) Check the sha256sum. It must return: `1b2f3763dd3f11fe63b0212dc21d20fa09a114760928e8a0656936632de2c1f8`

    To check it:
    ```
    sha256sum ./bcnad
    ```
    
3) Verify that version is `v0.3-beta`

```
chmod +x bcnad
./bcnad version
```
 
4) Move the newly built binary to it's directory
 ```
mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/indica/bin/bcnad
 ```
