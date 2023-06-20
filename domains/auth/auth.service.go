package auth

import (
	"context"
	auth "go-api-gateway/protos/go-api-gateway/auth"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func AuthServiceClient(conn *grpc.ClientConn) auth.AuthClient {
	return auth.NewAuthClient(conn)
}

func Login(client auth.AuthClient, u *auth.LoginRequest) (*auth.LoginResponse, error) {
	return client.Login(context.Background(), u)
}

func CreateLoginRequest(jsonQuery []byte) (*auth.LoginRequest, error) {
	u := auth.LoginRequest{}
	// input := []byte(jsonQuery)
	return &u, protojson.Unmarshal(jsonQuery, &u)
}

// ////////////////// Login end ///////////////////////
// ////////////////// signup //////////////////////////
func SignUp(client auth.AuthClient, u *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	return client.SignUp(context.Background(), u)
}
func CreateSignUpRequest(query []byte) (*auth.SignUpRequest, error) {
	u := auth.SignUpRequest{}
	return &u, protojson.Unmarshal(query, &u)
}
