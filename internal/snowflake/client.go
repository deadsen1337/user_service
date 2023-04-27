package snowflake

import (
	"context"
	snowflake_api "github.com/deadsen1337/snowflake_api"
	"google.golang.org/grpc"
)

type Client struct {
	Snowflake snowflake_api.SnowflakeServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	client := snowflake_api.NewSnowflakeServiceClient(conn)
	return &Client{Snowflake: client}
}

func (c *Client) NewID() (uint64, error) {
	resp, err := c.Snowflake.NewID(context.Background(), &snowflake_api.IdRequest{}, []grpc.CallOption{}...)
	if err != nil {
		return 0, err
	}

	return resp.GetId(), err
}
