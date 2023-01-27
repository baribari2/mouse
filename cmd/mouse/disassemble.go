package mouse

import (
	"log"
	"os/exec"

	"github.com/baribari2/mouse/decomp"
	"github.com/spf13/cobra"
)

var (
	code   string
	out    string
	source string
)

// TODO: Sanity checks
// TODO: Save to file
var disassembleCmd = &cobra.Command{
	Use:     "disassemble",
	Short:   "Disassemble a contract",
	Long:    `Disassemble a contract`,
	Aliases: []string{"ds"},
	Run: func(cmd *cobra.Command, args []string) {

		if code[0:2] == "0x" {
			code = code[2:]
		}

		m, err := decomp.DisassembleBytecode(code)
		if err != nil {
			log.Fatal(err)
		}

		c := exec.Command("clear")

		c.Start()
		if err := c.Wait(); err != nil {
			log.Fatal(err)
		}

		for i, op := range m.Target.Opcodes {
			if i > 650 {
				break
			}

			log.Printf("\x1b[33m%v \x1b[32m%s\x1b[0m\t %s", i, op.Instruction, op.Data)
		}

	},
}

func init() {
	disassembleCmd.Flags().StringVarP(&code, "code", "c", "", "Contract bytecode")
	disassembleCmd.Flags().StringVarP(&out, "out", "o", "", "Output file path (default: stdout)")
	disassembleCmd.Flags().StringVarP(&source, "source", "s", "", "Source file path (default: root directory)")

	rootCmd.AddCommand(disassembleCmd)
}
