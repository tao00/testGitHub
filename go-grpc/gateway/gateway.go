package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gogo/gateway"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/tao00/testGitHub/go-grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// var (
// 	echoEndpoint = flag.String("echo_endpoint", "localhost:50001", "endpoint of Gateway")
// )

// func run() error {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	mux := runtime.NewServeMux()
// 	opts := []grpc.DialOption{grpc.WithInsecure()}
// 	err := gw.RegisterActionHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("服务开启")
// 	return http.ListenAndServe(":8080", mux)
// }

// func main() {
// 	flag.Parse()
// 	defer glog.Flush()

// 	if err := run(); err != nil {
// 		glog.Fatal(err)
// 	}
// }

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:50001", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Println("Failed to dial server:", err)
	}
	mux := http.NewServeMux()

	jsonpb := &gateway.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		// This is necessary to get error details properly
		// marshalled in unary requests.
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	err = pb.RegisterActionHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Println("Failed to register gateway:", err)
	}

	mux.Handle("/", gwmux)

	gatewayAddr := fmt.Sprintf("%s:%d", "localhost", 8181)
	log.Println(fmt.Sprintf("Serving gRPC-Gateway on https://%s", gatewayAddr))

	// gwServer := http.Server{
	// 	Addr:    gatewayAddr,
	// 	Handler: mux,
	// }

	// gwServer.ListenAndServe()
	log.Println("RpcDispatch inited")

	flag.Parse()
	defer glog.Flush()

	if err := http.ListenAndServe(gatewayAddr, mux); err != nil {
		glog.Fatal(err)
	}
}
