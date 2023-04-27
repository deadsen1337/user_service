package user_service

import "user_service/internal/model"

type userRepo interface {
	Delete(id []byte) error
	Get(id []byte) ([]byte, error)
	Update(id []byte, user model.User) error
	Create(id []byte, user model.User) error
}

type snowflakeClient interface {
	NewID() (uint64, error)
}

type UserSevice struct {
	userRepo        userRepo
	snowflakeClient snowflakeClient
}

func NewService(repo userRepo, client snowflakeClient) *UserSevice {
	return &UserSevice{
		userRepo:        repo,
		snowflakeClient: client,
	}
}
