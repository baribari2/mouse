package types

type MouseContract struct {
	Address   string
	Bytecode  string
	Storage   map[string]interface{}
	Functions []string
	Opcodes   []*Opcode
}
