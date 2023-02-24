# Settings on chain

## Setting examples for admin

### set protocol fee receiver

```bash
furyad tx ledger set-protocol-fee-receiver furya1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --from admin --chain-id local-furya --keyring-backend file

furyad query ledger protocol-fee-receiver 
```

### add new rtoken

```bash
# set rtoken metadata
furyad tx rbank add-denom cosmos cosmosvaloper ./metadata/metadata_xatom.json --chain-id local-furya --from admin --keyring-backend file

furyad query bank denom-metadata

furyad query rbank address-prefix uxatom

# set relay fee receiver
furyad tx ledger set-relay-fee-receiver uxatom furya1mgjkpyfm00mxk0nmhvfvwhlr65067d538l6cxa --from admin --chain-id local-furya --keyring-backend file

furyad query ledger relay-fee-receiver uxatom

# this will init bonded pool, exchange rate, pipeline
furyad tx ledger init-pool uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 --from admin --chain-id local-furya --keyring-backend file

furyad query ledger bonded-pools uxatom

furyad query ledger exchange-rate uxatom

furyad query ledger bond-pipeline uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# add relayers
furyad tx relayers add-relayers ledger uxatom furya1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx:furya1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --from admin --chain-id local-furya

furyad query relayers relayers ledger uxatom

# set threshold
furyad tx relayers set-threshold ledger uxatom 1 --from admin --keyring-backend file --chain-id local-furya

furyad query relayers threshold ledger uxatom

# set params used by relay
furyad tx ledger set-r-params uxatom 0.00001stake 600 0 2 0stake --from admin --keyring-backend file --chain-id local-furya

furyad query ledger r-params uxatom

# set pool detail for multisig/ica pool
furyad tx ledger set-pool-detail uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl:cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x 1 --from admin --chain-id local-furya --keyring-backend file

furyad query ledger pool-detail uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# default 0.1
furyad tx ledger set-staking-reward-commission uxatom 0.15 --from admin --chain-id local-furya --keyring-backend file

furyad q ledger staking-reward-commission uxatom

# default 0.002
furyad tx ledger set-unbond-commission uxatom 0.0025 --from admin --chain-id local-furya --keyring-backend file

furyad q ledger unbond-commission uxatom

# default 1000000ufury
furyad tx ledger set-unbond-relay-fee uxatom 1000005ufury --from admin --chain-id local-furya --keyring-backend file

furyad q ledger unbond-relay-fee uxatom

```

### register ica pool
```
# register ica pool (need set rtoken metadata before this)
furyad tx ledger register-ica-pool uxatom connection-0 --keyring-backend file --from admin --chain-id local-furya --gas 410000

furyad q ledger ica-pool-list uxatom

# set withdrawal address
furyad tx ledger set-withdrawal-addr cosmos1gsth46z50w256p4kq36xquh4q90mfjq0t4lm9scln6zucg64epyqudzqzm --keyring-backend file --from admin --chain-id local-furya --gas 410000

```

### rvalidator

```bash
# add relayers
furyad tx relayers add-relayers rvalidator uxatom furya14z467aut40mcrt2ddyxf7e74fq99udul7kaf9g:furya15lne70yk254s0pm2da6g59r82cjymzjqvvqxz7 --keyring-backend file --from admin --chain-id local-furya

furyad q relayers relayers rvalidator uxatom

# set threshold
furyad tx relayers set-threshold rvalidator uxatom 1 --from admin --keyring-backend file --chain-id local-furya

furyad q relayers threshold rvalidator uxatom

# init rvalidator (should init target validators of pool before rtoken relay start)
furyad tx rvalidator init-r-validator uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-furya --keyring-backend file  

furyad q rvalidator r-validator-list uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# add rvalidator
furyad tx rvalidator add-r-validator uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv  --chain-id local-furya --keyring-backend file --from admin

furyad q rvalidator r-validator-list uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# rm rvalidator
furyad tx rvalidator rm-r-validator uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-furya --keyring-backend file
```



### bridge

