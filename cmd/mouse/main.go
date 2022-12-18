package main

import (
	"context"
	"log"
	_ "net/http/pprof"

	// "mouse/decomp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var SPAM_SIGNATURES = []string{
	"join_tg_invmru_haha",
	"CheckOutBoringSecurity",
	"niceFunctionHerePlzClick",
	"watch_tg_invmru",
	"func_2093253501",
	"many_msg_babbage",
	"sign_szabo_bytecode",
	"transfer_attention_tg",
	"JunionYoutubeXD_dashhvetozhe",
	"cryethereum_",
}

// TODO: Add support for multiple txpools
// TODO: Concurrently analyze txs
// TODO: Add support for contract deployment
// TODO: Add support for contract calls
func main() {
	rc, err := rpc.Dial(RPC_ENDPOINT)
	if err != nil {
		log.Fatalf("failed to dial rpc client: %v", err.Error())
	}

	ec, err := ethclient.Dial(RPC_ENDPOINT)
	if err != nil {
		log.Fatalf("failed to dial ethclient: %v", err.Error())
	}

	log.Printf("\x1b[32m%s\x1b[0m", "Starting mempool analyzer...")

	gc := gethclient.New(rc)

	ch := make(chan common.Hash)
	_, err = gc.SubscribePendingTransactions(context.Background(), ch)
	if err != nil {
		log.Fatalf("failed to sub to pending transactions: %v", err.Error())
		return
	}

	// check for a lot of leading 0's (filter out seaport)
	for h := range ch {
		tx, _, err := ec.TransactionByHash(context.Background(), h)
		if err != nil {
			continue
		}

		// goroutine to analyze contract
		if tx.To() == nil {
			log.Printf("\x1b[31m%s\x1b[0m%v", "Contract deployment! ", tx.Hash().String())
			log.Printf("\x1b[31m%s\x1b[0m", "TODO: Analyze contract")
			continue
		}

		if len(tx.Data()) > 3 {
			_, err = AnalyzeTx(ec, tx)
			if err != nil {
				log.Printf("\x1b[31m%s\x1b[0m%v", "failed to analyze tx: ", err.Error())
				continue
			}
		}
	}
}

// If not array or struct, and has more than the amount of parameters in the calldata, its invalid
