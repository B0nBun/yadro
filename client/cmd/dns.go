package cmd

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"yadro/gen/go/proto"
)

func init() {
	dnsCmd.AddCommand(addCmd)
	dnsCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(dnsCmd)
}

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Control dns-server list of the server",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := CallerFromFlagSet(cmd.Flags())
		if err != nil {
			Fatalf("%v", err)
		}
		defer c.Cleanup()

		resp := proto.DnsServers{}
		err = c.Call("ListDnsServers", &empty.Empty{}, &resp)
		if err != nil {
			Fatalf("request failed: %v", err)
		}
		if len(resp.List) == 0 {
			fmt.Println("No DNS servers")
			return
		}
		for _, server := range resp.List {
			fmt.Println(server.Address)
		}
	},
}

var addCmd = &cobra.Command{
	Use:   "add [dns-addrs...]",
	Short: "Add dns servers remotely",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toAdd := make([]*proto.DnsServer, 0, len(args))
		for _, addr := range args {
			toAdd = append(toAdd, &proto.DnsServer{Address: addr})
		}

		c, err := CallerFromFlagSet(cmd.Flags())
		if err != nil {
			Fatalf("%v", err)
		}
		defer c.Cleanup()

		resp := proto.DnsServers{}
		err = c.Call("AddDnsServers", &proto.DnsServers{
			List: toAdd,
		}, &resp)
		if err != nil {
			Fatalf("request failed: %v", err)
		}
		if len(resp.List) == 0 {
			fmt.Println("No DNS servers")
			return
		}
		for _, server := range resp.List {
			fmt.Println(server.Address)
		}
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove [dns-addrs...]",
	Short: "Remove dns servers remotely",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toRemove := make([]*proto.DnsServer, 0, len(args))
		for _, addr := range args {
			toRemove = append(toRemove, &proto.DnsServer{Address: addr})
		}

		c, err := CallerFromFlagSet(cmd.Flags())
		if err != nil {
			Fatalf("%v", err)
		}
		defer c.Cleanup()

		resp := proto.DnsServers{}
		err = c.Call("RemoveDnsServers", &proto.DnsServers{
			List: toRemove,
		}, &resp)
		if err != nil {
			Fatalf("request failed: %v", err)
		}
		if len(resp.List) == 0 {
			fmt.Println("No DNS servers")
			return
		}
		for _, server := range resp.List {
			fmt.Println(server.Address)
		}
	},
}
