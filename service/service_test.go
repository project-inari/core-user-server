package service

import (
	"context"
	"errors"
	"testing"

	"github.com/project-inari/core-user-server/dto"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

type mockDatabaseRepository struct {
	user *dto.UserEntity
	err  error
}

func (m *mockDatabaseRepository) CreateNewUser(_ dto.UserEntity) error {
	return m.err
}

func (m *mockDatabaseRepository) GetUserByUsername(_ string) (*dto.UserEntity, error) {
	return m.user, m.err
}

type mockCacheRepository struct {
	getRes                    *redis.StringCmd
	setUserVerifiedAccountRes *redis.StatusCmd
}

func (m *mockCacheRepository) Get(_ context.Context, _ string) *redis.StringCmd {
	return m.getRes
}

func (m *mockCacheRepository) SetUserVerifiedAccount(_ context.Context, _ dto.UserVerifiedAccountCache) *redis.StatusCmd {
	return m.setUserVerifiedAccountRes
}

const (
	mockUsername       = "username"
	mockUID            = "uid"
	mockFirstName      = "first_name"
	mockLastName       = "last_name"
	mockPhoneNo        = "phone_no"
	mockEmail          = "email"
	mockSelectedLocale = "EN"
)

func TestSignUp(t *testing.T) {
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		signupReq := dto.SignUpReq{
			Username:       mockUsername,
			UID:            mockUID,
			FirstName:      mockFirstName,
			LastName:       mockLastName,
			PhoneNo:        mockPhoneNo,
			Email:          mockEmail,
			SelectedLocale: mockSelectedLocale,
		}

		expectedRes := &dto.SignUpRes{
			Username: mockUsername,
			UID:      mockUID,
			Success:  true,
		}

		mockDatabaseRepository := &mockDatabaseRepository{
			user: &dto.UserEntity{
				Username:       mockUsername,
				UID:            mockUID,
				FirstName:      mockFirstName,
				LastName:       mockLastName,
				PhoneNo:        mockPhoneNo,
				Email:          mockEmail,
				SelectedLocale: mockSelectedLocale,
			},
			err: nil,
		}

		mockCacheRepository := &mockCacheRepository{
			setUserVerifiedAccountRes: redis.NewStatusCmd(ctx, "OK"),
		}

		svc := New(Dependencies{
			DatabaseRepository: mockDatabaseRepository,
			CacheRepository:    mockCacheRepository,
		})

		res, err := svc.SignUp(ctx, signupReq)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("error - when save to db failed", func(t *testing.T) {
		signupReq := dto.SignUpReq{
			Username:       mockUsername,
			UID:            mockUID,
			FirstName:      mockFirstName,
			LastName:       mockLastName,
			PhoneNo:        mockPhoneNo,
			Email:          mockEmail,
			SelectedLocale: mockSelectedLocale,
		}

		mockDatabaseRepository := &mockDatabaseRepository{
			err: errors.New("error"),
		}

		svc := New(Dependencies{
			DatabaseRepository: mockDatabaseRepository,
		})

		res, err := svc.SignUp(ctx, signupReq)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestInquiry(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		expectedRes := &dto.InquiryRes{
			Username:  mockUsername,
			UID:       mockUID,
			FirstName: mockFirstName,
			LastName:  mockLastName,
			PhoneNo:   mockPhoneNo,
			Email:     mockEmail,
		}

		mockDatabaseRepository := &mockDatabaseRepository{
			user: &dto.UserEntity{
				Username:  mockUsername,
				UID:       mockUID,
				FirstName: mockFirstName,
				LastName:  mockLastName,
				PhoneNo:   mockPhoneNo,
				Email:     mockEmail,
			},
			err: nil,
		}

		svc := New(Dependencies{
			DatabaseRepository: mockDatabaseRepository,
		})

		res, err := svc.Inquiry(mockUsername)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("error - when get user by username from db failed", func(t *testing.T) {
		mockDatabaseRepository := &mockDatabaseRepository{
			err: errors.New("error"),
		}

		svc := New(Dependencies{
			DatabaseRepository: mockDatabaseRepository,
		})

		res, err := svc.Inquiry(mockUsername)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
