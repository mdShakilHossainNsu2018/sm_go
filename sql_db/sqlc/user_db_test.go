package sql_db

import (
	"context"
	"database/sql"
	"github.com/mdShakilHossainNsu2018/sm_go/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueries_CreateUserUserDB(t *testing.T) {
	assertions := assert.New(t)
	arg := CreateUserParams{Username: sql.NullString{String: utils.RandomString(6), Valid: true},
		Password: sql.NullString{String: utils.RandomString(8), Valid: true}}

	user, err := testQueries.CreateUser(context.Background(), arg)

	assertions.Nil(err)
	assertions.Equal(arg.Username, user.Username)
	assertions.Equal(arg.Password, user.Password)
	assertions.NotZero(user.CreatedAt)
}

func TestQueries_ListUsers(t *testing.T) {
	assertions := assert.New(t)

	prams := ListUsersParams{Limit: 10, Offset: 1}

	users, err := testQueries.ListUsers(context.Background(), prams)
	assertions.Nil(err)
	assertions.NotZero(users)

}
