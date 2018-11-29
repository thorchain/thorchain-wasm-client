package client

import (
	"fmt"
	"syscall/js"

	"github.com/thorchain/thorchain-wasm-client/go/runner"
	// "syscall/js"
)

func RegisterFuncs(r *runner.Runner) {
	r.HandleFunc("decodeAccount", decodeAccount)
	// HandleFunc('pubKeyFromPriv', decodeAccount)
}

// func (b *Bridge) pubKeyFromPriv(args []js.Value) {
// 	privKey := args[0].String()
// 	pubKey := util.PubKeyFromPriv(privKey)

// 	jsonValue, err := b.cdc.MarshalJSON(&pubKey)
// 	fmt.Printf("jsonValue: %+v\n", jsonValue)
// 	if err != nil {
// 		panic(fmt.Sprintf("Unable to get private key: %+v\n", err))
// 	}
// 	jsCallback := b.getJSCallback(args)
// 	jsCallback.Invoke(string(jsonValue))
// 	return
// }

func decodeAccount(args []js.Value) []interface{} {
	strValue := args[0].String()

	return []interface{}{fmt.Sprintf("%v world!", strValue)}

	// respBytes := json.RawMessage(strValue)
	// fmt.Printf("strValue: %+v\n, respBytes: %+v\n", strValue, respBytes)
	// resp := &types.ResultABCIQuery{}
	// util.UnmarshalResponseBytes(b.cdc, respBytes, resp)
	// fmt.Printf("resp.Value: %+v\n", resp.Response.Value)
	// acc, err := util.DecodeAccount(b.cdc, resp.Response.Value)
	// fmt.Printf("account: %+v\n", acc)
	// // jsonValue, err := b.cdc.MarshalJSON(acctBytes)
	// if err != nil {
	// 	panic(err)
	// }
	// jsonValue, err := b.cdc.MarshalJSON(acc)
	// fmt.Printf("jsonValue: %+v\n", jsonValue)
	// jsCallback := b.getJSCallback(args)
	// jsCallback.Invoke(string(jsonValue))
}

// func (b *Bridge) sendMessage(args []js.Value) {
// 	fmt.Printf("args: %+v", args)
// 	from := args[0].String()
// 	to := args[1].String()
// 	coins := args[2].String()
// 	privKeyHex := args[3].String()
// 	txBytes, _ := tx.NewSendTx(from, to, coins, privKeyHex, b.cdc)
// 	jsonValue, _ := b.cdc.MarshalJSON(txBytes)
// 	fmt.Printf("jsonValue: %+v\n", string(jsonValue))
// 	jsCallback := b.getJSCallback(args)
// 	jsCallback.Invoke(string(jsonValue))
// }
