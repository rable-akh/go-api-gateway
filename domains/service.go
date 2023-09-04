package domains

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConn(addr string) (*grpc.ClientConn, error) {
	// return grpc.DialContext(context.Background(), addr, grpc.WithInsecure(), grpc.WithBlock())

	/// Dial
	return grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

}
