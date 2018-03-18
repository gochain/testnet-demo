var Web3 = require('web3')

// node1-testnet.gochain.io
var web3 = new Web3(new Web3.providers.HttpProvider("http://138.68.1.11:8545"));

module.exports = web3;
