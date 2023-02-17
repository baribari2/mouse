package types

type MouseContract struct {
	// Address of the contract
	address string

	// Bytecode of the contract
	bytecode string

	storage map[string]interface{}

	// Function names of the contract
	functions []string

	// Opcodes of the contract
	Opcodes []*Opcode
}

func NewMouseContract(address string, bytecode string, storage map[string]interface{}, functions []string, opcodes []*Opcode) *MouseContract {
	return &MouseContract{
		address:   address,
		bytecode:  bytecode,
		storage:   storage,
		functions: functions,
		Opcodes:   opcodes,
	}
}

func (c *MouseContract) Address() string {
	return c.address
}

func (c *MouseContract) Bytecode() string {
	return c.bytecode
}

func (c *MouseContract) Storage() map[string]interface{} {
	return c.storage
}

func (c *MouseContract) Functions() []string {
	return c.functions
}
