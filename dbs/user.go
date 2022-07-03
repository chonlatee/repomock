package dbs

import (
	"context"

	"github.com/chonlatee/repomock/models"
)

type User struct {
}

func (u *User) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	return nil, nil
}

func NewUserDB() *User {
	return &User{}
}
