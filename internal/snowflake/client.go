package snowflake

import (
	"google.golang.org/grpc"
	snowflake_api_service "user_service/pkg/snowflake-api"
)

func NewClient(conn *grpc.ClientConn) snowflake_api_service.SnowflakeServiceClient {
	client := snowflake_api_service.NewSnowflakeServiceClient(conn)
	return client
}
