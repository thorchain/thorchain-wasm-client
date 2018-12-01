import ThorchainWasmClient from 'thorchain-wasm-client'

declare var Go: any

async function main () {
    const { client, runner } = await ThorchainWasmClient(new Go())

    console.log('client ready')

    // succeeding communication
    const result = await client.helloWorld("Hellooo")
    console.log('got result ', result)

    // failing communication
    try {
        const result = await client.helloWorld()
        console.log('should have failed but did not, got result ', result)
    } catch (e) {
        console.log('failed successfully with err ', e)
    }

    // create a new key
    const key = await client.createKey()
    console.log("successfully created key", key)

    // get account
    const from = "t0accaddr1778wxtpj6879e8f5wa0kwh3h553kmydzvm5tth"
    const account = await fetch(`http://localhost:1317/accounts/${from}`).then(res => res.json())
    console.log('account', account)


    // sign a send tx
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

    const signedTx = await client.signTx(txContext, from, to, amount)
    console.log('successfully signedTx: ', { from, to, amount }, signedTx)

    const sent = await client.broadcastTx(signedTx)
    console.log('sent, got feedback', sent)
}

main()