# Mouse

Mouse is a tool for debugging data on Ethereum. Currently, one can decode ABI-encoded calldata, or disassemble smart contracts to EVM instructions.

## Installation
Download the file corresponding to your OS in the releases page

## Usage
```
mouse is a tool for analyzing Ethereum transactions

Usage:
  mouse [flags]
  mouse [command]

Available Commands:
  decode      Decode a transaction
  decompile   Decompile a contract
  disassemble Disassemble a contract
  help        Help about any command

Flags:
  -h, --help   help for mouse
```

#### Example
The command `$ mouse decode -o ./decode-test.txt -d 0x095ea7b3000000000...` will output the data below in a file named `decode-test.txt`
```
Decoded function signature: approve(address,uint256)
  Arg 0 (type): address
  Arg 1 (type): uint256
```

The command `$ mouse disassemble -s ./uniswap.hex -o ./uniswap.dis` will fetch bytecode from a file named `uniswap.hex` and output the data below in a file named `uniswap.dis`
```
0 CALLVALUE  
1 PUSH3 2fffff 
2 JUMPI  
3 PUSH1 04 
4 PUSH1 04 
5 PUSH1 04 
6 SHL  
7 SUB  
8 PUSH3 2fffff 
9 CODESIZE  
10 DUP2  
11 SWAP1  
12 SUB
// ...[snip]...
```

### Coming Soon
- [ ] Decompilation
- [ ] Prettier output