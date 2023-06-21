package auth

import (
	"context"
	"fmt"
	"go-api-gateway/config"
	"go-api-gateway/domains"
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

// ////////////////// signup end ///////////////////////
// ////////////////// Check Token //////////////////////////
func CheckToken(u *auth.CheckTokenRequest) (*auth.CheckTokenResponse, error) {
	conn, err := domains.ServiceConn(config.AuthServiceURI())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	//////// connect with client //////////
	client := auth.NewAuthClient(conn)
	//////// connect with client //////////

	return client.CheckToken(context.Background(), u)
}
func CreateTokenRequest(query []byte) (*auth.CheckTokenRequest, error) {
	u := auth.CheckTokenRequest{}
	return &u, protojson.Unmarshal(query, &u)
}
