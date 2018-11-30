const RPCMethods = {
    ABCI_QUERY: 'abci_query',
    BROADCAST_TX_COMMIT: 'broadcast_tx_commit',
}

export async function request(uri: string, method: string, params: object) {
    const payload = {
        id: 'jsonrpc-client',
        jsonrpc: '2.0',
        method,
        params,
    }

    const resp = await fetch(uri, {
        body: JSON.stringify(payload),
        method: 'POST',
    })

    const json = resp.json()

    return json
}

export async function broadcast(uri: string, signedTx: string) {
    const resp = await request(uri, RPCMethods.BROADCAST_TX_COMMIT, {
        tx: signedTx,
    })
    return resp
}
