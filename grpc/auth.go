package grpc

import (
	"context"

	"github.com/Penetration-Platform-Go/Auth-Service/lib"
	auth "github.com/Penetration-Platform-Go/gRPC-Files/Auth-Service"
)

// AuthService implement auth grpc service
type AuthService struct {
}

// GetJWTTokenKey returns token key
func (u *AuthService) GetJWTTokenKey(ctx context.Context, in *auth.Empty) (*auth.JWTTokenKey, error) {
	return &auth.JWTTokenKey{
		JWTTokenKey: lib.JWTKey,
	}, nil
}

// GetToken return token
func (u *AuthService) GetToken(ctx context.Context, in *auth.TokenMessages) (*auth.Token, error) {
	token, err := lib.GenerateJWT(in.Username, in.IsValid)
	return &auth.Token{
		JWT: token,
	}, err
}
