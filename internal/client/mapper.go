package client

import (
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
	authv1 "github.com/SamuraiJeka/SnapLink-proto/gen/go/auth/v1"
	linkv1 "github.com/SamuraiJeka/SnapLink-proto/gen/go/link/v1"
)

func toProtoCreateRequest(
	req dto.CreateLink,
	user_id uint64,
) *linkv1.CreateLinkRequest {
	return &linkv1.CreateLinkRequest{
		UserId: user_id,
		Url: req.Url,
	}
}

func toDtoLink(link *linkv1.Link) dto.Link {
	return dto.Link{
		Id: link.Id,
		UserId: link.UserId,
		BaseUrl: link.BaseUrl,
		ShortUrl: link.ShortUrl,
	}
}

func toProtoRegister(req dto.Register) *authv1.RegisterRequest {
	return &authv1.RegisterRequest{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}
}

func toProtoLogin(req dto.Login) *authv1.LoginRequest {
	return &authv1.LoginRequest{
		Email: req.Email,
		Password: req.Password,
	}
}

func toDtoToken(token *authv1.TokenResponse) dto.TokenResponse {
	return dto.TokenResponse{
		Access_token: token.AccessToken,
		Refresh_token: token.RefreshToken,
	}
}

func toDtoValidate(vld *authv1.ValidateAccessTokenResponse) dto.ValidateResponse {
	return dto.ValidateResponse{
		Valid: vld.Valid,
		User_id: vld.UserId,
	}
}

func toDtoPublicKey(req *authv1.GetPyblicKeyResponse) []dto.PublicKey {
	public_keys := make([]dto.PublicKey, len(req.Keys))
	
	for _, key := range req.Keys {
		public_keys = append(
			public_keys,
			dto.PublicKey{
				Id: key.Kid,
				Algorithm: key.Algorithm,
				KeyType: key.KeyType,
				Curve: key.Curve,
				Key: key.PublicKey,
			},
		)
	}

	return public_keys
}
