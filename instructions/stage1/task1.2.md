# STAGE 1.Task 2. Connect your validator to our Prometheus analytics.

We want collect as many metrics as we can for monitoring.
We need you to:
1. Enable the prometheus metrics in config.toml
```
nano $HOME/.bcna/config/config.toml
```
Go to the end of the file and change this params:
```
#######################################################
###       Instrumentation Configuration Options     ###
#######################################################
[instrumentation]

# When true, Prometheus metrics are served under /metrics on
# PrometheusListenAddr.
# Check out the documentation for the list of available metrics.
prometheus = true

# Address to listen for Prometheus collector(s) connections
prometheus_listen_addr = "0.0.0.0:26660"

# Maximum number of simultaneous connections.
# If you want to accept a larger number than the default, make sure
# you increase your OS limits.
# 0 - unlimited.
max_open_connections = 3

# Instrumentation namespace
namespace = "tendermint"
```
in `nano` type CTRL + X to exit + Y to save and exit.
Reset the daemon:
```
sudo service bcnad restart
```
2. Enable a firewall rule to let us collect the data.
For Ubuntu user:
```
sudo ufw allow from 167.172.43.16 proto tcp to any port 26660
```
**Please keep this port open until the end of the Testnet.**

3. Submit your public IP in the Task Center Dashboard (proof of task).
