package keyprovider

import (
	"context"
	"crypto/ed25519"
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrWrongTokenType = errors.New("worng token type")
	ErrKeyNotFound = errors.New("signing key not found")
)

const AccessToken = "access"

type Claims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

type Verifier struct {
	keys *KeyProvider
}

func NewVerifier(keys *KeyProvider) *Verifier {
	return &Verifier{
		keys: keys,
	}
}

func (v *Verifier) Verify(
	ctx context.Context,
	tokenStr string,
) (uint64, error) {
	userId, err := v.verify(tokenStr)
	if err == nil {
		return userId, nil
	}

	if !errors.Is(err, ErrKeyNotFound) {
		return 0, err
	}

	if err := v.keys.Refresh(ctx); err != nil {
		return 0, ErrInvalidToken
	}

	return v.verify(tokenStr)
}

func (v *Verifier) verify(tokenStr string) (uint64, error) {
	var claims Claims

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodEdDSA {
				return nil, ErrInvalidToken
			}

			kid, ok := token.Header["kid"].(string)
			if !ok || kid == "" {
				return nil, ErrInvalidToken
			}

			key, ok := v.keys.Get(kid)
			if !ok {
				return nil, ErrKeyNotFound
			}

			if key.Algorithm != "EdDSA" {
				return nil, ErrInvalidToken
			}

			if len(key.Key) != ed25519.PublicKeySize {
				return nil, ErrInvalidToken
			}

			return ed25519.PublicKey(key.Key), nil
		},
	)

	if err != nil {
		return 0, ErrInvalidToken
	}

	if !token.Valid {
		return 0, ErrInvalidToken
	}

	if claims.Type != AccessToken {
		return 0, ErrWrongTokenType
	}

	userId, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, ErrInvalidToken
	}

	return userId, nil
}
