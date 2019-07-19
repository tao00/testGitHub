package main

import (
	"log"
	"net"

	pb "github.com/tao00/testGitHub/go-grpc/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	PORT = ":50001"
)

type server struct{}

func (s *server) GrpcTest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Println("request: ", in.Name)
	return &pb.Response{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterActionServer(s, &server{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)

	// go func() {
	// 	err = s.Serve(lis)
	// 	if err != nil {
	// 		log.Println("Fail to start gRPC Server")
	// 	} else {
	// 		log.Println("rpc服务已经开启")
	// 	}
	// }()

	// conn, err := grpc.DialContext(context.Background(), "localhost:50001", grpc.WithInsecure(), grpc.WithBlock())

	// if err != nil {
	// 	log.Println("Failed to dial server:", err)
	// }
	// mux := http.NewServeMux()

	// jsonpb := &gateway.JSONPb{
	// 	EmitDefaults: true,
	// 	Indent:       "  ",
	// 	OrigName:     true,
	// }
	// gwmux := runtime.NewServeMux(
	// 	runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
	// 	// This is necessary to get error details properly
	// 	// marshalled in unary requests.
	// 	runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	// )

	// err = pb.RegisterActionHandler(context.Background(), gwmux, conn)
	// if err != nil {
	// 	log.Println("Failed to register gateway:", err)
	// }

	// mux.Handle("/", gwmux)

	// gatewayAddr := fmt.Sprintf("%s:%d", "10.15.101.25", 8080)
	// log.Println(fmt.Sprintf("Serving gRPC-Gateway on https://%s", gatewayAddr))

	// gwServer := http.Server{
	// 	Addr:    gatewayAddr,
	// 	Handler: mux,
	// }

	// go gwServer.ListenAndServe()
	// log.Println("RpcDispatch inited")
}
