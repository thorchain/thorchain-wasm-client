package sdk

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/tendermint/tendermint/libs/bech32"
)

// Bech32 prefixes
const (
	// Bech32 prefixes
	Bech32PrefixAccAddr = "t0accaddr"
	Bech32PrefixAccPub  = "t0accpub"
	Bech32PrefixValAddr = "t0valaddr"
	Bech32PrefixValPub  = "t0valpub"
)

type AccAddress []byte

// create an AccAddress from a hex string
func AccAddressFromHex(address string) (addr AccAddress, err error) {
	if len(address) == 0 {
		return addr, errors.New("decoding bech32 address failed: must provide an address")
	}
	bz, err := hex.DecodeString(address)
	if err != nil {
		return nil, err
	}
	return AccAddress(bz), nil
}

// create an AccAddress from a bech32 string
func AccAddressFromBech32(address string) (addr AccAddress, err error) {
	bz, err := GetFromBech32(address, Bech32PrefixAccAddr)
	if err != nil {
		return nil, err
	}
	return AccAddress(bz), nil
}

// Allow it to fulfill various interfaces in light-client, etc...
func (bz AccAddress) Bytes() []byte {
	return bz
}

func (bz AccAddress) String() string {
	bech32Addr, err := bech32.ConvertAndEncode(Bech32PrefixAccAddr, bz.Bytes())
	if err != nil {
		panic(err)
	}
	return bech32Addr
}

// decode a bytestring from a bech32-encoded string
func GetFromBech32(bech32str, prefix string) ([]byte, error) {
	if len(bech32str) == 0 {
		return nil, errors.New("decoding bech32 address failed: must provide an address")
	}
	hrp, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
		return nil, err
	}

	if hrp != prefix {
		return nil, fmt.Errorf("invalid bech32 prefix. Expected %s, Got %s", prefix, hrp)
	}

	return bz, nil
}
