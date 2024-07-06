package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/golang/protobuf/ptypes/empty"
	"yadro/client/service"
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
		addr, _ := cmd.Flags().GetString("addr")
		timeout, _ := cmd.Flags().GetDuration("timeout")
		rest, _ := cmd.Flags().GetBool("rest")

		c, err := service.NewCaller(addr, timeout, rest)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Cleanup()

		resp := proto.Hostname{}
		err = c.Call("GetHostname", &empty.Empty{}, &resp)
		if err != nil {
			log.Fatalf("request failed: %v", err)
		}
		fmt.Println(resp.GetName())
	},
}

var setCmd = &cobra.Command{
	Use: "set [hostname]",
	Short: "Set the hostname on the server",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr, _ := cmd.Flags().GetString("addr")
		timeout, _ := cmd.Flags().GetDuration("timeout")
		rest, _ := cmd.Flags().GetBool("rest")

		c, err := service.NewCaller(addr, timeout, rest)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Cleanup()

		hostname := args[0]
		resp := proto.Hostname{}
		err = c.Call("SetHostname", &proto.Hostname{Name: hostname}, &resp)
		if err != nil {
			log.Fatalf("rpc failed: %v", err)
		}
		fmt.Println(resp.GetName())
	},
}