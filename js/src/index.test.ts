import * as fetch from 'node-fetch'
import ThorchainWasmClient from '.'

declare var global: any

global.fetch = fetch

it('works', async () => {
    ThorchainWasmClient({} as any)
})
