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
	dnsCmd.AddCommand(addCmd)
	dnsCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(dnsCmd)
}

var dnsCmd = &cobra.Command{
	Use: "dns",
	Short: "Control dns-server list of the server",
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

		resp := proto.DnsServers{}
		err = c.Call("ListDnsServers", &empty.Empty{}, &resp)
		if err != nil {
			log.Fatalf("request failed: %v", err)
		}
		if len(resp.List) == 0 {
			fmt.Println("No DNS servers")
			return
		}
		for _, server := range resp.List {
			fmt.Println(server.Ip)
		}
	},
}

// TODO: Parse IP to validate it (net.ParseIP)
var addCmd = &cobra.Command {
	Use: "add [dns-ips...]",
	Short: "Add dns servers remotely",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toAdd := make([]*proto.DnsServer, 0, len(args))
		for _, ip := range args {
			toAdd = append(toAdd, &proto.DnsServer{ Ip: ip })
		}

		addr, _ := cmd.Flags().GetString("addr")
		timeout, _ := cmd.Flags().GetDuration("timeout")
		rest, _ := cmd.Flags().GetBool("rest")

		c, err := service.NewCaller(addr, timeout, rest)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Cleanup()

		resp := proto.DnsServers{}
		err = c.Call("AddDnsServers", &proto.DnsServers{
			List: toAdd,
		}, &resp)
		if err != nil {
			log.Fatalf("request failed: %v", err)
		}
		if len(resp.List) == 0 {
			fmt.Println("No DNS servers")
			return
		}
		for _, server := range resp.List {
			fmt.Println(server.Ip)
		}
	},
}

var removeCmd = &cobra.Command {
	Use: "remove [dns-ips...]",
	Short: "Remove dns servers remotely",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toRemove := make([]*proto.DnsServer, 0, len(args))
		for _, ip := range args {
			toRemove = append(toRemove, &proto.DnsServer{ Ip: ip })
		}

		addr, _ := cmd.Flags().GetString("addr")
		timeout, _ := cmd.Flags().GetDuration("timeout")
		rest, _ := cmd.Flags().GetBool("rest")

		c, err := service.NewCaller(addr, timeout, rest)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Cleanup()

		resp := proto.DnsServers{}
		err = c.Call("RemoveDnsServers", &proto.DnsServers{
			List: toRemove,
		}, &resp)
		if err != nil {
			log.Fatalf("request failed: %v", err)
		}
		if len(resp.List) == 0 {
			fmt.Println("No DNS servers")
			return
		}
		for _, server := range resp.List {
			fmt.Println(server.Ip)
		}
	},
}