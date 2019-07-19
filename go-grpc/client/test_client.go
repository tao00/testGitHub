package main

import (
	"log"
	"os"

	pb "github.com/tao00/testGitHub/go-grpc/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewActionClient(conn)

	name := "lin"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.GrpcTest(context.Background(), &pb.Request{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Message)
}
