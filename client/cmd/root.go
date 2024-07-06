package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "control the hostname and dns servers remotely",
}

func init() {
	rootCmd.PersistentFlags().BoolP("rest", "R", false, "set to true if the client should send REST http request (otherwise it uses grpc)")
	rootCmd.PersistentFlags().StringP("addr", "a", "0.0.0.0:1235", "server address")
	rootCmd.PersistentFlags().DurationP("timeout", "t", time.Second, "request timeout")
}

func Execute() error {
	return rootCmd.Execute()
}
