# STAGE 2. Task 2. - Get the definitive Genesis file and start your validator.


### Step 1. Get the definitive Genesis file and configure the server.

> Time interval: From 16th Wednesday 10.00h UTC to 17th Thursday  13.59h UTC

1. Fetch genesis.json into bcna's config directory, it will automatically overwrite the genesis.json from stage 2, task 1.
```
curl -s https://raw.githubusercontent.com/BitCannaGlobal/testnet-bcna-cosmos/main/instructions/stage2/genesis.json > ~/.bcna/config/genesis.json
```

2. Ensure you retrieve the right file. The md5sum must be `45c19f9448e0afcd44573b1addfe1223`

```
md5sum ~/.bcna/config/genesis.json 
45c19f9448e0afcd44573b1addfe1223  /home/user/.bcna/config/genesis.json
```

### Step 2. Start the daemon and wait.

> The Bitcanna team will be present in the Discord Testnet channel to coordinate and help with issues.

1. Reset the old chain, start your node and leave it online for genesis time. You will not need to be present or awake at genesis time, as long as your node is running (but being online might be beneficial in case problems occur).

```
bcnad unsafe-reset-all
sudo systemctl start bcnad
```

2. Monitor your log to check the consensus/launch (CTRL + C to exit):
```
sudo journalctl -u bcnad -f
```

Wait until the date/time specified in the Genesis file and the chain will start automatically.

You can expect a message like this:
```
3:49PM INF ABCI Replay Blocks appHeight=0 module=consensus stateHeight=0 storeHeight=0
3:49PM INF asserting crisis invariants inv=0/11 module=x/crisis name=staking/module-accounts
3:49PM INF asserting crisis invariants inv=1/11 module=x/crisis name=staking/nonnegative-power
3:49PM INF asserting crisis invariants inv=2/11 module=x/crisis name=staking/positive-delegation
3:49PM INF asserting crisis invariants inv=3/11 module=x/crisis name=staking/delegator-shares
3:49PM INF asserting crisis invariants inv=4/11 module=x/crisis name=distribution/nonnegative-outstanding
3:49PM INF asserting crisis invariants inv=5/11 module=x/crisis name=distribution/can-withdraw
3:49PM INF asserting crisis invariants inv=6/11 module=x/crisis name=distribution/reference-count
3:49PM INF asserting crisis invariants inv=7/11 module=x/crisis name=distribution/module-account
3:49PM INF asserting crisis invariants inv=8/11 module=x/crisis name=gov/module-account
3:49PM INF asserting crisis invariants inv=9/11 module=x/crisis name=bank/nonnegative-outstanding
3:49PM INF asserting crisis invariants inv=10/11 module=x/crisis name=bank/total-supply
3:49PM INF asserted all invariants duration=1.475903 height=0 module=x/crisis
3:49PM INF created new capability module=ibc name=ports/transfer
3:49PM INF port binded module=x/ibc/port port=transfer
3:49PM INF claimed capability capability=1 module=transfer name=ports/transfer
3:49PM INF Completed ABCI Handshake - Tendermint and App are synced appHash= appHeight=0 module=consensus
3:49PM INF Version info block=11 p2p=8 software=v0.34.10
3:49PM INF This node is not a validator addr=61E121975FE49F3ADA61550CE3E8CC080CAFFCBF module=consensus pubKey=V+XvGRmBwzv6DOq6LhW3+NAgrV9qptkY5LPcnKYupTo=
3:49PM INF P2P Node ID ID=7c030b0172dcc19fad261dac5601c1282282e747 file=/home/linux/.bcna/config/node_key.json module=p2p
3:49PM INF Adding persistent peers addrs=["8e241ba2e8db2e83bb5d80473b4fd4d901043dda@178.128.247.173:26656","a76427360ea0418986a86c574432f40eb7258d80@159.65.198.245:26656"] module=p2p
3:49PM INF Adding unconditional peer ids ids=[] module=p2p
3:49PM INF Add our address to book addr={"id":"7c030b0172dcc19fad261dac5601c1282282e747","ip":"0.0.0.0","port":26656} book=/home/linux/.bcna/config/addrbook.json module=p2p
3:49PM INF Starting Node service impl=Node
3:49PM INF Genesis time is in the future. Sleeping until then... genTime=2021-06-17T14:00:00Z
3:49PM INF Starting pprof server laddr=localhost:6060

```

## Step 3. The Launch 2021-06-17T14:00:00  UTC
> The Bitcanna team will be present in the Discord Testnet channel to coordinate and help with issues.

To monitor the consensus, we can check the status of the votes on the current proposed block (you can open another terminal):

```
curl http://localhost:26657/consensus_state | jq '.result.round_state.height_vote_set[0].prevotes_bit_array'
```

The goal is to get >66% for each block to commit.

```
"BA{138:xxxxx_x_xxxxxx_xxxxxxxx__x___xxxxxx_x_x_x__xxxxx_x_xxx_xxxx_xxxx_x_xx__xx__x_xxxx__x__xxx_x_x_x_x_xxxxx_x_xxx___xxx_xxx_xx_xxx_x_xxx_x_xxx} 4488/6684 = 0.67"
```
0% is normal after a new block has begun.
```
"BA{138:__________________________________________________________________________________________________________________________________________} 0/6684 = 0.00" 
```

In between, we will be waiting for more validators to join and vote (the vote is automatic, you only need to start the chain).


# Let's hope for a good launch! :rocket: :moon:
