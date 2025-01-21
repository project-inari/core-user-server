package service

import (
	"context"

	"github.com/project-inari/core-user-server/dto"
)

func (s *service) SignUp(ctx context.Context, req dto.SignUpReq) (*dto.SignUpRes, error) {
	user := dto.UserEntity{
		Username:       req.Username,
		UID:            req.UID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		PhoneNo:        req.PhoneNo,
		Email:          req.Email,
		SelectedLocale: req.SelectedLocale,
	}

	err := s.databaseRepository.CreateNewUser(user)
	if err != nil {
		return nil, err
	}

	cacheVal := dto.UserVerifiedAccountCache{
		Username:  req.Username,
		UID:       req.UID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		PhoneNo:   req.PhoneNo,
		Email:     req.Email,
	}

	c := s.cacheRepository.SetUserVerifiedAccount(ctx, cacheVal)
	if c.Err() != nil {
		return nil, c.Err()
	}

	return &dto.SignUpRes{
		Username: req.Username,
		UID:      req.UID,
		Success:  true,
	}, nil
}
