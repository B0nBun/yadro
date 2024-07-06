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
	hostnameCmd.Flags().StringP("set", "s", "", "set new hostname on the server")
	rootCmd.AddCommand(hostnameCmd)
}

var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Control hostname of the server",
	Run: func(cmd *cobra.Command, args []string) {
		addr, _ := cmd.Flags().GetString("addr")
		timeout, _ := cmd.Flags().GetDuration("timeout")
		rest, _ := cmd.Flags().GetBool("rest")
		setFlag := cmd.Flags().Lookup("set")

		c, err := service.NewCaller(addr, timeout, rest)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Cleanup()

		if setFlag.Changed {
			hostname, _ := cmd.Flags().GetString("set")
			resp := proto.Hostname{}
			err := c.Call("SetHostname", &proto.Hostname{Name: hostname}, &resp)
			if err != nil {
				log.Fatalf("rpc failed: %v", err)
			}
			fmt.Println(resp.GetName())
		} else {
			resp := proto.Hostname{}
			err := c.Call("GetHostname", &empty.Empty{}, &resp)
			if err != nil {
				log.Fatalf("request failed: %v", err)
			}
			fmt.Println(resp.GetName())
		}
	},
}
