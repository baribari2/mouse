package main

import (
	"errors"
	"fmt"
	"log"
	"mouse/calldec"
	"mouse/calldec/types"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etypes "github.com/ethereum/go-ethereum/core/types"
)

func AnalyzeTx(tx *etypes.Transaction) (mouseTx *types.MouseTx, err error) {
	var ab abi.ABI
	mouseTx = &types.MouseTx{
		Hash:     tx.Hash(),
		GasLimit: tx.Gas(),
		GasPrice: tx.GasPrice(),
		Cost:     tx.Cost(),
		Nonce:    tx.Nonce(),
		To:       tx.To(),
		Value:    tx.Value(),
		Type:     tx.Type(),
	}

	mouseTx, err = calldec.DecodeCalldata(mouseTx, tx.Data())
	if err != nil {
		return nil, err
	}

	log.Printf("%s\x1b[32m%s\x1b[0m%s", "------------------------- ", "Transaction Found", " -------------------------")
	log.Printf("\x1b[32m%s\x1b[0m%v", "Hash: ", mouseTx.Hash)
	log.Printf("\x1b[32m%s\x1b[0m%v", "To: ", mouseTx.To)
	log.Printf("\x1b[32m%s\x1b[0m%v", "Cost: ", mouseTx.Cost)
	log.Printf("\x1b[32m%s\x1b[0m%v", "Gas Limit: ", mouseTx.GasLimit)
	log.Printf("\x1b[32m%s\x1b[0m%v", "Nonce: ", mouseTx.Nonce)

	if len(mouseTx.PossibleSignatures) > 0 {
		log.Printf("\x1b[32m%s\x1b[0m%v", "Function Signature: ", mouseTx.PossibleSignatures[0].TextSignature)
		for _, sig := range mouseTx.PossibleSignatures[1:] {
			log.Printf("\x1b[33m%s\x1b[0m%v", "\tPossible Function Signature: ", sig.TextSignature)
			log.Printf("\x1b[33m%s\x1b[0m%v", "\tPossible Function Arguments: ", sig.Arguments.TextArguments)
		}

		if len(mouseTx.PossibleSignatures[0].Arguments.TextArguments) > 0 {
			log.Printf("\x1b[32m%s\x1b[0m%v", "Arguments: ", mouseTx.PossibleSignatures[0].Arguments.TextArguments)

			as := "[{\"inputs\":["

			for _, sig := range mouseTx.PossibleSignatures {
				for _, arg := range sig.Arguments.ABIArguments {
					as += arg
				}

				as += "],\"name\":\"" + sig.Name + "\",\"outputs\":[{\"name\":\"\", \"type\":\"bytes32\"}],\"stateMutability\":\"public\",\"type\":\"function\"}]"

				// log.Printf("\x1b[32m%s\x1b[0m%v", "ABI String: ", as)

				ab, err = abi.JSON(strings.NewReader(as))
				if err != nil {
					log.Printf("\x1b[31m%s\x1b[0m%v", "failed to parse abi: ", err.Error())
					if errors.New("improperly formatted output") == err {
						continue
					}

					return
				}

			}
		}

		for i, d := range mouseTx.Calldata {
			log.Printf("\x1b[32m%s\x1b[0m%v\n", fmt.Sprintf("\tInput Data (index %v): ", i), d)
		}

		for _, sig := range mouseTx.PossibleSignatures {
			dt, err := hexutil.Decode(sig.HexSignature)
			if err != nil {
				log.Printf("\x1b[31m%s\x1b[0m%v", "failed to decode hex: ", err.Error())
				continue
			}

			_, err = ab.MethodById([]byte(dt))
			if err != nil {
				log.Printf("\x1b[31m%s\x1b[0m%v", "failed to get method by id: ", err.Error())
				continue
			}
		}
		log.Printf("---------------------------------------------------------------------\n")
	}

	return mouseTx, nil
}
