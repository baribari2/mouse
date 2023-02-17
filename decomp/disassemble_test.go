package decomp

import (
	"testing"
)

func TestDissassembleBytecode(t *testing.T) {
	bc := "0x60806040"

	tx, err := DisassembleBytecode(bc)
	if err != nil {
		t.Fatal(err)
	}

	if tx.Target.Opcodes[0].Instruction() != "PUSH1" {
		t.Fatal("PUSH1 not found")
	}

	if tx.Target.Opcodes[1].Instruction() != "PUSH1" {
		t.Fatal("PUSH1 not found")
	}
}
