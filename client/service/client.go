package service

import (
	"fmt"
	"context"
	"time"
	"net/http"
	"net/url"
	"reflect"
	"bytes"
	"encoding/json"

	"yadro-dns/gen/go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	Call(method string, args, resp interface{}) error
}

func NewClient(addr string, timeout time.Duration, REST bool) (c Client, err error, cleanup func()) {
	if REST {
		crest := ClientREST{
			HTTP: http.Client{Timeout: timeout},
			Addr: addr,
		}
		c = &crest
		cleanup = func() { crest.HTTP.CloseIdleConnections() }
	} else {
		var conn *grpc.ClientConn
		conn, err = grpc.NewClient(
			addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			err = fmt.Errorf("failed to connect: %w", err)
			return
		}
		cleanup = func() { conn.Close() }
		c = &ClientGRPC{
			ServiceClient: proto.NewDnsServiceClient(conn),
			Timeout: timeout,
		}
	}
	return
}

type ClientGRPC struct {
	ServiceClient proto.DnsServiceClient
	Timeout time.Duration
}

func (c *ClientGRPC) Call(method string, args, resp interface{}) error {
	av := reflect.ValueOf(args)
	if av.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", av.Kind()))
	}
	rv := reflect.ValueOf(resp)
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", rv.Kind()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	methodv := reflect.ValueOf(c.ServiceClient).MethodByName(method)
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

type ClientREST struct {
	HTTP http.Client
	Addr string
}

func (c *ClientREST) Call(rpcMethod string, args, resp interface{}) error {
	av := reflect.ValueOf(args)
	if av.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", av.Kind()))
	}
	rv := reflect.ValueOf(resp)
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", rv.Kind()))
	}

	method, relUrl := rpcToRest(rpcMethod)
	apiUrl, err := url.JoinPath("http://", c.Addr, relUrl)
	if err != nil {
		return fmt.Errorf("couldn't join urls (%v, %v): %v", c.Addr, relUrl, err)
	}
	encoded, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("failed to encode to json: %v", err)
	}

	req, err := http.NewRequest(method, apiUrl, bytes.NewBuffer(encoded))
	r, err := c.HTTP.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return fmt.Errorf("failed to decode json: %v", err)
	}
	return nil
}

func rpcToRest(rpcMethod string) (method string, url string) {
	switch rpcMethod {
	case "GetHostname": return "GET", "/api/hostname"
	case "SetHostname": return "POST", "/api/hostname"
	default: panic(fmt.Sprintf("Got unexpected rpc method '%v'", rpcMethod))
	}
}