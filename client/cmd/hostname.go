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

// TODO: A lot of repeating code

var hostnameCmd = &cobra.Command{
	Use: "hostname",
	Short: "Control hostname of the server",
	Long: "",
	Run: func (cmd *cobra.Command, args []string) {
		addr, _ := cmd.Flags().GetString("addr")
		setFlag := cmd.Flags().Lookup("set")
		if setFlag.Changed {
			hostname, _ := cmd.Flags().GetString("set")
			hostname, err := setHostname(addr, hostname)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(hostname)
		} else {
			hostname, err := getHostname(addr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(hostname)
		}
	},
}

func init() {
	hostnameCmd.Flags().StringP("set", "s", "", "set new hostname on the server")
	rootCmd.AddCommand(hostnameCmd)
}

func getHostname(addr string) (string, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return "", fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()
	c := proto.NewDnsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetHostname(ctx, &proto.GetHostnameParams{})
	if err != nil {
		return "", fmt.Errorf("rpc failed: %w", err)
	}
	return r.GetName(), nil
}

func setHostname(addr, hostname string) (string, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return "", fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()
	c := proto.NewDnsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SetHostname(ctx, &proto.Hostname{Name: hostname})
	if err != nil {
		return "", fmt.Errorf("rpc failed: %w", err)
	}
	return r.GetName(), nil
}
