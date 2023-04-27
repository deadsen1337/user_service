package snowflake

import (
	"google.golang.org/grpc"
	"log"
)

func NewConnection() *grpc.ClientConn {
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}

	conn, err := grpc.Dial("localhost:7002", opts...)
	if err != nil {
		log.Fatal("grpc connect ffs error: ", err)
	}

	return conn
}
