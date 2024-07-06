package cmd

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yadro-dns/gen/go/proto"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "server",
	Short: "start grpc and http servers for the dns-service",
	Run: runClient, // TODO: use Command.RunE instead
}

func init() {
	rootCmd.PersistentFlags().BoolP("rest", "R", false, "set to true if the client should send REST http request (otherwise it uses grpc)")
	rootCmd.PersistentFlags().StringP("addr", "a", "0.0.0.0:1235", "server address")
}

func runClient(cmd *cobra.Command, _args []string) {
	// TODO: Add support for client REST http request -- `useRest, _ := cmd.PersistentFlags().GetBool("rest")`
	addr, _ := cmd.PersistentFlags().GetString("addr")

	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewDnsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetHostname(ctx, &proto.GetHostnameParams{})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Recieved back: %s", r.GetName())
}

func Execute() error {
	return rootCmd.Execute()
}