package service

import (
	"context"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
)

type LinkService struct {
	client LinkClient
}

func NewLinkService(client LinkClient) *LinkService {
	return &LinkService{
		client: client,
	}
}

func (s *LinkService) Create(
	ctx context.Context,
	req dto.CreateLink,
	userId uint64,
) (dto.Link, error) {
	return s.client.Create(ctx, req, userId)
}

func (s *LinkService) Get(
	ctx context.Context,
	idx uint64,
) (dto.Link, error) {
	return s.client.Get(ctx, idx)
}

func (s *LinkService) GetByUser(
	ctx context.Context,
	userId uint64,
	limit uint32,
	offset uint32,
) ([]dto.Link, error) {
	return s.client.GetByUser(ctx, userId, limit, offset)
}

func (s *LinkService) Delete(
	ctx context.Context,
	idx uint64,
) error {
	return s.client.Delete(ctx, idx)
}
