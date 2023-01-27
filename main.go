package main

import (
	cmd "github.com/baribari2/mouse/cmd/mouse"
)

// TODO: Add support for multiple txpools
// TODO: Concurrently analyze txs
// TODO: Add support for contract deployment
// TODO: Add support for contract calls
// TODO: Switch to ethereum logging lib
func main() {
	cmd.Execute(RPC_ENDPOINT, ADD)
}

// If not array or struct, and has more than the amount of parameters in the calldata, its invalid
