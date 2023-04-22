package snowflake

import (
	"google.golang.org/grpc"
	"log"
)

func NewConnection() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:7002", opts...)
	if err != nil {
		log.Fatal("grpc connect ffs error: ", err)
	}

	return conn
}
