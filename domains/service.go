package domains

import (
	"context"

	"google.golang.org/grpc"
)

func ServiceConn(addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(context.Background(), addr, grpc.WithInsecure(), grpc.WithBlock())
}
