package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"yadro-dns/gen/go/proto"
	"yadro-dns/client/service"
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
		setFlag := cmd.Flags().Lookup("set")

		if setFlag.Changed {
			hostname, _ := cmd.Flags().GetString("set")
			resp := proto.Hostname{}
			err := service.CallGRPC(addr, "SetHostname", timeout, &proto.Hostname{Name: hostname}, &resp)
			if err != nil {
				log.Fatal("rpc failed: %v", err)
			}
			fmt.Println(resp.GetName())
		} else {
			resp := proto.Hostname{}
			err := service.CallGRPC(addr, "GetHostname", timeout, &proto.GetHostnameParams{}, &resp)
			if err != nil {
				log.Fatal("rpc failed: %v", err)
			}
			fmt.Println(resp.GetName())
		}
	},
}
