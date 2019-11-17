package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	pb "github.com/newbmiao/go-staffs/grpc/helloworld"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

//server side:
// check token scope
func tokenVerifyFunc(ctx context.Context) (context.Context, error) {
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil || tokenString == "" {
		log.Println("parse token err:", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")

	}
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, nil)

	claims, _ := token.Claims.(*CustomClaims)

	hasScope := false
	result := strings.Split(claims.Scope, " ")
	for i := range result {
		//check scope "read:users" ; if not contains this, will failed pass
		if result[i] == "read:users" {
			hasScope = true
		}
	}
	if hasScope {
		return ctx, nil
	}
	//mark some ctx field...

	return nil, status.Error(codes.PermissionDenied, "invalid scope")
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//add jwt token interceptor
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_auth.UnaryServerInterceptor(tokenVerifyFunc),
	)))
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
