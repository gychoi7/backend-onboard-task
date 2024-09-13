package services

import (
	_ "errors"
	"gorm.io/gorm"
	"onycom/models"
	"onycom/utils"
	"time"
)

// models.User 가져와서 CRUD 함수 구현
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := utils.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func CreateUser(user *models.User) (models.User, error) {
	err := utils.DB.Create(user).Error
	if err != nil {
		return *user, err
	}
	return *user, nil
}

func SaveOrUpdateToken(userID uint, refreshToken string) error {
	var token models.Token

	// DB에서 해당 유저의 토큰 찾기
	err := utils.DB.Where("user_id = ?", userID).First(&token).Error

	if err != nil {
		// 토큰이 존재하지 않는 경우 새로 생성
		if err == gorm.ErrRecordNotFound {
			token = models.Token{
				UserID:    userID,
				Refresh:   refreshToken,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			err = utils.DB.Create(&token).Error
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// 토큰이 존재하면 업데이트
		token.Refresh = refreshToken
		token.UpdatedAt = time.Now()
		err = utils.DB.Save(&token).Error
		if err != nil {
			return err
		}
	}

	return nil
}
