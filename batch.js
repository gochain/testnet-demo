web3 = require('./init.js')

exports.run = async function(account1, account2, numTxs) {
    console.log("Sending", numTxs, "from", account1.address, "to", account2.address)
    let response = await web3.eth.getTransactionCount(account1.address);
    console.log("nonce/txcount:", response); 
    account1.nonce = response;
    for (let i = 0; i < numTxs; i++){
        try {
            let tx = {to: account2.address, value: web3.utils.toWei('1', 'ether'), nonce: account1.nonce, gas: '2000000'}
            response = await web3.eth.accounts.signTransaction(tx, account1.privateKey);
            console.log("signed:", response); 
            tx.signed = response;

            web3.eth.sendSignedTransaction(tx.signed.rawTransaction).on('receipt', function(response){console.log("send tx response:", response)}).then(function(response){ console.log(response) }).catch(console.log)
            
            // update nonce
            account1.nonce++;

        } catch(err) {
            console.log('ERROR', err)
        }
    }
}
