package service

import (
	"fmt"
	"context"
	"time"
	"reflect"

	"yadro-dns/gen/go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallGRPC(addr, method string, timeout time.Duration, args, resp interface{}) error {
	av := reflect.ValueOf(args)
	if av.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", av.Kind()))
	}
	rv := reflect.ValueOf(resp)
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", rv.Kind()))
	}

	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	c := proto.NewDnsServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	methodv := reflect.ValueOf(c).MethodByName(method)
	values := methodv.Call([]reflect.Value {
		reflect.ValueOf(ctx),
		av,
	})
	result := values[0]
	errv := values[1]
	rv.Elem().Set(result.Elem())
	if errv.IsNil() {
		return nil
	} else {
		return errv.Interface().(error)
	}
}
