package helpers

import (
	"syscall/js"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/thorchain/thorchain-wasm-client/go/client/sdk"
)

// NewTxContextFromJsValue returns a new initialized TxContext with parameters from the js.Value
func NewTxContextFromJsValue(arg js.Value) (sdk.TxContext, error) {
	privKeyStr, err := ParseStringProp(arg, "priv_key", true)
	if err != nil {
		return sdk.TxContext{}, err
	}
	var privKeyBytes [32]byte
	copy(privKeyBytes[:], privKeyStr)
	privKey := secp256k1.PrivKeySecp256k1(privKeyBytes)

	accountNumber, err := ParseIntProp(arg, "account_number", true)
	if err != nil {
		return sdk.TxContext{}, err
	}

	sequence, err := ParseIntProp(arg, "sequence", true)
	if err != nil {
		return sdk.TxContext{}, err
	}

	gas, err := ParseIntProp(arg, "gas", true)
	if err != nil {
		return sdk.TxContext{}, err
	}

	feeStr, err := ParseStringProp(arg, "fee", false)
	if err != nil {
		return sdk.TxContext{}, err
	}
	var fee sdk.Coins
	if feeStr != "" {
		fee, err = sdk.ParseCoins(feeStr)
		if err != nil {
			return sdk.TxContext{}, err
		}
	} else {
		fee = sdk.Coins{sdk.NewInt64Coin("RUNE", 0)}
	}

	chainID, err := ParseStringProp(arg, "chain_id", true)
	if err != nil {
		return sdk.TxContext{}, err
	}

	memo, err := ParseStringProp(arg, "memo", false)
	if err != nil {
		return sdk.TxContext{}, err
	}

	return sdk.TxContext{
		PrivKey:       privKey,
		AccountNumber: int64(accountNumber),
		Sequence:      int64(sequence),
		Gas:           int64(gas),
		ChainID:       chainID,
		Memo:          memo,
		Fee:           fee,
	}, nil
}
