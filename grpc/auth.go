package grpc

import (
	"context"

	"github.com/Penetration-Platform-Go/Auth-Service/lib"
	auth "github.com/Penetration-Platform-Go/gRPC-Files/Auth-Service"
)

// AuthService implement auth grpc service
type AuthService struct {
}

// GetUsernameByToken returns username
func (u *AuthService) GetUsernameByToken(ctx context.Context, in *auth.Token) (*auth.Username, error) {
	username, err := lib.CheckJWT(in.JWT)
	return &auth.Username{
		Username: username,
	}, err
}

// GetToken return token
func (u *AuthService) GetToken(ctx context.Context, in *auth.Username) (*auth.Token, error) {
	token, err := lib.GenerateJWT(in.Username)
	return &auth.Token{
		JWT: token,
	}, err
}
