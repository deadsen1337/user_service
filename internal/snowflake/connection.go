package snowflake

import (
	"google.golang.org/grpc"
)

func NewConnection() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}

	conn, err := grpc.Dial("localhost:7002", opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
