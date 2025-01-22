// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/project-inari/core-user-server/dto"
	"github.com/project-inari/core-user-server/repository"
)

// Port represents the service layer functions
type Port interface {
	SignUp(ctx context.Context, req dto.SignUpReq) (*dto.SignUpRes, error)
	Inquiry(username string) (*dto.InquiryRes, error)
}

type service struct {
	databaseRepository repository.DatabaseRepository
	cacheRepository    repository.CacheRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	DatabaseRepository repository.DatabaseRepository
	CacheRepository    repository.CacheRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		databaseRepository: d.DatabaseRepository,
		cacheRepository:    d.CacheRepository,
	}
}
