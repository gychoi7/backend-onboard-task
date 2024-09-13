package services

import (
	"errors"
	"onycom/models"
	"onycom/utils"
	"time"
)

var (
	ErrNotAuthorized = errors.New("권한이 없습니다.")
	ErrPostNotFound  = errors.New("게시글을 찾을 수 없습니다.")
)

func CreatePost(post *models.Post) error {
	err := utils.DB.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPosts(userID uint, offset int, limits int) (int64, []models.Post, error) {
	var totalCount int64
	err := utils.DB.Model(&models.Post{}).Where("user_id = ?", userID).Count(&totalCount).Error
	if err != nil {
		return 0, nil, err
	}

	var posts []models.Post
	err = utils.DB.Where("user_id = ?", userID).Offset(offset).Limit(limits).Find(&posts).Error
	if err != nil {
		return 0, nil, err
	}
	return totalCount, posts, nil
}

func GetPost(id uint) (models.Post, error) {
	var post models.Post
	err := utils.DB.Where("id = ?", id).First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func UpdatePost(id uint, userID uint, title string, content string) error {
	var post models.Post
	if err := utils.DB.First(&post, id).Error; err != nil {
		return ErrPostNotFound
	}

	if post.UserID != userID {
		return ErrNotAuthorized
	}

	post.UpdatedAt = time.Now()

	return utils.DB.Save(&post).Error
}

func DeletePost(id uint, userID uint) error {
	var post models.Post
	if err := utils.DB.First(&post, id).Error; err != nil {
		return ErrPostNotFound
	}

	if post.UserID != userID {
		return ErrNotAuthorized
	}

	return utils.DB.Delete(&post).Error
}
