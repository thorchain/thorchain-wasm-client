package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"syscall/js"
	"time"

	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	sdk "github.com/thorchain/thorchain-wasm-client/go/client/cosmos-sdk"
	"github.com/thorchain/thorchain-wasm-client/go/client/thorchain/clp"
	"github.com/thorchain/thorchain-wasm-client/go/client/thorchain/exchange"
	"github.com/thorchain/thorchain-wasm-client/go/helpers"
	"github.com/thorchain/thorchain-wasm-client/go/runner"
)

func RegisterFuncs(r *runner.Runner) {
	r.HandleFunc("createKey", createKey)
	r.HandleFunc("getPubAndAddrFromPrivKey", getPubAndAddrFromPrivKey)
	r.HandleFunc("helloWorld", helloWorld)
	r.HandleFunc("signSendTx", signSendTx)
	r.HandleFunc("signClpTradeTx", signClpTradeTx)
	r.HandleFunc("signExchangeCreateLimitOrderTx", signExchangeCreateLimitOrderTx)
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

func getPubAndAddrFromPrivKey(args []js.Value) (interface{}, error) {
	privStr, err := helpers.ParseString(args, 0)
	if err != nil {
		return nil, err
	}

	privBytes, err := base64.StdEncoding.DecodeString(privStr)
	if err != nil {
		return nil, err
	}

	priv, err := cryptoAmino.PrivKeyFromBytes(privBytes)
	if err != nil {
		return nil, err
	}

	pub := priv.PubKey()
	addr := sdk.AccAddress(pub.Address())

	bz, err := json.Marshal(key{Priv: priv.Bytes(), Pub: pub.Bytes(), Addr: addr.String()})
	if err != nil {
		return nil, err
	}

	return string(bz), nil
}

func signSendTx(args []js.Value) (interface{}, error) {
	txContextObj, err := helpers.ParseObject(args, 0)
	if err != nil {
		return nil, err
	}
	txContext, err := helpers.NewTxContextFromJsValue(txContextObj)
	if err != nil {
		return nil, err
	}

	fromStr, err := helpers.ParseString(args, 1)
	if err != nil {
		return nil, err
	}
	from, err := sdk.AccAddressFromBech32(fromStr)
	if err != nil {
		return nil, err
	}

	toStr, err := helpers.ParseString(args, 2)
	if err != nil {
		return nil, err
	}
	to, err := sdk.AccAddressFromBech32(toStr)
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

	stdSignMsg, err := sdk.Build(txContext, []sdk.Msg{msg})
	if err != nil {
		return nil, err
	}

	//sign
	txBytes, err := sdk.PrivSign(cdc, txContext.PrivKey, stdSignMsg)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.EncodeToString(txBytes), nil
}

func signClpTradeTx(args []js.Value) (interface{}, error) {
	txContextObj, err := helpers.ParseObject(args, 0)
	if err != nil {
		return nil, err
	}
	txContext, err := helpers.NewTxContextFromJsValue(txContextObj)
	if err != nil {
		return nil, err
	}

	fromStr, err := helpers.ParseString(args, 1)
	if err != nil {
		return nil, err
	}

	from, err := sdk.AccAddressFromBech32(fromStr)
	if err != nil {
		return nil, err
	}

	fromTicker, err := helpers.ParseString(args, 2)
	if err != nil {
		return nil, err
	}
	toTicker, err := helpers.ParseString(args, 3)
	if err != nil {
		return nil, err
	}

	fromAmount, err := helpers.ParseInt(args, 4)
	if err != nil {
		return nil, err
	}

	msg := clp.NewMsgTrade(from, fromTicker, toTicker, fromAmount)

	stdSignMsg, err := sdk.Build(txContext, []sdk.Msg{msg})
	if err != nil {
		return nil, err
	}

	// printCdcTypes(cdc)

	//sign
	txBytes, err := sdk.PrivSign(cdc, txContext.PrivKey, stdSignMsg)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.EncodeToString(txBytes), nil
}

func signExchangeCreateLimitOrderTx(args []js.Value) (interface{}, error) {
	txContextObj, err := helpers.ParseObject(args, 0)
	if err != nil {
		return nil, err
	}
	txContext, err := helpers.NewTxContextFromJsValue(txContextObj)
	if err != nil {
		return nil, err
	}

	senderStr, err := helpers.ParseString(args, 1)
	if err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(senderStr)
	if err != nil {
		return nil, err
	}

	kindStr, err := helpers.ParseString(args, 2)
	if err != nil {
		return nil, err
	}
	kind, err := exchange.ParseKind(kindStr)
	if err != nil {
		return nil, err
	}

	amountStr, err := helpers.ParseString(args, 3)
	if err != nil {
		return nil, err
	}
	amount, err := sdk.ParseCoin(amountStr)
	if err != nil {
		return nil, err
	}

	priceStr, err := helpers.ParseString(args, 4)
	if err != nil {
		return nil, err
	}
	price, err := sdk.ParseCoin(priceStr)
	if err != nil {
		return nil, err
	}

	expiresAtStr, err := helpers.ParseString(args, 5)
	if err != nil {
		return nil, err
	}
	expiresAt, err := time.Parse(time.RFC3339, expiresAtStr)
	if err != nil {
		return nil, err
	}

	msg := exchange.NewMsgCreateLimitOrder(sender, kind, amount, price, expiresAt)

	stdSignMsg, err := sdk.Build(txContext, []sdk.Msg{msg})
	if err != nil {
		return nil, err
	}

	// printCdcTypes(cdc)

	//sign
	txBytes, err := sdk.PrivSign(cdc, txContext.PrivKey, stdSignMsg)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.EncodeToString(txBytes), nil
}

//__________________________________________________________________

var cdc *sdk.Codec

func init() {
	cdc = sdk.NewCodec()
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterConcrete(sdk.StdTx{}, "auth/StdTx", nil)
	cdc.RegisterConcrete(sdk.MsgSend{}, "cosmos-sdk/Send", nil)
	cdc.RegisterConcrete(clp.MsgTrade{}, "clp/MsgTrade", nil)
	cdc.RegisterConcrete(exchange.MsgCreateLimitOrder{}, "exchange/MsgCreateLimitOrder", nil)
	sdk.RegisterCrypto(cdc)
}

func printCdcTypes(cdc *sdk.Codec) {
	var b bytes.Buffer
	cdc.PrintTypes(&b)
	fmt.Println("", b.String())
}
