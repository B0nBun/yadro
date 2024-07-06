package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yadro-dns/gen/go/proto"
)

func main() {
	addr := flag.String("addr", "localhost:1235", "the address to connect to")
	flag.Parse()

	conn, err := grpc.NewClient(
		*addr,
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
