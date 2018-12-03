package sdk

import "github.com/tendermint/tendermint/crypto"

// TxContext implements a transaction context created in SDK modules.
type TxContext struct {
	// Codec         *sdk.Codec
	PrivKey       crypto.PrivKey
	AccountNumber int64
	Sequence      int64
	Gas           int64
	ChainID       string
	Memo          string
	Fee           Coins
}
