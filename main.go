package main

import (
    "log"
    "context"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/ethclient/gethclient"
    "github.com/ethereum/go-ethereum/common"
)

func main() {
    rc, err := rpc.Dial("ws://localhost:8546")
    if err != nil {
        log.Fatalf("failed to dial rpc client: %v", err.Error())
    }

    ec, err := ethclient.Dial("ws://localhost:8546")
    if err != nil {
        log.Fatalf("failed to dial ethclient: %v", err.Error())
    }

    log.Printf("\x1b[32m%s\x1b[0m", "Starting mempool analyzer...")

    gc := gethclient.New(rc)

    ch := make(chan common.Hash)
    _, err = gc.SubscribePendingTransactions(context.Background(), ch)
    if err != nil {
        log.Fatalf("failed to sub to pending transactions: %v", err.Error())
    }

    for h := range ch {
        tx, _, err := ec.TransactionByHash(context.Background(), h)
        if err != nil {
            continue
        }

        log.Printf("%s\x1b[32m%s\x1b[0m%s", "---------- ", "Transaction Found", " ----------")
        log.Printf("Hash: %v", h)
        log.Printf("To: %v", tx.To())
        log.Printf("Cost: %v", tx.Cost())
        log.Printf("Gas Limit: %v", tx.Gas())
        log.Printf("Nonce: %v", tx.Nonce())
        log.Printf("Data: %v", tx.Data())

        d, err := decodeData(tx.Data())
        if err != nil {
            log.Printf("\x1b[31m%s\x1b[0m%v", "failed to decode data: ", err.Error())
        }

        log.Printf("Decoded data: %v", d)
    }
}

type Data struct {
    Sig []byte
    Data [][]byte
}

func decodeData(data []byte) (*Data, error) {
    var d *Data
    if len(data) >= 4 {
        d.Sig = data[:4]
    }

    for i := 0; i < len(data) - 1; {
        d.Data = append(d.Data, data[i:i+31])
    }

    return d, nil
}
