package accounts

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/user_pb"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AccountServer struct {
	//profile_pb.UnimplementedProfileServiceServer
	user_pb.UnimplementedUserServiceServer
}

type User struct {
	id        string
	username  string
	password  string
	createdAt string
	updatedAt string
}

// Dummy db
var users []*user_pb.CreateUserRequest

func (*AccountServer) CreateUser(ctx context.Context, req *user_pb.CreateUserRequest) (*user_pb.CreateUserResponse, error) {
	log.Println("CreateUser called")
	log.Println(req.GetUser())

	users = append(users, req)

	return &user_pb.CreateUserResponse{UserId: "1"}, nil
}

func (*AccountServer) GetUser(ctx context.Context, req *user_pb.GetUserRequest) (*user_pb.GetUserResponse, error) {

	return nil, nil
}

func (*AccountServer) UpdateUser(ctx context.Context, req *user_pb.UpdateUserRequest) (*user_pb.UpdateUserResponse, error) {
	return nil, nil
}

func (*AccountServer) DeleteUser(ctx context.Context, req *user_pb.DeleteUserRequest) (*user_pb.DeleteUserResponse, error) {
	return nil, nil
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatalf("Unable to hash password: %v", err)
	}
	return string(hash)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString("SECRET_KEY")
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}
