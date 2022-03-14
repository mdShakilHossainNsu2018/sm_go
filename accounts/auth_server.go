package accounts

import (
	"context"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/auth_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type AuthServer struct {
	auth_pb.UnimplementedAuthServiceServer
	userStore  UserStore
	jwtManager *JWTManager
}

// NewAuthServer returns a new auth server
func NewAuthServer(userStore UserStore, jwtManager *JWTManager) auth_pb.AuthServiceServer {
	return &AuthServer{userStore: userStore, jwtManager: jwtManager}
}

// Login is a unary RPC to login user
func (server *AuthServer) Login(ctx context.Context, req *auth_pb.LoginRequest) (*auth_pb.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &auth_pb.LoginResponse{AccessToken: token}
	return res, nil
}
