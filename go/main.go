package main

import (
	"github.com/thorchain/thorchain-wasm-client/go/client"
	"github.com/thorchain/thorchain-wasm-client/go/runner"
)

func main() {
	c := make(chan struct{}, 0)

	r := runner.NewRunner(namespace)
	client.RegisterFuncs(&r)

	println("thorchain_wasm_client initialized")

	<-c
}
