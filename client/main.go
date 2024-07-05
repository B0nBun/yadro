package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "yadro-dns/gen/go/proto"
)

func main() {
	addr := flag.String("addr", "localhost:1234", "the address to connect to")
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDnsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.StringMessage{Value: "Echo this!"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Recieved back: %s", r.GetValue())
}