```bash
furyad tx bridge add-chain-id 1 --from admin --keyring-backend file --chain-id local-furya

furyad query bridge chain-ids



furyad tx relayers add-relayers bridge 1 furya1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-furya

furyad query relayers relayers bridge 1



furyad tx relayers set-threshold bridge 1 1 --from admin --keyring-backend file --chain-id local-furya

furyad query relayers threshold bridge 1



furyad tx bridge set-resourceid-to-denom  000000000000000000000000000000a9e0095b8965c01e6a09c97938f3860901 uxatom NATIVE --from admin --keyring-backend file --chain-id local-furya

furyad query bridge resourceid-to-denoms



furyad tx bridge set-relay-fee-receiver furya1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-furya

furyad query bridge relay-fee-receiver



furyad tx bridge set-relay-fee 1 1000000ufury --from admin --keyring-backend file --chain-id local-furya

furyad query bridge  relay-fee 1


furyad tx bridge add-banned-denom 1 uxatom --from admin --keyring-backend file --chain-id local-furya

furyad q bridge banned-denom-list
```

### migrate rtoken (after adding new rtoken step)

```bash
furyad tx ledger migrate-init uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100000000 150000000 200000000 300000000 1.23 --from admin --keyring-backend file --chain-id local-furya

furyad query bank  total 

furyad query ledger exchange-rate uxatom

furyad query ledger bond-pipeline uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



furyad tx ledger migrate-unbondings uxatom --unbondings ./unbondings_example.json --from admin --keyring-backend file --chain-id local-furya

furyad query ledger pool-unbond uxatom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 3



furyad tx bridge set-denom-type uxatom  1 --from admin --keyring-backend file --chain-id local-furya

furyad query bridge denom-types
```

### rdex

```bash
furyad tx rdex create-pool 10ufury 20uxatom --from admin --chain-id local-furya --keyring-backend file

furyad tx rdex add-provider furya1qzt0qajzr9df3en5sk06xlk26n30003c8uhdkg --from admin --chain-id local-furya --keyring-backend file

furyad tx rdex add-liquidity  100ufury 200uxatom --from admin --chain-id local-furya --keyring-backend file

furyad tx rdex remove-liquidity 10 5 1uxatom 1ufury ufury --from admin --chain-id local-furya --keyring-backend file

furyad tx rdex swap 2ufury 1uxatom  --from admin --chain-id local-furya --keyring-backend file
```

### mining

```bash
furyad tx mining add-mining-provider furya1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx  --from admin --chain-id local-furya --keyring-backend file

furyad tx mining add-reward-token ufury 200 --from admin --chain-id local-furya --keyring-backend file


furyad tx mining add-stake-pool ufury ./add_stake_pool_example.json  --from relay1 --chain-id local-furya --keyring-backend file

furyad tx mining stake 0 10ufury 0 --from my-account --chain-id local-furya --keyring-backend file 

furyad tx mining claim-reward 0 0 --from my-account --chain-id local-furya --keyring-backend file

furyad tx mining add-reward 1 0 300 0 0 --from relay1 --chain-id local-furya --keyring-backend file

furyad tx mining withdraw 1 10ufury 0 --from test --chain-id local-furya --keyring-backend file
```



## Operate examples for user

### liquidity bond (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1000000stake --note 1:furya1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --chain-id local-cosmos
```

### recover (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1stake --note 2:furya1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc:9A80F3E6A007E1144BE34F4A0AC35B9288C19641BCAD3464277168000AF5FC66 --keyring-backend file --chain-id local-cosmos
```

### liquidity unbond

```bash
furyad tx ledger liquidity-unbond cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100uxatom cosmos1j9dues7ey2a39nes4ewfvyma96d3f5zrdhnfan --keyring-backend file --from user --home /Users/tpkeeper/gowork/furya/rtoken-relay-core/keys/furya --chain-id local-furya
```

### deposit (transfer token to external chain)

```bash
furyad tx bridge deposit 1 uxatom 800 dccf954570063847d73746afa0b0878f2c779d42089c5d9a107f2aca176e985f --from my-account --chain-id local-furya --keyring-backend file
```
