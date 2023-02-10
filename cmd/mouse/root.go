package mouse

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "mouse",
		Short: "mouse is a tool for analyzing ethereum transactions",
		Long:  "mouse is a tool for analyzing ethereum transactions",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	RPC_ENDPOINT string
	ADDRESS      string
)

func Execute(rpc, address string) {
	RPC_ENDPOINT = rpc
	ADDRESS = address

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
