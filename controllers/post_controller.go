package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onycom/models"
	"onycom/services"
	"strconv"
	"time"
)

// @Summary GetPosts
// @Description GetPosts
// @Tags posts
// @Accept  json
// @Param Authorization header string true "Bearer {token}"
// @Param page query int false "페이지 번호"
// @Param limits query int false "페이지 당 게시글 수"
// @Produce  json
// @Success 200 {object} object "게시글 목록"
// @Failure 401 {string} string "인증되지 않은 사용자입니다."
// @Router /posts [get]
func GetPosts(c *gin.Context) {
	authUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "인증되지 않은 사용자입니다."})
		return
	}

	userID := authUserID.(uint)

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limits", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 페이지 번호입니다."})
		return
	}

	limits, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 limits 값입니다."})
		return
	}

	offset := (page - 1) * limits

	totalCount, posts, err := services.GetPosts(userID, offset, limits)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "게시글 목록을 불러오는데 실패했습니다."})
		return
	}

	totalPage := totalCount / int64(limits)
	if totalCount%int64(limits) != 0 {
		totalPage++
	}
	c.JSON(http.StatusOK, gin.H{
		"posts":      posts,
		"totalPosts": totalCount,
		"totalPage":  totalPage,
		"page":       page,
	})
}

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// @Summary CreatePost
// @Description CreatePost
// @Tags posts
// @Accept  json
// @Param Authorization header string true "Bearer {token}"
// @Param body body PostRequest true "제목과 내용"
// @Produce  json
// @Success 201 {string} string "게시글이 성공적으로 생성되었습니다."
// @Failure 400 {string} string "제목과 내용을 모두 입력해야 합니다."
// @Failure 401 {string} string "인증되지 않은 사용자입니다."
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	authUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "인증되지 않은 사용자입니다."})
		return
	}

	userID := authUserID.(uint)

	var input PostRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 입력 값입니다."})
		return
	}

	if input.Title == "" || input.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "제목과 내용을 모두 입력해야 합니다."})
		return
	}

	post := &models.Post{
		Title:     input.Title,
		Content:   input.Content,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 게시글 생성 서비스 호출
	if err := services.CreatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "게시글 생성 중 오류가 발생했습니다."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "게시글이 성공적으로 생성되었습니다."})
}

// @Summary GetPost
// @Description GetPost
// @Tags posts
// @Accept  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path int true "게시글 ID"
// @Produce  json
// @Success 200 {object} object "게시글"
// @Failure 400 {string} string "게시글 ID가 필요합니다."
// @Failure 400 {string} string "잘못된 게시글 ID입니다."
// @Failure 401 {string} string "인증되지 않은 사용자입니다."
// @Failure 404 {string} string "게시글을 찾을 수 없습니다."
// @Router /posts/{id} [get]

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "사용자 ID가 필요합니다."})
		return
	}

	parseID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 사용자 ID입니다."})
		return
	}

	postID := uint(parseID)

	post, err := services.GetPost(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "게시글을 찾을 수 없습니다."})
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary UpdatePost
// @Description UpdatePost
// @Tags posts
// @Accept  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path int true "게시글 ID"
// @Param body body PostRequest true "제목과 내용"
// @Produce  json
// @Success 200 {string} string "게시글이 성공적으로 수정되었습니다."
// @Failure 400 {string} string "게시글 ID가 필요합니다."
// @Failure 400 {string} string "잘못된 게시글 ID입니다."
// @Failure 400 {string} string "잘못된 입력 값입니다."
// @Failure 401 {string} string "인증되지 않은 사용자입니다."
// @Failure 403 {string} string "게시글 수정 권한이 없습니다."
// @Failure 404 {string} string "게시글을 찾을 수 없습니다."
// @Failure 500 {string} string "게시글 수정 중 오류가 발생했습니다."
// @Router /posts/{id} [put]
func UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "게시글 ID가 필요합니다."})
		return
	}

	parseID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 게시글 ID입니다."})
		return
	}

	postID := uint(parseID)

	var input PostRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 입력 값입니다."})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "인증되지 않은 사용자입니다."})
		return
	}

	err = services.UpdatePost(postID, userID.(uint), input.Title, input.Content)
	if err != nil {
		if err == services.ErrNotAuthorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "게시글 수정 권한이 없습니다."})
		} else if err == services.ErrPostNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "게시글을 찾을 수 없습니다."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "게시글 수정 중 오류가 발생했습니다."})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "게시글이 성공적으로 수정되었습니다."})
}

// @Summary DeletePost
// @Description DeletePost
// @Tags posts
// @Accept  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path int true "게시글 ID"
// @Produce  json
// @Success 200 {string} string "게시글이 성공적으로 삭제되었습니다."
// @Failure 400 {string} string "게시글 ID가 필요합니다."
// @Failure 400 {string} string "잘못된 게시글 ID입니다."
// @Failure 401 {string} string "인증되지 않은 사용자입니다."
// @Failure 403 {string} string "게시글 삭제 권한이 없습니다."
// @Failure 404 {string} string "게시글을 찾을 수 없습니다."
// @Failure 500 {string} string "게시글 삭제 중 오류가 발생했습니다."
// @Router /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "게시글 ID가 필요합니다."})
		return
	}

	parseID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 게시글 ID입니다."})
		return
	}

	postID := uint(parseID)

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "인증되지 않은 사용자입니다."})
		return
	}

	err = services.DeletePost(postID, userID.(uint))
	if err != nil {
		if err == services.ErrNotAuthorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "게시글 삭제 권한이 없습니다."})
		} else if err == services.ErrPostNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "게시글을 찾을 수 없습니다."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "게시글 삭제 중 오류가 발생했습니다."})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "게시글이 성공적으로 삭제되었습니다."})
}
