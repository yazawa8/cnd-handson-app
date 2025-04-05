package repository

import (
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/user/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/db"
	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*model.User, error) {
	var users model.User
	if err := db.DB.Where("email = ?", email).First(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &users, nil
}

func CreateUser(user *model.User) error {
	if err := db.DB.Model(&model.User{}).Create(user).Error; err != nil {
		return err
	}
	return nil
}
