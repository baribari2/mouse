package evm

import (
	"log"
	"mouse/calldec/types"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
)

func ExecuteCalldata(tx *types.MouseTx) (err error) {
	cd, err := hexutil.Decode("0x" + strings.Join(tx.Calldata, ""))
	if err != nil {
		return err
	}

	res, _, err := runtime.Execute([]byte(tx.TargetCode), cd, &runtime.Config{
		GasLimit: tx.GasLimit,
		GasPrice: tx.GasPrice,
		Value:    tx.Value,
	})
	if err != nil {
		return err
	}

	tx.ExecuteResult = res

	log.Printf("Byte Res: %v", tx.ExecuteResult)
	log.Printf("String Res: %v", string(tx.ExecuteResult))

	return nil
}
