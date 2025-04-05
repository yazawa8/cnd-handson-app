package service

import (
	"time"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/user/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/user/repository"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/auth"
	"github.com/google/uuid"
)

// Register処理
func RegisterUser(Email string, Password string) error {
	HashPassword, err := auth.HashPassword(Password)
	if err != nil {
		return err
	}
	user := model.User{
		ID:           uuid.New(), // UUIDを生成
		Name:         "Guest",    // デフォルト値
		Email:        Email,
		PasswordHash: HashPassword,
		RoleID:       uuid.Nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return repository.CreateUser(&user)
}

// Auth処理
func AuthenticateUser(email, password string) (*model.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	if !auth.ComparePassword(user.PasswordHash, password) {
		return nil, nil
	}
	return user, nil
}
