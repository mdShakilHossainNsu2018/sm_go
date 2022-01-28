package accounts

import (
	"context"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/user_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(bufSize)

	server := grpc.NewServer()

	user_pb.RegisterUserServiceServer(server, &AccountServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestCreateAccount(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialer()), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)
	client := user_pb.NewUserServiceClient(conn)
	resp, err := client.CreateUser(ctx, &user_pb.CreateUserRequest{
		User: &user_pb.User{
			UserId:   "1",
			Username: "Test",
			Password: "tom2Jery",
		}})
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}
