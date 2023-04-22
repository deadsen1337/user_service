package user_service

import (
	"context"
	"google.golang.org/grpc"
	snowflake_api_service "user_service/pkg/snowflake-api"
)

type userRepo interface {
	Delete(id []byte) error
	Get(id []byte) ([]byte, error)
	Update(id, user []byte) error
	Create(id, user []byte) error
}

type snowflakeClient interface {
	NewID(ctx context.Context, in *snowflake_api_service.IdRequest, opts ...grpc.CallOption) (*snowflake_api_service.IdResponse, error)
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

