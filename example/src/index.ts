import ThorchainWasmClient from 'thorchain-wasm-client'

declare var Go: any

async function main () {
    const { client, runner } = await ThorchainWasmClient(new Go())

    console.log('client ready')

    // succeeding communication
    const result = await client.helloWorld("Hellooo")
    console.log('Hello world result ', result)

    // failing communication
    try {
        const result = await client.helloWorld()
        console.error('Hello world without args should have failed but did not, got result ', result)
    } catch (e) {
        console.log('Hello world without args failed successfully with err ', e)
    }

    // create a new key
    const key = await client.createKey()
    console.log("successfully created key", key, typeof key)

    // get pub key and address from a private key
    const pubAndAddr = await client.getPubAndAddrFromPrivKey(key.priv)
    console.log("successfully got pub and addr from priv", pubAndAddr)

    // get account
    const from = "t0accaddr1778wxtpj6879e8f5wa0kwh3h553kmydzvm5tth"
    const account = await fetch(`http://localhost:1317/accounts/${from}`).then(res => res.json())
    console.log('Successfully got account', account)

    // sign and broadcast a send tx
    const priv_key = "4bD3myAZXUfvoP6cdkfkgwigDzMltovbwcmkNcIxWRiWNYfRcg=="
    const accountNumber = parseInt(account.value.account_number, 10)
    const sequence = parseInt(account.value.sequence, 10)
    const txContext = {
        priv_key,
        account_number: accountNumber,
        sequence,
        gas: 20000,
        chain_id: "test-chain-local"
    }
    const to = "t0accaddr17xhjfa7tj6vzmmwdfa0dcphrsudlrsthwmzfck"
    const amount = "1RUNE"

    const signedSendTx = await client.signSendTx(txContext, from, to, amount)
    const resSignedSendTx = await client.broadcastTx(signedSendTx)
    console.log('Successfully broadcast signedSendTx: ', { from, to, amount, signedSendTx, resSignedSendTx })

    // sign and broadcast a clp trade tx
    const txContext2 = {
        ...txContext,
        sequence: txContext.sequence + 1,
    }
    const fromTicker = 'RUNE'
    const toTicker = 'XMR'
    const fromAmount = 1

    const signedClpTradeTx = await client.signClpTradeTx(txContext2, from, fromTicker, toTicker, fromAmount)
    const resSignedClpTradeTx = await client.broadcastTx(signedClpTradeTx)
    console.log('Successfully broadcast signedClpTradeTx: ', {
        from, fromTicker, toTicker, fromAmount, signedClpTradeTx, resSignedClpTradeTx })

    // sign and broadcast an exchange trade tx
    const txContext3 = {
        ...txContext2,
        sequence: txContext2.sequence + 1,
    }
    const kind = 'buy'
    const amount2 = '2ETH'
    const price = '4RUNE'
    const expiresAt = '2099-12-31T11:45:05.000Z'

    const signedExchangeCreateLimitOrderTx = await client.signExchangeCreateLimitOrderTx(
        txContext3, from, kind, amount2, price, expiresAt)
    const resSignedExchangeCreateLimitOrderTx = await client.broadcastTx(signedExchangeCreateLimitOrderTx)
    console.log('Successfully broadcast SignExchangeCreateLimitOrderTx: ', {
        from, kind, amount2, price, expiresAt, signedExchangeCreateLimitOrderTx, resSignedExchangeCreateLimitOrderTx })
}

main()
