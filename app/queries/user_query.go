package queries

import (
	"gorm-authentication/app/model"
)

func CreateUser(user *model.User) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
