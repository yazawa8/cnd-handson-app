package repository

import (
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/user/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserByID(userID uuid.UUID) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func GetUserByName(name string) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("name = ?", name).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *model.User) (*model.User, error) {
	if err := db.DB.Model(&model.User{}).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *model.User) (*model.User, error) {
	if err := db.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	user, err := GetUserByID(user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(userID uuid.UUID) error {
	if err := db.DB.Where("id = ?", userID).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
