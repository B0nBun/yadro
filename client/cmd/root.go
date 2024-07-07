package cmd

import (
	"time"

	"github.com/spf13/cobra"
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

// For each command we need a service.Caller, so we could create one here in rootCmd.PreRun, however, I don't like to
// diffuse the logic across commands, so I think repeating the same CallerFromFlagSet(...) at the start of each Run is better
func Execute() error {
	return rootCmd.Execute()
}
