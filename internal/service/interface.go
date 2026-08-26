package service

import (
	"context"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
)


type LinkClient interface {
	Create(ctx context.Context, req dto.CreateLink, user_id uint64) (dto.Link, error)
	Get(ctx context.Context, idx uint64) (dto.Link, error)
	GetByUser(ctx context.Context, user_id uint64, limit uint32, offset uint32) ([]dto.Link, error)
	Delete(ctx context.Context, idx uint64) error
}

type UserClient interface {
	Get(ctx context.Context, idx uint64)
	Patch(ctx context.Context)
}

type AuthClient interface {
	Registration(ctx context.Context)
	Login(ctx context.Context)
	Refresh(ctx context.Context)
}
