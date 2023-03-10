package types

var (
	OPCODES = map[string]*Opcode{
		"00": {"STOP", "00", 0, 0, ""},
		"01": {"ADD", "01", 2, 1, ""},
		"02": {"MUL", "02", 2, 1, ""},
		"03": {"SUB", "03", 2, 1, ""},
		"04": {"DIV", "04", 2, 1, ""},
		"05": {"SDIV", "05", 2, 1, ""},
		"06": {"MOD", "06", 2, 1, ""},
		"07": {"SMOD", "07", 2, 1, ""},
		"08": {"ADDMOD", "08", 3, 1, ""},
		"09": {"MULMOD", "09", 3, 1, ""},
		"0a": {"EXP", "0a", 2, 1, ""},
		"0b": {"SIGNEXTEND", "0b", 2, 1, ""},
		"10": {"LT", "10", 2, 1, ""},
		"11": {"GT", "11", 2, 1, ""},
		"12": {"SLT", "12", 2, 1, ""},
		"13": {"SGT", "13", 2, 1, ""},
		"14": {"EQ", "14", 2, 1, ""},
		"15": {"ISZERO", "15", 1, 1, ""},
		"16": {"AND", "16", 2, 1, ""},
		"17": {"OR", "17", 2, 1, ""},
		"18": {"XOR", "18", 2, 1, ""},
		"19": {"NOT", "19", 1, 1, ""},
		"1a": {"BYTE", "1a", 2, 1, ""},
		"1b": {"SHL", "1b", 2, 1, ""},
		"1c": {"SHR", "1c", 2, 1, ""},
		"1d": {"SAR", "1d", 2, 1, ""},
		"20": {"SHA3", "20", 2, 1, ""},
		"30": {"ADDRESS", "30", 0, 1, ""},
		"31": {"BALANCE", "31", 1, 1, ""},
		"32": {"ORIGIN", "32", 0, 1, ""},
		"33": {"CALLER", "33", 0, 1, ""},
		"34": {"CALLVALUE", "34", 0, 1, ""},
		"35": {"CALLDATALOAD", "35", 1, 1, ""},
		"36": {"CALLDATASIZE", "36", 0, 1, ""},
		"37": {"CALLDATACOPY", "37", 3, 0, ""},
		"38": {"CODESIZE", "38", 0, 1, ""},
		"39": {"CODECOPY", "39", 3, 0, ""},
		"3a": {"GASPRICE", "3a", 0, 1, ""},
		"3b": {"EXTCODESIZE", "3b", 1, 1, ""},
		"3c": {"EXTCODECOPY", "3c", 4, 0, ""},
		"3d": {"RETURNDATASIZE", "3d", 0, 1, ""},
		"3e": {"RETURNDATACOPY", "3e", 3, 0, ""},
		"40": {"BLOCKHASH", "40", 1, 1, ""},
		"41": {"COINBASE", "41", 0, 1, ""},
		"42": {"TIMESTAMP", "42", 0, 1, ""},
		"43": {"NUMBER", "43", 0, 1, ""},
		"44": {"DIFFICULTY", "44", 0, 1, ""},
		"45": {"GASLIMIT", "45", 0, 1, ""},
		"46": {"CHAINID", "46", 0, 1, ""},
		"47": {"SELFBALANCE", "47", 0, 1, ""},
		"48": {"BASEFEE", "48", 0, 1, ""},
		"50": {"POP", "50", 1, 0, ""},
		"51": {"MLOAD", "51", 1, 1, ""},
		"52": {"MSTORE", "52", 2, 0, ""},
		"53": {"MSTORE8", "53", 2, 0, ""},
		"54": {"SLOAD", "54", 1, 1, ""},
		"55": {"SSTORE", "55", 2, 0, ""},
		"56": {"JUMP", "56", 1, 0, ""},
		"57": {"JUMPI", "57", 2, 0, ""},
		"58": {"PC", "58", 0, 1, ""},
		"59": {"MSIZE", "59", 0, 1, ""},
		"5a": {"GAS", "5a", 0, 1, ""},
		"5b": {"JUMPDEST", "5b", 0, 0, ""},
		"60": {"PUSH1", "60", 0, 1, ""},
		"61": {"PUSH2", "61", 0, 1, ""},
		"62": {"PUSH3", "62", 0, 1, ""},
		"63": {"PUSH4", "63", 0, 1, ""},
		"64": {"PUSH5", "64", 0, 1, ""},
		"65": {"PUSH6", "65", 0, 1, ""},
		"66": {"PUSH7", "66", 0, 1, ""},
		"67": {"PUSH8", "67", 0, 1, ""},
		"68": {"PUSH9", "68", 0, 1, ""},
		"69": {"PUSH10", "69", 0, 1, ""},
		"6a": {"PUSH11", "6a", 0, 1, ""},
		"6b": {"PUSH12", "6b", 0, 1, ""},
		"6c": {"PUSH13", "6c", 0, 1, ""},
		"6d": {"PUSH14", "6d", 0, 1, ""},
		"6e": {"PUSH15", "6e", 0, 1, ""},
		"6f": {"PUSH16", "6f", 0, 1, ""},
		"70": {"PUSH17", "70", 0, 1, ""},
		"71": {"PUSH18", "71", 0, 1, ""},
		"72": {"PUSH19", "72", 0, 1, ""},
		"73": {"PUSH20", "73", 0, 1, ""},
		"74": {"PUSH21", "74", 0, 1, ""},
		"75": {"PUSH22", "75", 0, 1, ""},
		"76": {"PUSH23", "76", 0, 1, ""},
		"77": {"PUSH24", "77", 0, 1, ""},
		"78": {"PUSH25", "78", 0, 1, ""},
		"79": {"PUSH26", "79", 0, 1, ""},
		"7a": {"PUSH27", "7a", 0, 1, ""},
		"7b": {"PUSH28", "7b", 0, 1, ""},
		"7c": {"PUSH29", "7c", 0, 1, ""},
		"7d": {"PUSH30", "7d", 0, 1, ""},
		"7e": {"PUSH31", "7e", 0, 1, ""},
		"7f": {"PUSH32", "7f", 0, 1, ""},
		"80": {"DUP1", "80", 1, 2, ""},
		"81": {"DUP2", "81", 2, 3, ""},
		"82": {"DUP3", "82", 3, 4, ""},
		"83": {"DUP4", "83", 4, 5, ""},
		"84": {"DUP5", "84", 5, 6, ""},
		"85": {"DUP6", "85", 6, 7, ""},
		"86": {"DUP7", "86", 7, 8, ""},
		"87": {"DUP8", "87", 8, 9, ""},
		"88": {"DUP9", "88", 9, 10, ""},
		"89": {"DUP10", "89", 10, 11, ""},
		"8a": {"DUP11", "8a", 11, 12, ""},
		"8b": {"DUP12", "8b", 12, 13, ""},
		"8c": {"DUP13", "8c", 13, 14, ""},
		"8d": {"DUP14", "8d", 14, 15, ""},
		"8e": {"DUP15", "8e", 15, 16, ""},
		"8f": {"DUP16", "8f", 16, 17, ""},
		"90": {"SWAP1", "90", 2, 2, ""},
		"91": {"SWAP2", "91", 3, 3, ""},
		"92": {"SWAP3", "92", 4, 4, ""},
		"93": {"SWAP4", "93", 5, 5, ""},
		"94": {"SWAP5", "94", 6, 6, ""},
		"95": {"SWAP6", "95", 7, 7, ""},
		"96": {"SWAP7", "96", 8, 8, ""},
		"97": {"SWAP8", "97", 9, 9, ""},
		"98": {"SWAP9", "98", 10, 10, ""},
		"99": {"SWAP10", "99", 11, 11, ""},
		"9a": {"SWAP11", "9a", 12, 12, ""},
		"9b": {"SWAP12", "9b", 13, 13, ""},
		"9c": {"SWAP13", "9c", 14, 14, ""},
		"9d": {"SWAP14", "9d", 15, 15, ""},
		"9e": {"SWAP15", "9e", 16, 16, ""},
		"9f": {"SWAP16", "9f", 17, 17, ""},
		"a0": {"LOG0", "a0", 2, 0, ""},
		"a1": {"LOG1", "a1", 3, 0, ""},
		"a2": {"LOG2", "a2", 4, 0, ""},
		"a3": {"LOG3", "a3", 5, 0, ""},
		"a4": {"LOG4", "a4", 6, 0, ""},
		"f0": {"CREATE", "f0", 3, 1, ""},
		"f1": {"CALL", "f1", 7, 1, ""},
		"f2": {"CALLCODE", "f2", 7, 1, ""},
		"f3": {"RETURN", "f3", 2, 0, ""},
		"f4": {"DELEGATECALL", "f4", 6, 1, ""},
		"f5": {"CALLBLACKBOX", "f5", 7, 1, ""},
		"f6": {"STATICCALL", "f6", 6, 1, ""},
		"fa": {"STATICCALL", "fa", 6, 1, ""},
		"fd": {"REVERT", "fd", 2, 0, ""},
		"fe": {"INVALID", "fe", 0, 0, ""},
		"ff": {"SELFDESTRUCT", "ff", 1, 0, ""},
	}
)

type Opcode struct {
	// The name of the opcode
	instruction string

	// The number of the opcode in the instruction set
	number string

	// The number of inputs the opcode requires
	inputs int64

	// The number of outputs the opcode produces
	outputs int64

	// The data associated with the opcode
	Data string
}

func NewOpcode(instruction string, number string, inputs int64, outputs int64, data string) *Opcode {
	return &Opcode{
		instruction: instruction,
		number:      number,
		inputs:      inputs,
		outputs:     outputs,
		Data:        data,
	}
}

// Returns the name of the opcode
func (o *Opcode) Instruction() string {
	return o.instruction
}

// Returns the number of the opcode in the instruction set
func (o *Opcode) Number() string {
	return o.number
}

// Returns the number of inputs the opcode requires
func (o *Opcode) Inputs() int64 {
	return o.inputs
}

// Returns the number of outputs the opcode produces
func (o *Opcode) Outputs() int64 {
	return o.outputs
}

// Returns the opcode associated with the given name
func MatchOpcode(op string) *Opcode {
	opcode, ok := OPCODES[op]
	if !ok {
		return nil
	}

	return opcode
}
