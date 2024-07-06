package cmd

import (
	"fmt"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yadro-dns/gen/go/proto"
	"github.com/spf13/cobra"
)


func init() {
	hostnameCmd.Flags().StringP("set", "s", "", "set new hostname on the server")
	rootCmd.AddCommand(hostnameCmd)
}

var hostnameCmd = &cobra.Command{
	Use: "hostname",
	Short: "Control hostname of the server",
	Long: "",
	Run: func (cmd *cobra.Command, args []string) {
		addr, _ := cmd.Flags().GetString("addr")
		setFlag := cmd.Flags().Lookup("set")

		conn, err := grpc.NewClient(
			addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to connect: %v", err)
		}
		defer conn.Close()

		c := proto.NewDnsServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		var r *proto.Hostname

		if setFlag.Changed {
			hostname, _ := cmd.Flags().GetString("set")
			r, err = c.SetHostname(ctx, &proto.Hostname{Name: hostname})
		} else {
			r, err = c.GetHostname(ctx, &proto.GetHostnameParams{})
		}

		if err != nil {
			log.Fatalf("rpc failed: %v", err)
		}
		fmt.Println(r.GetName())
	},
}

