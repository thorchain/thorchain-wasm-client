# THORChain WASM Client

## Installation

1. `make prepare`<sup>1</sup>
2. `cd js && yarn && cd ..`
3. `cd example && yarn && cd ..`

## Adding functionality and debugging it

1. Register your function and implement it in the Go WASM module: `go/client/client.go`
2. Build the Go WASM module: `make build`
3. Add methods to invoke the functions of the Go WASM module from the TypeScript library in `js/src/client/Client.ts`
4. Build the TypeScript library: `cd js && yarn build && cd ..`
5. Create a link to the TypeScript module: `cd js && yarn link && cd ..`
6. Add code to the example in `example/src/index.ts`
7. Use the link that we created above: `cd example && yarn link "thorchain-wasm-client" && cd ..`
8. Run the example: `cd example && yarn start && cd ..`

<sup>1</sup>This installs dependencies and patches the vendor directory with workarounds for missing WASM support.
The workaround is needed to resolve compile time errors – some dependencies do not support WASM right now. This
project does not depend on any unsupported functionality, so this does not cause problems for now.