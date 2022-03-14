package accounts

import (
	"fmt"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/user_pb"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string
	Username  string
	Password  string
	Role      string
	CreatedAt string
	UpdatedAt string
}

// NewUser returns a new user
func NewUser(username string, password string, role string) (*user_pb.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &user_pb.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}
	return user, nil
}

// IsCorrectPassword checks if the provided password is correct or not
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// Clone returns a clone of this user
func (user *User) Clone() *User {
	return &User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
}
