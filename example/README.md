# Thorchain WebAssembly Bridge Example

This example requires the thorchaind and thorchaincli installed locally, see https://github.com/thorchain/THORChain

## Run the example

1. Build the bridge: `cd .. && make prepare && make build && cd example`
2. Build the bridge: `cd ../js && yarn build && cd ../example`
3. Run thorchain node: `thorchaind start`
3. Create a CLP for the example: `thorchaincli clp create XMR "Monero" 18 100 5000000 5000000 --chain-id test-chain-local --from local_validator`
4. Run thorchain LCD: `thorchaincli advanced rest-server`
5. Install dependencies: `yarn`
6. Serve the example: `yarn start`
