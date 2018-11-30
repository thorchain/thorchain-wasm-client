package sdk

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"
)

// StdFee includes the amount of coins paid in fees and the maximum
// gas to be used by the transaction. The ratio yields an effective "gasprice",
// which must be above some miminum to be accepted into the mempool.
type StdFee struct {
	Amount Coins `json:"amount"`
	Gas    int64 `json:"gas"`
}

func NewStdFee(gas int64, amount ...Coin) StdFee {
	return StdFee{
		Amount: amount,
		Gas:    gas,
	}
}

// fee bytes for signing later
func (fee StdFee) Bytes() []byte {
	// normalize. XXX
	// this is a sign of something ugly
	// (in the lcd_test, client side its null,
	// server side its [])
	if len(fee.Amount) == 0 {
		fee.Amount = Coins{}
	}
	bz, err := msgCdc.MarshalJSON(fee) // TODO
	if err != nil {
		panic(err)
	}
	return bz
}

//__________________________________________________________

// StdSignDoc is replay-prevention structure.
// It includes the result of msg.GetSignBytes(),
// as well as the ChainID (prevent cross chain replay)
// and the Sequence numbers for each signature (prevent
// inchain replay and enforce tx ordering per account).
type StdSignDoc struct {
	AccountNumber int64             `json:"account_number"`
	ChainID       string            `json:"chain_id"`
	Fee           json.RawMessage   `json:"fee"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Sequence      int64             `json:"sequence"`
}

var msgCdc = NewCodec()

// StdSignBytes returns the bytes to sign for a transaction.
func StdSignBytes(chainID string, accnum int64, sequence int64, fee StdFee, msgs []Msg, memo string) []byte {
	var msgsBytes []json.RawMessage
	for _, msg := range msgs {
		msgsBytes = append(msgsBytes, json.RawMessage(msg.GetSignBytes()))
	}
	bz, err := msgCdc.MarshalJSON(StdSignDoc{
		AccountNumber: accnum,
		ChainID:       chainID,
		Fee:           json.RawMessage(fee.Bytes()),
		Memo:          memo,
		Msgs:          msgsBytes,
		Sequence:      sequence,
	})
	if err != nil {
		panic(err)
	}
	return MustSortJSON(bz)
}

// StdSignMsg is a convenience structure for passing along
// a Msg with the other requirements for a StdSignDoc before
// it is signed. For use in the CLI.
type StdSignMsg struct {
	ChainID       string
	AccountNumber int64
	Sequence      int64
	Fee           StdFee
	Msgs          []Msg
	Memo          string
}

// get message bytes
func (msg StdSignMsg) Bytes() []byte {
	return StdSignBytes(msg.ChainID, msg.AccountNumber, msg.Sequence, msg.Fee, msg.Msgs, msg.Memo)
}

// Standard Signature
type StdSignature struct {
	crypto.PubKey `json:"pub_key"` // optional
	Signature     []byte           `json:"signature"`
	AccountNumber int64            `json:"account_number"`
	Sequence      int64            `json:"sequence"`
}

// Build builds a single message to be signed from a TxContext given a set of
// messages. It returns an error if a fee is supplied but cannot be parsed.
func Build(chainID string, accountNumber, sequence, gas int64, fee Coin, msgs []Msg, memo string) (StdSignMsg, error) {
	if chainID == "" {
		return StdSignMsg{}, errors.Errorf("chain ID required but not specified")
	}

	return StdSignMsg{
		ChainID:       chainID,
		AccountNumber: accountNumber,
		Sequence:      sequence,
		Fee:           NewStdFee(gas, fee),
		Msgs:          msgs,
		Memo:          memo,
	}, nil
}

//Sign a transaction with a given private key
// func PrivSign(priv crypto.PrivKey, msg StdSignMsg) ([]byte, error) {
// 	sig, err := priv.Sign(msg.Bytes())
// 	if err != nil {
// 		return nil, err
// 	}
// 	pubkey := priv.PubKey()

// 	sigs := []StdSignature{{
// 		AccountNumber: msg.AccountNumber,
// 		Sequence:      msg.Sequence,
// 		PubKey:        pubkey,
// 		Signature:     sig,
// 	}}

// 	return txCtx.Codec.MarshalBinary(NewStdTx(msg.Msgs, msg.Fee, sigs, msg.Memo))
// }
