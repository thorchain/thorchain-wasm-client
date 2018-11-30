package client

import (
	"fmt"
	"syscall/js"

	"github.com/thorchain/thorchain-wasm-client/go/client/sdk"
	"github.com/thorchain/thorchain-wasm-client/go/helpers"
	"github.com/thorchain/thorchain-wasm-client/go/runner"
)

func RegisterFuncs(r *runner.Runner) {
	r.HandleFunc("helloWorld", helloWorld)
	r.HandleFunc("signSendTx", signSendTx)
}

func helloWorld(args []js.Value) (interface{}, error) {
	if args[0].Type() != js.TypeString {
		return nil, fmt.Errorf("Arg 0 must be a string, got type %v", args[0].Type())
	}
	strValue := args[0].String()

	return fmt.Sprintf("%v world!", strValue), nil
}

func signSendTx(args []js.Value) (interface{}, error) {
	from, err := helpers.ParseString(args, 0)
	if err != nil {
		return nil, err
	}

	to, err := helpers.ParseString(args, 1)
	if err != nil {
		return nil, err
	}

	amount, err := helpers.ParseString(args, 2)
	if err != nil {
		return nil, err
	}
	coins, err := sdk.ParseCoins(amount)
	if err != nil {
		return nil, err
	}

	chainID, err := helpers.ParseString(args, 3)
	if err != nil {
		return nil, err
	}

	msg := sdk.MsgSend{[]sdk.Input{sdk.Input{[]byte(from), coins}}, []sdk.Output{sdk.Output{[]byte(to), coins}}}

	// TODO use real values
	stdSignMsg, err := sdk.Build(chainID, 0, 0, 20000, sdk.Coin{}, []sdk.Msg{msg}, "")
	if err != nil {
		return nil, err
	}

	return string(stdSignMsg.Bytes()), nil
}
