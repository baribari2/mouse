package decomp

import (
	"log"
	"strconv"
	"strings"

	"github.com/baribari2/mouse/common/types"
)

var UNKNOWN = map[string]bool{
	"22": true,
	"0d": true,
	"d4": true,
	"e9": true,
	"fe": true,
	"e5": true,
	"bf": true,
	"dd": true,
	"af": true,
	"cb": true,
	"2d": true,
	"b5": true,
	"c8": true,
	"d3": true,
}

type Decompiler struct{}

func New() *Decompiler {
	return &Decompiler{}
}

func (d *Decompiler) AnalyzeContract(m *types.MouseTx) error {
	err := DisassembleContract(m)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Change param type to interface
func DisassembleContract(m *types.MouseTx) error {
	log.Println("Disassembling contract...")
	var pc int64 = 0

	for pc < int64(len(m.TargetCode)) {
		o := m.TargetCode[pc : pc+2]
		op := types.MatchOpcode(o)
		if op == nil {
			op = &types.Opcode{
				Instruction: "XXXXXXXXXX",
				Number:      "xx",
				Inputs:      0,
				Outputs:     0,
			}
		}

		if strings.Contains(op.String(), "PUSH") {
			var num int
			if len(op.String()[4:]) > 1 {
				num, _ = strconv.Atoi(op.String()[4:])
				if pc+2+int64(num)*2 < int64(len(m.TargetCode)) {
					op.Data = m.TargetCode[pc+2 : pc+2+int64(num)*2]
				}
			} else {
				num, _ = strconv.Atoi(string(op.String()[4]))
				op.Data = m.TargetCode[pc+2 : pc+2+int64(num)*2]
			}

			m.Target.Opcodes = append(m.Target.Opcodes, op)
			pc += 2 + int64(num)*2
			continue
		}

		m.Target.Opcodes = append(m.Target.Opcodes, op)
		pc += 2
	}

	return nil
}

func DisassembleBytecode(code string) (*types.MouseTx, error) {
	log.Println("Disassembling contract...")
	var pc int64 = 0
	var m = &types.MouseTx{
		TargetCode: code,
	}

	for pc < int64(len(m.TargetCode)) {
		o := m.TargetCode[pc : pc+2]
		op := types.MatchOpcode(o)
		if op == nil {
			op = &types.Opcode{
				Instruction: "XXXXXXXXXX",
				Number:      "xx",
				Inputs:      0,
				Outputs:     0,
			}
		}

		if strings.Contains(op.String(), "PUSH") {
			var num int
			if len(op.String()[4:]) > 1 {
				num, _ = strconv.Atoi(op.String()[4:])
				if pc+2+int64(num)*2 < int64(len(m.TargetCode)) {
					op.Data = m.TargetCode[pc+2 : pc+2+int64(num)*2]
				}
			} else {
				num, _ = strconv.Atoi(string(op.String()[4]))
				op.Data = m.TargetCode[pc+1 : pc+1+int64(num)*2]
			}

			m.Target.Opcodes = append(m.Target.Opcodes, op)
			pc += 2 + int64(num)*2
			continue
		}

		m.Target.Opcodes = append(m.Target.Opcodes, op)
		pc += 2
	}

	return m, nil
}
