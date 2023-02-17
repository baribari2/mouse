package types

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// Data structure for easier handling of contract decompilation and transaction decoding
type MouseTx struct {
	// The list of possible function signatures, after seaching a database
	PossibleSignatures []*Signature

	// The hexadecimal representation of the calldata???
	ByteSignature string

	// The text representation of the calldata, chunked by 2
	Calldata []string

	// The raw calldata, as a string
	RawCalldata string

	// The hash of the transaction
	Hash common.Hash

	// The cost of the transaction
	Cost *big.Int

	// The gas limit of the transaction
	GasLimit uint64

	// The gas price of the transaction
	GasPrice *big.Int

	// The nonce of the transaction
	Nonce uint64

	// Whether of not the transaction is protected (sent through a relay)
	Protected bool

	// The target address of the transaction
	To *common.Address

	// The value of the transaction
	Value *big.Int

	// The type of the transaction
	Type uint8

	// The code of the target address of the transaction
	TargetCode string

	// A custom representation of the target contract of the transaction
	Target MouseContract

	// Unimplemented
	ExecuteResult []byte

	// Unimplemented
	DecodedInputs []string

	//Unimplemented
	MouseOrigin string

	sync.Mutex
}

// Data structure for easier handling of function parameters
type Argument struct {
	TextArguments []string
	ABIArguments  []string
}

// Data structure for easier handling of function signatures
type Signature struct {
	// The name of the function
	Name string

	// The text representation of the function signature
	TextSignature string

	// The hexadecimal representation of the function signature
	HexSignature string

	// The function parameters
	Arguments *Argument
}

func NewSignature(name string, textSignature string, hexSignature string, arguments *Argument) *Signature {
	return &Signature{
		Name:          name,
		TextSignature: textSignature,
		HexSignature:  hexSignature,
		Arguments:     arguments,
	}
}
