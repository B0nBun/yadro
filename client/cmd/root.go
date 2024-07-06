package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "client",
	Short: "control the hostname and dns servers remotely",
}

// TODO: request timeout in cli args
// TODO: Add support for client REST http request -- `useRest, _ := cmd.PersistentFlags().GetBool("rest")`
func init() {
	rootCmd.PersistentFlags().BoolP("rest", "R", false, "set to true if the client should send REST http request (otherwise it uses grpc)")
	rootCmd.PersistentFlags().StringP("addr", "a", "0.0.0.0:1235", "server address")
}

func Execute() error {
	return rootCmd.Execute()
}