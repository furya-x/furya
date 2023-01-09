# Settings on chain

## Setting examples for admin

### set protocol fee receiver

```bash
katanad tx ledger set-protocol-fee-receiver katana1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --from admin --chain-id local-katana --keyring-backend file

katanad query ledger protocol-fee-receiver 
```

### add new rtoken

```bash
# set rtoken metadata
katanad tx rbank add-denom cosmos cosmosvaloper ./metadata/metadata_ratom.json --chain-id local-katana --from admin --keyring-backend file

katanad query bank denom-metadata

katanad query rbank address-prefix uratom

# set relay fee receiver
katanad tx ledger set-relay-fee-receiver uratom katana1mgjkpyfm00mxk0nmhvfvwhlr65067d538l6cxa --from admin --chain-id local-katana --keyring-backend file

katanad query ledger relay-fee-receiver uratom

# this will init bonded pool, exchange rate, pipeline
katanad tx ledger init-pool uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 --from admin --chain-id local-katana --keyring-backend file

katanad query ledger bonded-pools uratom

katanad query ledger exchange-rate uratom

katanad query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# add relayers
katanad tx relayers add-relayers ledger uratom katana1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx:katana1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --from admin --chain-id local-katana

katanad query relayers relayers ledger uratom

# set threshold
katanad tx relayers set-threshold ledger uratom 1 --from admin --keyring-backend file --chain-id local-katana

katanad query relayers threshold ledger uratom

# set params used by relay
katanad tx ledger set-r-params uratom 0.00001stake 600 0 2 0stake --from admin --keyring-backend file --chain-id local-katana

katanad query ledger r-params uratom

# set pool detail for multisig/ica pool
katanad tx ledger set-pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl:cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x 1 --from admin --chain-id local-katana --keyring-backend file

katanad query ledger pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# default 0.1
katanad tx ledger set-staking-reward-commission uratom 0.15 --from admin --chain-id local-katana --keyring-backend file

katanad q ledger staking-reward-commission uratom

# default 0.002
katanad tx ledger set-unbond-commission uratom 0.0025 --from admin --chain-id local-katana --keyring-backend file

katanad q ledger unbond-commission uratom

# default 1000000ukata
katanad tx ledger set-unbond-relay-fee uratom 1000005ukata --from admin --chain-id local-katana --keyring-backend file

katanad q ledger unbond-relay-fee uratom

```

### register ica pool
```
# register ica pool (need set rtoken metadata before this)
katanad tx ledger register-ica-pool uratom connection-0 --keyring-backend file --from admin --chain-id local-katana --gas 410000

katanad q ledger ica-pool-list uratom

# set withdrawal address
katanad tx ledger set-withdrawal-addr cosmos1gsth46z50w256p4kq36xquh4q90mfjq0t4lm9scln6zucg64epyqudzqzm --keyring-backend file --from admin --chain-id local-katana --gas 410000

```

### rvalidator

```bash
# add relayers
katanad tx relayers add-relayers rvalidator uratom katana14z467aut40mcrt2ddyxf7e74fq99udul7kaf9g:katana15lne70yk254s0pm2da6g59r82cjymzjqvvqxz7 --keyring-backend file --from admin --chain-id local-katana

katanad q relayers relayers rvalidator uratom

# set threshold
katanad tx relayers set-threshold rvalidator uratom 1 --from admin --keyring-backend file --chain-id local-katana

katanad q relayers threshold rvalidator uratom

# init rvalidator (should init target validators of pool before rtoken relay start)
katanad tx rvalidator init-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-katana --keyring-backend file  

katanad q rvalidator r-validator-list uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# add rvalidator
katanad tx rvalidator add-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv  --chain-id local-katana --keyring-backend file --from admin

katanad q rvalidator r-validator-list uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# rm rvalidator
katanad tx rvalidator rm-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-katana --keyring-backend file
```



### bridge

