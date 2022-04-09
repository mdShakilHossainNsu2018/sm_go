package main

import (
	"database/sql"
	"fmt"
	"github.com/mdShakilHossainNsu2018/sm_go/accounts"
	"github.com/mdShakilHossainNsu2018/sm_go/databases"
	"github.com/mdShakilHossainNsu2018/sm_go/interceptors"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/auth_pb"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/user_pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
)

var db *sql.DB

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func main() {
	log.Println("Main Program started")

	fmt.Println("Connected!")

	albums, err := databases.AlbumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userStore := accounts.NewInMemoryUserStore()
	jwtManager := accounts.NewJWTManager(secretKey, tokenDuration)
	authService := accounts.NewAuthServer(userStore, jwtManager)
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryInterceptor),
		grpc.StreamInterceptor(interceptors.StreamInterceptor))

	user_pb.RegisterUserServiceServer(s, &accounts.AccountServer{})
	auth_pb.RegisterAuthServiceServer(s, authService)
	log.Printf("Starting gRPC listener on port" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
