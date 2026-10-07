package keyprovider

import (
	"context"
	"sync"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/client"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
)

type KeyProvider struct {
	client client.AuthClient

	mu sync.RWMutex
	keys map[string]dto.PublicKey
}

func NewKeyProvider(
	client client.AuthClient,
) *KeyProvider {
	return &KeyProvider{
		client: client,
		keys: make(map[string]dto.PublicKey),
	}
}

func (p *KeyProvider) Refresh(ctx context.Context) error {
	keys, err := p.client.GetPublicKey(ctx)
	if err != nil {
		return err
	}

	newKeys := make(map[string]dto.PublicKey, len(keys))
	
	for _, key := range keys {
		newKeys[key.Id] = key
	}

	p.mu.Lock()
	p.keys = newKeys
	p.mu.Unlock()

	return nil
}

func (p *KeyProvider) Get(
	kid string,
) (dto.PublicKey, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	key, ok := p.keys[kid]

	return key, ok
}
