# Hardware Requirements
Here are the minimal hardware configs required for running a validator/sentry node

* 8GB RAM
* 4vCPUs
* 200GB Disk space

# Software Requirements
Install deps
``` 
sudo apt-get install build-essential jq
```

# Compile instructions

## Install Go 1.15.x (there are some issues with SDK keyring and 1.16.x)
The official instructions can be found here: https://golang.org/doc/install

First remove any existing old Go installation as root
```
sudo rm -rf /usr/local/go
``` 

Download the software:
```
curl https://dl.google.com/go/go1.15.11.linux-amd64.tar.gz | sudo tar -C/usr/local -zxvf -
```
Update environment variables to include go (copy everything and paste)
```
cat <<'EOF' >>$HOME/.profile
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
EOF
source $HOME/.profile
```

## Download source code and compile
```
git clone https://github.com/BitCannaGlobal/testnet-bcna-cosmos.git
cd testnet-bcna-cosmos
git checkout v.testnet5
make build   #it build the binary in build/ folder
make install #it move the binary to $GOBIN path folder
```
