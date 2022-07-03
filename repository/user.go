package repository

//go:generate mockgen -source ./user.go -destination ./mock/user_mock.go

import (
	"context"

	"github.com/chonlatee/repomock/models"
)

type Users interface {
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
}
