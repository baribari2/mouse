package decomp

type Contract struct {
	Address   string
	Bytecode  string
	Storage   map[string]interface{}
	Functions []string
}

type Decompiler struct{}

// 6080604052
func New() *Decompiler {
	return &Decompiler{}
}

func (d *Decompiler) AnalyzeContract(c *Contract) error {

	return nil
}
