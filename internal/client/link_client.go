package client

import (
	"context"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
	linkv1 "github.com/SamuraiJeka/SnapLink-proto/gen/go/link/v1"

	"google.golang.org/grpc"
)

type LinkClient struct {
	client linkv1.LinkServiceClient
}

func NewLinkClient(conn *grpc.ClientConn) *LinkClient {
	return &LinkClient{
		client: linkv1.NewLinkServiceClient(conn),
	}
}

func (c *LinkClient) Create(
	ctx context.Context,
	req dto.CreateLink,
	user_id uint64,
) (dto.Link, error) {
	resp, err := c.client.Create(
		ctx,
		toProtoCreateRequest(req, user_id),
	)
	if err != nil {
		return dto.Link{}, err
	}

	return toDtoLink(resp.Link), nil
}

func (c *LinkClient) Get(
	ctx context.Context,
	idx uint64,
) (dto.Link, error) {
	resp, err := c.client.Get(
		ctx,
		&linkv1.GetLinkRequest{
			Id: idx,
		},
	)
	if err != nil {
		return dto.Link{}, err
	}

	return toDtoLink(resp.Link), nil
}

func (c *LinkClient) GetByUser(
	ctx context.Context,
	user_id uint64,
	limit uint32,
	offset uint32,
) ([]dto.Link, error) {
	resp, err := c.client.GetByUser(
		ctx,
		&linkv1.GetLinkByUserRequest{
			UserId: user_id,
			Limit: limit,
			Offset: offset,
		},
	)
	if err != nil {
		return nil, err 
	}

	links := make([]dto.Link, 0, len(resp.Link))

	for _, link  := range resp.Link {
		links = append(links, toDtoLink(link))
	}

	return links, nil
}

func (c *LinkClient) Delete(
	ctx context.Context,
	idx uint64,
) error {
	_, err := c.client.Delete(
		ctx,
		&linkv1.DeleteLinkRequest{
			Id: idx,
		},
	)
	return err
}
