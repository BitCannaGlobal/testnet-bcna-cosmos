# STAGE 2. TASK 3. Never be jailed.

It's very difficult with the current parameters to be jailed.

```
        "signed_blocks_window": "20000",
        "min_signed_per_window": "0.500000000000000000",
```

It means you need miss 10000 continuous blocks to be jailed (you need to be disconnected for roughly 16 hours)

To avoid this jailing, you will have to monitor your validator and reconnect if you lost connection for example.

**Please, feel free to contribute making a Pull Request**

## Some tips to help to improve your server/service/peer/validator response.

### 1. Increase the value  `LimitNOFILE`  to `LimitNOFILE=8192` at service file `/lib/systemd/system/bcnad.service`.
```
sudo nano /lib/systemd/system/bcnad.service
```

Don't forget to reload the systemd and restart the service bcnad
```
sudo systemctl reenable  bcnad.service && sudo systemctl start bcnad.service
```

### 2. There are some params that you can change in config.toml

#### [P2P] Part

You can increase these values to get connected to more peers

```
# Maximum number of inbound peers
max_num_inbound_peers = 40

# Maximum number of outbound peers to connect to, excluding persistent peers
max_num_outbound_peers = 10
```
You can increase these values to get connected faster to peers:

```
# Rate at which packets can be sent, in bytes/second
send_rate = 5120000

# Rate at which packets can be received, in bytes/second
recv_rate = 5120000
```

#### [mempool] part

**Be careful with this part. The recommendable is leave as is by default. But IF YOUR VALIDATOR is stucked, you can increase or decrease this values to try and optimize. Don't touch if everthing goes right.**

You can increase the size to accept more or decrease to reject
```
# Maximum number of transactions in the mempool
size = 5000

# Limit the total size of all txs in the mempool.
# This only accounts for raw transactions (e.g. given 1MB transactions and

# max_txs_bytes=5MB, mempool will only accept 5 transactions).
max_txs_bytes = 1073741824

# Maximum size of a single transaction.
# NOTE: the max size of a tx transmitted over the network is {max_tx_bytes}.
max_tx_bytes = 5048576
```
Also you could manipulate the `max_txs_bytes` & `max_tx_bytes`

## Summarizing 

The default parameters used to be right. The most important aspect to avoid get jailed is to have a good:
* Bandwith 
* CPU cores
* Memory 
* quick SSD or nve disk

Don't forget to restart the daemon (`sudo service bcnad restart`)

## Create a swapfile for the system to breathe in tought times (eg: spam, busy network)

Check if your system already has a Swapfile

`sudo swapon --show `

if no results, it means you have no swapfile.

Allocate space to the swapfile (in this case, 4GB)

`sudo fallocate -l 4G /swapfile`


Give the swapfile the correct permissions - owner only.

`sudo chmod 600 /swapfile`


Set up the swap area

`sudo mkswap /swapfile`


Active swap

`sudo swapon /swapfile`


Now make it permanent

`sudo pico /etc/fstab` - and paste this line

` /swapfile swap swap defaults 0 0`

Verify swap creation: 

`sudo swapon --show` or `sudo free -h` , also shows in `htop`


### Now, configure kernel module for swap support

Add to swappiness kernel

`cat /proc/sys/vm/swappiness` #output should be 60 or near, which is good, but we need a lower value for production so let's change it to 10:

`sudo sysctl vm.swappiness=10`


Make the change permanent

```sudo pico /etc/sysctl.conf ``` - and add this line at the bottom vm.swappiness=10


Reboot your system check again with `sudo swapon --show` or `sudo free -h`

If you did all the steps correctly it should be fine.

