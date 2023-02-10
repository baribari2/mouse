package mouse

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/baribari2/mouse/decomp"
	"github.com/spf13/cobra"
)

var (
	code   string
	out    string
	source string
)

var disassembleCmd = &cobra.Command{
	Use:     "disassemble",
	Short:   "Disassemble a contract",
	Long:    `Disassemble a contract`,
	Aliases: []string{"ds"},
	Run: func(cmd *cobra.Command, args []string) {
		if code != "" {
			if code[:2] == "0x" {
				code = code[2:]
			}

		} else if source != "" {
			d, err := os.ReadFile(source)
			if err != nil {
				log.Fatalf("error reading file: %v", err)
				return
			}

			code = string(d)
		} else {
			log.Fatal("No bytecode or source file provided")
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

		if out != "" {
			f, err := os.Create(out)
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			for i, op := range m.Target.Opcodes {

				if i > 100 {
					break
				}

				_, err := f.WriteString(fmt.Sprintf("%d %s %s \n", i, op.Instruction, op.Data))
				if err != nil {
					log.Fatal(err)
				}

			}
		} else {
			for i, op := range m.Target.Opcodes {
				log.Printf("\x1b[33m%v \x1b[32m%s\x1b[0m\t %s", i, op.Instruction, op.Data)
			}
		}
	},
}

func init() {
	disassembleCmd.Flags().StringVarP(&code, "code", "c", "", "Contract bytecode")
	disassembleCmd.Flags().StringVarP(&out, "out", "o", "", "Output file path (default: stdout)")
	disassembleCmd.Flags().StringVarP(&source, "source", "s", "", "Source file path (default: root directory)")

	rootCmd.AddCommand(disassembleCmd)
}
