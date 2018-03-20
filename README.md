# Testing GoChain

This will walk you through creating wallets and sending some transactions using the console.

## Setup

* Node 9.X installed.
* Clone this repo: `git clone https://github.com/gochain-io/testnet-demo.git` then `cd testnet-demo`
* Run `npm install`

## Demo

Run `node`, then in console:

```sh
web3 = require('./init.js')
# Create account
account1 = web3.eth.accounts.create()
# Save the address and privateKey you get back, could also make these on myetherwallet or other tools
# Send us a message on telegram and we'll send you some GOC to use
# Can recover an account with (ensure 0x in the beginning of privateKey): `account = web3.eth.accounts.privateKeyToAccount(privateKey)`

# Make a second account to send tx to
account2 = web3.eth.accounts.create()

# Make and send a transaction - just copy and paste each line directly
web3.eth.getTransactionCount(account1.address).then(function(response){console.log(response); account1.nonce = response;})
tx = {to: account2.address, value: web3.utils.toWei('1', 'ether'), nonce: account1.nonce, gas: '2000000'}
web3.eth.accounts.signTransaction(tx, account1.privateKey).then(function(response){ console.log(response); tx.signed = response; })
web3.eth.sendSignedTransaction(tx.signed.rawTransaction).on('receipt', console.log).then(function(response){ console.log(response) }).catch(console.log)

# Now check the new balances
web3.eth.getBalance(account1.address).then(function(response){console.log(web3.utils.fromWei(response, 'ether'))})
web3.eth.getBalance(account2.address).then(function(response){console.log(web3.utils.fromWei(response, 'ether'))})
```

BOOM!

### Batch time

Now let's send a bunch of txs at once. This will send from `account1` to `account2`.

```sh
batch = require('./batch.js')
batch.run(account1, account2, 500)

# Wait for tx's to complete, then check balances again
web3.eth.getBalance(account1.address).then(function(response){console.log(web3.utils.fromWei(response, 'ether'))})
web3.eth.getBalance(account2.address).then(function(response){console.log(web3.utils.fromWei(response, 'ether'))})
```
