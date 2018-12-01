package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/thorchain/thorchain-wasm-client/go/client/sdk"
	"github.com/thorchain/thorchain-wasm-client/go/helpers"
	"github.com/thorchain/thorchain-wasm-client/go/runner"
)

func RegisterFuncs(r *runner.Runner) {
	r.HandleFunc("createKey", createKey)
	r.HandleFunc("helloWorld", helloWorld)
	r.HandleFunc("signTx", signTx)
}

func helloWorld(args []js.Value) (interface{}, error) {
	if args[0].Type() != js.TypeString {
		return nil, fmt.Errorf("Arg 0 must be a string, got type %v", args[0].Type())
	}
	strValue := args[0].String()

	return fmt.Sprintf("%v world!", strValue), nil
}

type key struct {
	Priv []byte `json:"priv"`
	Pub  []byte `json:"pub"`
	Addr string `json:"addr"`
}

func createKey(args []js.Value) (interface{}, error) {
	priv := secp256k1.GenPrivKey()
	pub := priv.PubKey()
	addr := sdk.AccAddress(pub.Address())

	bz, err := json.Marshal(key{Priv: priv.Bytes(), Pub: pub.Bytes(), Addr: addr.String()})
	if err != nil {
		return nil, err
	}

	return string(bz), nil
}

func signTx(args []js.Value) (interface{}, error) {
	txContextObj, err := helpers.ParseObject(args, 0)
	if err != nil {
		return nil, err
	}
	txContext, err := helpers.NewTxContextFromJsValue(txContextObj)
	if err != nil {
		return nil, err
	}

	from, err := helpers.ParseString(args, 1)
	if err != nil {
		return nil, err
	}

	to, err := helpers.ParseString(args, 2)
	if err != nil {
		return nil, err
	}

	amount, err := helpers.ParseString(args, 3)
	if err != nil {
		return nil, err
	}
	coins, err := sdk.ParseCoins(amount)
	if err != nil {
		return nil, err
	}

	msg := sdk.MsgSend{[]sdk.Input{sdk.Input{[]byte(from), coins}}, []sdk.Output{sdk.Output{[]byte(to), coins}}}

	fmt.Println("msg", msg)

	stdSignMsg, err := sdk.Build(txContext, []sdk.Msg{msg})
	if err != nil {
		return nil, err
	}

	fmt.Println("stdSignMsg", stdSignMsg)

	priv := secp256k1.GenPrivKey()

	// printCdcTypes(cdc)

	//sign
	txBytes, err := sdk.PrivSign(cdc, priv, stdSignMsg)
	if err != nil {
		return nil, err
	}

	fmt.Println("txBytes", string(txBytes))

	return base64.StdEncoding.EncodeToString(txBytes), nil
}

//__________________________________________________________________

var cdc *sdk.Codec

func init() {
	cdc = sdk.NewCodec()
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterConcrete(sdk.StdTx{}, "auth/StdTx", nil)
	cdc.RegisterConcrete(sdk.MsgSend{}, "cosmos-sdk/Send", nil)
	sdk.RegisterCrypto(cdc)
}

func printCdcTypes(cdc *sdk.Codec) {
	var b bytes.Buffer
	cdc.PrintTypes(&b)
	fmt.Println("", b.String())
}
