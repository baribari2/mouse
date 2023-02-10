package mouse

import (
	"fmt"
	"log"
	"os"

	"github.com/baribari2/mouse/calldec"
	"github.com/baribari2/mouse/common/types"
	"github.com/spf13/cobra"
)

var (
	possible bool
	outp     string
	sourcep  string
	data     string
)

var decodeCmd = &cobra.Command{
	Use:     "decode",
	Short:   "Decode a transaction",
	Long:    `Decode a transaction`,
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		m := &types.MouseTx{}

		if data != "" {
			m.RawCalldata = data
		} else {
			d, err := os.ReadFile(sourcep)
			if err != nil {
				log.Fatalf("error reading file: %v", err)
				return
			}

			m.RawCalldata = string(d)
		}

		err := calldec.DecodeCalldata(m)
		if err != nil {
			log.Printf("error decoding calldata: %v", err)
			return
		}

		if possible {
			log.Printf("Possible signatures:")
			for _, s := range m.PossibleSignatures[1:] {
				log.Printf("  %v", s)
			}
		}

		if outp != "" {
			f, err := os.Create(outp)
			if err != nil {
				log.Fatalf("error creating file: %v", err)
				return
			}

			defer f.Close()

			_, err = f.WriteString(fmt.Sprintf("Decoded function signature: %v\n", m.PossibleSignatures[0].TextSignature))
			for i, a := range m.PossibleSignatures[0].Arguments.TextArguments {
				_, err = f.WriteString(fmt.Sprintf("  Arg %v (type): %v\n", i, a))
			}
		}
	},
}

func init() {
	decodeCmd.Flags().BoolVarP(&possible, "possible", "p", false, "Show possible decoded function signatures")
	decodeCmd.Flags().StringVarP(&outp, "out", "o", "", "Output file path (default: stdout)")
	decodeCmd.Flags().StringVarP(&data, "data", "d", "", "Calldata to decode")
	decodeCmd.Flags().StringVarP(&sourcep, "source", "s", "", "Source file path (default: stdin)")

	rootCmd.AddCommand(decodeCmd)
}
