package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/golang/protobuf/ptypes/empty"
	"yadro/gen/go/proto"
)

func init() {
	hostnameCmd.AddCommand(setCmd)
	rootCmd.AddCommand(hostnameCmd)
}

var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Get hostname of the server",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := CallerFromFlagSet(cmd.Flags())
		if err != nil {
			Fatalf("%v", err)
		}
		defer c.Cleanup()

		resp := proto.Hostname{}
		err = c.Call("GetHostname", &empty.Empty{}, &resp)
		if err != nil {
			Fatalf("request failed: %v", err)
		}
		fmt.Println(resp.GetName())
	},
}

var setCmd = &cobra.Command{
	Use: "set [hostname]",
	Short: "Set the hostname on the server",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := CallerFromFlagSet(cmd.Flags())
		if err != nil {
			Fatalf("%v", err)
		}
		defer c.Cleanup()

		hostname := args[0]
		resp := proto.Hostname{}
		err = c.Call("SetHostname", &proto.Hostname{Name: hostname}, &resp)
		if err != nil {
			Fatalf("rpc failed: %v", err)
		}
		fmt.Println(resp.GetName())
	},
}