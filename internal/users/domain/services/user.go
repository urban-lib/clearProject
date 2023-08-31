package services

import "context"

type UserService interface {
	FindUserByID(ctx context.Context, userID int)
	FindUserByLogin(ctx context.Context, login string)
}
