package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yadro/gen/go/proto"
)

type Caller interface {
	Call(method string, args, resp interface{}) error
	Cleanup()
}

func NewCaller(addr string, timeout time.Duration, REST bool) (Caller, error) {
	if REST {
		return NewRESTCaller(addr, timeout), nil
	} else {
		return NewGRPCCaller(addr, timeout)
	}
}

type GRPCCaller struct {
	conn          *grpc.ClientConn
	serviceClient proto.DnsServiceClient
	timeout       time.Duration
}

func NewGRPCCaller(addr string, timeout time.Duration) (c *GRPCCaller, err error) {
	c = new(GRPCCaller)
	c.timeout = timeout
	c.conn, err = grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		err = fmt.Errorf("failed to connect: %w", err)
		return
	}
	c.serviceClient = proto.NewDnsServiceClient(c.conn)
	return
}

func (c *GRPCCaller) Call(method string, args, resp interface{}) error {
	av := reflect.ValueOf(args)
	if av.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", av.Kind()))
	}
	rv := reflect.ValueOf(resp)
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", rv.Kind()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	methodv := reflect.ValueOf(c.serviceClient).MethodByName(method)
	values := methodv.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		av,
	})
	result := values[0]
	errv := values[1]
	if errv.IsNil() {
		rv.Elem().Set(result.Elem())
		return nil
	} else {
		return errv.Interface().(error)
	}
}

func (c *GRPCCaller) Cleanup() {
	c.conn.Close()
}

type RESTCaller struct {
	httpC http.Client
	addr  string
}

func NewRESTCaller(addr string, timeout time.Duration) *RESTCaller {
	return &RESTCaller{
		httpC: http.Client{Timeout: timeout},
		addr:  addr,
	}
}

func (c *RESTCaller) Call(rpcMethod string, args, resp interface{}) error {
	av := reflect.ValueOf(args)
	if av.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", av.Kind()))
	}
	rv := reflect.ValueOf(resp)
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected pointer, but got %v", rv.Kind()))
	}

	method, relUrl := rpcToRest(rpcMethod)
	apiUrl, err := url.JoinPath("http://", c.addr, relUrl)
	if err != nil {
		return fmt.Errorf("couldn't join urls (%v, %v): %v", c.addr, relUrl, err)
	}
	encoded, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("failed to encode to json: %v", err)
	}

	req, err := http.NewRequest(method, apiUrl, bytes.NewBuffer(encoded))
	r, err := c.httpC.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return fmt.Errorf("failed to decode json: %v", err)
	}
	return nil
}

func rpcToRest(rpcMethod string) (method string, url string) {
	switch rpcMethod {
	case "GetHostname":
		return "GET", "/api/hostname"
	case "SetHostname":
		return "POST", "/api/hostname"
	case "ListDnsServers":
		return "GET", "/api/dns-servers"
	case "AddDnsServers":
		return "POST", "/api/dns-servers"
	case "RemoveDnsServers":
		return "PUT", "/api/dns-servers/delete"
	default:
		panic(fmt.Sprintf("Got unexpected rpc method '%v'", rpcMethod))
	}
}

func (c *RESTCaller) Cleanup() {}
