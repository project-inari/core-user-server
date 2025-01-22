package service

import (
	"github.com/project-inari/core-user-server/dto"
)

func (s *service) Inquiry(username string) (*dto.InquiryRes, error) {
	user, err := s.databaseRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return &dto.InquiryRes{
		Username:  user.Username,
		UID:       user.UID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		PhoneNo:   user.PhoneNo,
		Email:     user.Email,
	}, nil
}
