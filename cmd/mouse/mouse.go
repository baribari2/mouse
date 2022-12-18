package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mouse/calldec"
	"mouse/calldec/types"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AnalyzeTx(ec *ethclient.Client, tx *etypes.Transaction) (mouseTx *types.MouseTx, err error) {
	var ab abi.ABI
	ca, err := ec.CodeAt(context.Background(), *tx.To(), nil)
	if err != nil {
		return nil, err
	}

	if ca == nil {
		return nil, errors.New("no contract code found")
	}

	mouseTx = &types.MouseTx{
		Hash:       tx.Hash(),
		GasLimit:   tx.Gas(),
		GasPrice:   tx.GasPrice(),
		Cost:       tx.Cost(),
		Nonce:      tx.Nonce(),
		To:         tx.To(),
		Value:      tx.Value(),
		Type:       tx.Type(),
		TargetCode: hexutil.Encode(ca)[3:],
	}

	log.Printf("%s\x1b[32m%s\x1b[0m%s", "------------------------- ", "Transaction Found", " -------------------------")
	mouseTx, err = calldec.DecodeCalldata(mouseTx, tx.Data())
	if err != nil {
		return nil, err
	}

	// ---------------------------------------------------------- //
	//							  Output						  //
	// ---------------------------------------------------------- //
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

		if len(mouseTx.PossibleSignatures[0].Arguments.TextArguments) > 1 {
			log.Printf("\x1b[32m%s\x1b[0m%v", "Arguments: ", mouseTx.PossibleSignatures[0].Arguments.TextArguments)

			if len(mouseTx.PossibleSignatures[0].Arguments.TextArguments) == len(mouseTx.Calldata) {
				for i, d := range mouseTx.Calldata {
					log.Printf("\x1b[32m%s\x1b[0m%v\n", fmt.Sprintf("\tInput Data (index %v) is of type %v: ", i, mouseTx.PossibleSignatures[0].Arguments.TextArguments[i]), d)
				}
			} else {
				for i, d := range mouseTx.Calldata {
					log.Printf("\x1b[32m%s\x1b[0m%v\n", fmt.Sprintf("\tInput Data (index %v): ", i), d)
				}
			}

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

		for _, sig := range mouseTx.PossibleSignatures {
			dt, err := hexutil.Decode(sig.HexSignature)
			if err != nil {
				// log.Printf("\x1b[31m%s\x1b[0m%v", "failed to decode hex: ", err.Error())
				continue
			}

			_, err = ab.MethodById([]byte(dt))
			if err != nil {
				// log.Printf("\x1b[31m%s\x1b[0m%v", "failed to get method by id: ", err.Error())
				continue
			}
		}
	}

	log.Printf("---------------------------------------------------------------------\n")

	return mouseTx, nil
}