```bash
katanad tx bridge add-chain-id 1 --from admin --keyring-backend file --chain-id local-katana

katanad query bridge chain-ids



katanad tx relayers add-relayers bridge 1 katana1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-katana

katanad query relayers relayers bridge 1



katanad tx relayers set-threshold bridge 1 1 --from admin --keyring-backend file --chain-id local-katana

katanad query relayers threshold bridge 1



katanad tx bridge set-resourceid-to-denom  000000000000000000000000000000a9e0095b8965c01e6a09c97938f3860901 uratom NATIVE --from admin --keyring-backend file --chain-id local-katana

katanad query bridge resourceid-to-denoms



katanad tx bridge set-relay-fee-receiver katana1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-katana

katanad query bridge relay-fee-receiver



katanad tx bridge set-relay-fee 1 1000000ukata --from admin --keyring-backend file --chain-id local-katana

katanad query bridge  relay-fee 1


katanad tx bridge add-banned-denom 1 uratom --from admin --keyring-backend file --chain-id local-katana

katanad q bridge banned-denom-list
```

### migrate rtoken (after adding new rtoken step)

```bash
katanad tx ledger migrate-init uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100000000 150000000 200000000 300000000 1.23 --from admin --keyring-backend file --chain-id local-katana

katanad query bank  total 

katanad query ledger exchange-rate uratom

katanad query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



katanad tx ledger migrate-unbondings uratom --unbondings ./unbondings_example.json --from admin --keyring-backend file --chain-id local-katana

katanad query ledger pool-unbond uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 3



katanad tx bridge set-denom-type uratom  1 --from admin --keyring-backend file --chain-id local-katana

katanad query bridge denom-types
```

### rdex

```bash
katanad tx rdex create-pool 10ukata 20uratom --from admin --chain-id local-katana --keyring-backend file

katanad tx rdex add-provider katana1qzt0qajzr9df3en5sk06xlk26n30003c8uhdkg --from admin --chain-id local-katana --keyring-backend file

katanad tx rdex add-liquidity  100ukata 200uratom --from admin --chain-id local-katana --keyring-backend file

katanad tx rdex remove-liquidity 10 5 1uratom 1ukata ukata --from admin --chain-id local-katana --keyring-backend file

katanad tx rdex swap 2ukata 1uratom  --from admin --chain-id local-katana --keyring-backend file
```

### mining

```bash
katanad tx mining add-mining-provider katana1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx  --from admin --chain-id local-katana --keyring-backend file

katanad tx mining add-reward-token ukata 200 --from admin --chain-id local-katana --keyring-backend file


katanad tx mining add-stake-pool ukata ./add_stake_pool_example.json  --from relay1 --chain-id local-katana --keyring-backend file

katanad tx mining stake 0 10ukata 0 --from my-account --chain-id local-katana --keyring-backend file 

katanad tx mining claim-reward 0 0 --from my-account --chain-id local-katana --keyring-backend file

katanad tx mining add-reward 1 0 300 0 0 --from relay1 --chain-id local-katana --keyring-backend file

katanad tx mining withdraw 1 10ukata 0 --from test --chain-id local-katana --keyring-backend file
```



## Operate examples for user

### liquidity bond (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1000000stake --note 1:katana1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --chain-id local-cosmos
```

### recover (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1stake --note 2:katana1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc:9A80F3E6A007E1144BE34F4A0AC35B9288C19641BCAD3464277168000AF5FC66 --keyring-backend file --chain-id local-cosmos
```

### liquidity unbond

```bash
katanad tx ledger liquidity-unbond cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100uratom cosmos1j9dues7ey2a39nes4ewfvyma96d3f5zrdhnfan --keyring-backend file --from user --home /Users/tpkeeper/gowork/katana/rtoken-relay-core/keys/katana --chain-id local-katana
```

### deposit (transfer token to external chain)

```bash
katanad tx bridge deposit 1 uratom 800 dccf954570063847d73746afa0b0878f2c779d42089c5d9a107f2aca176e985f --from my-account --chain-id local-katana --keyring-backend file
```
