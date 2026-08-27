package client

import (
	"context"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
	authv1 "github.com/SamuraiJeka/SnapLink-proto/gen/go/auth/v1"

	"google.golang.org/grpc"
)

type AuthClient struct {
	client authv1.AuthServiceClient
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		client: authv1.NewAuthServiceClient(conn),
	}
}

func (c *AuthClient) Register(
	ctx context.Context,
	req dto.Register,
) (dto.TokenResponse, error) {
	resp, err := c.client.Register(
		ctx,
		toProtoRegister(req),
	)
	if err != nil {
		return dto.TokenResponse{}, err
	}

	return toDtoToken(resp), nil
}

func (c *AuthClient) Login(
	ctx context.Context,
	req dto.Login,
) (dto.TokenResponse, error) {
	resp, err := c.client.Login(
		ctx,
		toProtoLogin(req),
	)
	if err != nil {
		return dto.TokenResponse{}, nil
	}

	return toDtoToken(resp), nil
}

func (c *AuthClient) Refresh(
	ctx context.Context,
	req dto.Refresh,
) (dto.TokenResponse, error) {
	resp, err := c.client.Refresh(
		ctx,
		&authv1.RefreshRequest{
			RefreshToken: req.Refresh_token,
		},
	)
	if err != nil {
		return dto.TokenResponse{}, nil
	}

	return toDtoToken(resp), nil
}

func (c *AuthClient) Logout(
	ctx context.Context,
	req dto.Refresh,
) error {
	_, err := c.client.Logout(
		ctx,
		&authv1.LogoutRequest{
			RefreshToken: req.Refresh_token,
		},
	)
	if err != nil {
		return err
	}
	
	return nil
}

func (c *AuthClient) ValidateAccessToken(
	ctx context.Context,
	req dto.ValidateRequest,
) (dto.ValidateResponse, error) {
	resp, err := c.client.ValideteAccessToken(
		ctx,
		&authv1.ValidateAccessTokenRequest{
			AccessToken: req.Access_token,
		},
	)
	if err != nil {
		return dto.ValidateResponse{}, nil
	}

	return toDtoValidate(resp), nil
}
