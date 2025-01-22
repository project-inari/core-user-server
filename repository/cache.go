package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/project-inari/core-user-server/dto"
	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	client                 *redis.Client
	keyUserVerifiedAccount string
	ttlUserVerifiedAccount time.Duration
}

// CacheRepositoryConfig represents the configuration of the cache repository
type CacheRepositoryConfig struct {
	KeyUserVerifiedAccount string
	TTLUserVerifiedAccount time.Duration
}

// CacheRepositoryDependencies represents the dependencies of the cache repository
type CacheRepositoryDependencies struct {
	Client *redis.Client
}

// NewCacheRepository creates a new cache repository
func NewCacheRepository(c CacheRepositoryConfig, d CacheRepositoryDependencies) CacheRepository {
	return &cacheRepository{
		client:                 d.Client,
		keyUserVerifiedAccount: c.KeyUserVerifiedAccount,
		ttlUserVerifiedAccount: c.TTLUserVerifiedAccount,
	}
}

func (r *cacheRepository) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *cacheRepository) SetUserVerifiedAccount(ctx context.Context, p dto.UserVerifiedAccountCache) *redis.StatusCmd {
	marshalValue, _ := json.Marshal(p)
	return r.client.Set(ctx, fmt.Sprintf("%s:%s", r.keyUserVerifiedAccount, p.Username), marshalValue, r.ttlUserVerifiedAccount)
}
