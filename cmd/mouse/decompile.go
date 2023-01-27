package mouse

import (
	"github.com/spf13/cobra"
)

var decompileCmd = &cobra.Command{
	Use:     "decompile",
	Short:   "Decompile a contract",
	Long:    `Decompile a contract`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(decompileCmd)
}
