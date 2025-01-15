// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/project-inari/core-user-server/dto"
)

// DatabaseRepository represents the repository layer functions of database repository
type DatabaseRepository interface {
	CreateNewUser(newUser dto.UserEntity) error
}

// CacheRepository represents the repository layer functions of cache repository
type CacheRepository interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	SetUserVerifiedAccount(ctx context.Context, p dto.SignUpReq) *redis.StatusCmd
}
