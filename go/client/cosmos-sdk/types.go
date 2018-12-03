package sdk

// Transactions messages must fulfill the Msg
type Msg interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	Type() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() Error

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []AccAddress
}

// type cmnError = cmn.Error

// sdk Error type
type Error interface {
	// Implements cmn.Error
	// Error() string
	// Stacktrace() cmn.Error
	// Trace(offset int, format string, args ...interface{}) cmn.Error
	// Data() interface{}
	// cmnError

	// convenience
	// TraceSDK(format string, args ...interface{}) Error

	// set codespace
	// WithDefaultCodespace(CodespaceType) Error

	// Code() CodeType
	// Codespace() CodespaceType
	// ABCILog() string
	// ABCICode() ABCICodeType
	// Result() Result
	// QueryResult() abci.ResponseQuery
}
