package types

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type Argument struct {
	TextArguments []string
	ABIArguments  []string
}

type Signature struct {
	Name          string
	TextSignature string
	HexSignature  string
	Arguments     *Argument
}

// Local data structure for easier handling of incoming transactions
type MouseTx struct {
	PossibleSignatures []*Signature
	ByteSignature      string
	Calldata           []string
	Hash               common.Hash
	Cost               *big.Int
	GasLimit           uint64
	GasPrice           *big.Int
	Nonce              uint64
	Protected          bool
	To                 *common.Address
	Value              *big.Int
	Type               uint8
	ExecuteResult      []byte
	TargetCode         string
	DecodedInputs      []string
	sync.Mutex
}
