package main

import (
	"github.com/mdShakilHossainNsu2018/sm_go/accounts"
	"github.com/mdShakilHossainNsu2018/sm_go/interceptors"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/user_pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	log.Println("Main Program started")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryInterceptor),
		grpc.StreamInterceptor(interceptors.StreamInterceptor))

	user_pb.RegisterUserServiceServer(s, &accounts.AccountServer{})
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
