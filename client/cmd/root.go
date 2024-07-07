package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "control the hostname and dns servers remotely",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("rest", "R", false, "if set to true, client will use REST http request instead of grpc")
	rootCmd.PersistentFlags().StringP("addr", "a", "", "server address")
	rootCmd.MarkPersistentFlagRequired("addr")
	rootCmd.PersistentFlags().DurationP("timeout", "t", time.Second, "request timeout")
}

// We could use Command.PreRun with Command.ExecuteContext(context.Context) to create service.Caller only once, here,
// but I don't like to diffuse the logic across commands, so I just use the same code in every command (CallerFromFlagSet(...) from "utils")
func Execute() error {
	return rootCmd.Execute()
}
