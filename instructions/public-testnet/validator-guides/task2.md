# Task 2 Connect your validator to Prometheus Analytics

By connecting your validator to Prometheus analytics your node will provide valuable data, which will be shown in a public dashboard to better monitor the network.

We need you to:
1. Enable the prometheus metrics in config.toml
```
nano $HOME/.bcna/config/config.toml
```
Go to the end of the file and change these parameters:
```
#######################################################
###       Instrumentation Configuration Options     ###
#######################################################
[instrumentation]

prometheus = true

prometheus_listen_addr = "0.0.0.0:26660"

```
>in `nano` type CTRL + X to exit + Y to save and exit.
Reset the daemon:
```
sudo service bcnad restart
```
2. Enable a firewall rule to let us collect the data.
For Ubuntu user:
```
sudo ufw allow from 167.172.43.16 proto tcp to any port 26660
```

3. Pass the server IP that runs your node to Raul BitCanna ES#3102 on Discord.

4. Check if your server is working when you receive the OK from the Team:
  * https://monitor.bitcanna.io/d/UJyurCTWz/bitcanna-testnet-dashboard?orgId=2 
  

**Please keep this port open until the end of the Public Testnet.**

