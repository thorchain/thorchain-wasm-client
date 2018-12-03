package exchange

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/thorchain/thorchain-wasm-client/go/client/cosmos-sdk"
)

// OrderKind is the kind an order can have, namely buy or sell
type OrderKind byte

const (
	// BuyOrder is the kind for buy orders
	BuyOrder OrderKind = 0x01
	// SellOrder is the kind for sell orders
	SellOrder OrderKind = 0x02
)

// Create type
type MsgCreateLimitOrder struct {
	Sender    sdk.AccAddress
	Kind      OrderKind
	Amount    sdk.Coin
	Price     sdk.Coin
	ExpiresAt time.Time
}

// new create message
func NewMsgCreateLimitOrder(
	sender sdk.AccAddress, kind OrderKind, amount sdk.Coin, price sdk.Coin, expiresAt time.Time) MsgCreateLimitOrder {
	return MsgCreateLimitOrder{
		Sender:    sender,
		Kind:      kind,
		Amount:    amount,
		Price:     price,
		ExpiresAt: expiresAt,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = MsgCreateLimitOrder{}

// Parser for OrderKind. Returns an error if str is neither "buy" nor "sell"
func ParseKind(str string) (OrderKind, error) {
	if str == "buy" {
		return BuyOrder, nil
	}
	if str == "sell" {
		return SellOrder, nil
	}
	return 0x03, fmt.Errorf("kind must be 'buy' for buy orders or 'sell' for sell orders")
}

//Get MsgCreate Type
func (msg MsgCreateLimitOrder) Type() string { return "exchange" }

//Get Create Signers
func (msg MsgCreateLimitOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgCreateLimitOrder) String() string {
	return fmt.Sprintf("MsgCreateLimitOrder{Sender: %v, Kind: %v, Amount: %v, Price: %v, ExpiresAt: %v}",
		msg.Sender, msg.Kind, msg.Amount, msg.Price, msg.ExpiresAt)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg MsgCreateLimitOrder) ValidateBasic() sdk.Error {
	return nil
}

// Get the bytes for the message signer to sign on
func (msg MsgCreateLimitOrder) GetSignBytes() []byte {
	// ensure expires at is in UTC to have deterministic sign bytes
	msgUtc := msg
	msgUtc.ExpiresAt = msgUtc.ExpiresAt.UTC()

	b, err := json.Marshal(msgUtc)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
