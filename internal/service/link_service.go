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

func (s *LinkService) Create(ctx context.Context, req dto.CreateLink) (dto.Link, error) {
	
}

func (s *LinkService) Get(ctx context.Context, idx uint64) (dto.Link, error)  {

}

func (s *LinkService) GetByUser(ctx context.Context, limit uint32, offset uint32) ([]dto.Link, error) {

}

func (s *LinkService) Delete(ctx context.Context, idx uint64) error {

}
