package client

import (
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
	linkv1 "github.com/SamuraiJeka/SnapLink-proto/gen/go/link/v1"
)

func toProtoCreateRequest(
	req dto.CreateLink,
	user_id uint64,
) *linkv1.CreateLinkRequest {
	return & linkv1.CreateLinkRequest{
		UserId: user_id,
		Url: req.Url,
	}
}

func toDtoLink(
	link *linkv1.Link,
) dto.Link {
	return dto.Link{
		Id: link.Id,
		UserId: link.UserId,
		BaseUrl: link.BaseUrl,
		ShortUrl: link.ShortUrl,
	}
}
