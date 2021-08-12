# Task 4 Jail your validator on purpose, then unjail your validator

Jailing occurs when a validator fails to sign blocks for a consecutive period of 24 hours. Getting jailed has consequences for both your balance, and delegators' delegations made to you.
You can read more about this here: https://testnet.bitcanna.io/concepts-and-terminology/slashing

1. Stop signing blocks using the command `sudo service bcnad stop`.
2. Wait for about 24 hours for your validator to be jailed (if you fail to sign blocks for roughly 24 hours you will be jailed) To check the jailed/unjailed status:
`bcnad query staking validator bcnavaloper9430390djoidsajxxxxxxxxxxxxxxxxxx`
3. Send the `unjail` transaction, you can use `bcnad tx --help` to navigate through the command-line to find the correct unjail transaction. Good luck!


There is another fast (an advanced) way to be jailed. It's based in loose the self-delegation stake in the validator. Ask in the Discord group if you are interested. 
