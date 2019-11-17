package main

import (
	"context"
	"github.com/newbmiao/go-staffs/jwk/utils"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"time"

	pb "github.com/newbmiao/go-staffs/grpc/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

//client side
// interceptor
type JWTInterceptor struct {
	Token string
}

func (t *JWTInterceptor) UnaryClientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Add the current bearer token to the metadata and call the RPC
	// command
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "bearer "+t.Token)
	return invoker(ctx, method, req, reply, cc, opts...)
}
func main() {
	//add jwt token with dial.
	var ji = new(JWTInterceptor)
	ji.Token = utils.GetToken()

	conn, err := grpc.Dial(address, grpc.WithUnaryInterceptor(ji.UnaryClientInterceptor),
		grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
