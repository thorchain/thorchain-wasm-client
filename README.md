# THORChain WASM Client

Client to communicate with thorchaind from the browser. The code is written in Go and compiled to WebAssembly (WASM).
The reason for this package is the lack of a stable amino encoding library in JavaScript. The biggest drawback to this
approach is the file size of the WASM module, which is >10 MB uncompressed, after manual tree shaking (copying relevant
code into `./go/client/cosmos-sdk` and `./go/client/thorchain` instead of importing the whole modules). For now, this
file size is a secondary concern since the module will be loaded asynchronously in the browser – until a user wants to sign a tx, it should be loaded. It is planned to replace the WASM module with a JS implemtation in the near future.

## Usage in the browser

1. `yarn add thorchain-wasm-client`
2. Add the wasm_exec.js file from `thorchain-wasm-client/dist/wasm_exec.js` to your index.html (see
    `example/src/index.html` for an example)
3. Import the module in your frontend code: `import ThorchainWasmClient from 'thorchain-wasm-client'`
4. Instantiate the client: `const { client, runner } = await ThorchainWasmClient(new Go(), 'http://localhost:26657')`
5. Run methods of the client, e. g.: `const key = await client.createKey` (see `js/src/client/Client.ts` for all
    signatures and `example/src/index.ts` for examples)

## Development

### Installation

1. `make prepare`<sup>1</sup>
2. `cd js && yarn && cd ..`
3. `cd example && yarn && cd ..`

### Adding functionality and debugging it

1. Register your function and implement it in the Go WASM module: `go/client/client.go`
2. Build the Go WASM module: `make build`
3. Add methods to invoke the functions of the Go WASM module from the TypeScript library in `js/src/client/Client.ts`
4. Build the TypeScript library: `cd js && yarn build && cd ..`
5. Create a link to the TypeScript module: `cd js && yarn link && cd ..`
6. Add code to the example in `example/src/index.ts`
7. Use the link that we created above: `cd example && yarn link "thorchain-wasm-client" && cd ..`
8. Run the example: `cd example && yarn start && cd ..`
9. Publish the module: `cd js && yarn publish --access public && cd ..`

<sup>1</sup>This installs dependencies and patches the vendor directory with workarounds for missing WASM support.
Some dependencies do not support WASM yet and produce compilation errors. The vendor_patches folder contains stubs
to allow compilation.
