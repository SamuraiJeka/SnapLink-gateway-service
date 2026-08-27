package service

import (
	"context"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
)

type AuthService struct {
	client AuthClient
}

func NewAuthService(client AuthClient) *AuthService {
	return &AuthService{
		client: client,
	}
}

func (s *AuthService) Register(ctx context.Context, req dto.Register) (dto.TokenResponse, error) {
	return s.client.Register(ctx, req)
}

func (s *AuthService) Login(ctx context.Context, req dto.Login) (dto.TokenResponse, error) {
	return s.client.Login(ctx, req)
}

func (s *AuthService) Refresh(ctx context.Context, req dto.Refresh) (dto.TokenResponse, error) {
	return s.client.Refresh(ctx, req)
}

func (s *AuthService) Logout(ctx context.Context, req dto.Refresh) error {
	return s.client.Logout(ctx, req)
}

func (s *AuthService) ValidateAccessToken(ctx context.Context, req dto.ValidateRequest) (dto.ValidateResponse, error) {
	return s.client.ValidateAccessToken(ctx, req)
}
