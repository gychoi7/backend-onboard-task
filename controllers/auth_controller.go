package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onycom/models"
	"onycom/services"
	"onycom/utils"
	"regexp"
	"time"
)


type SignInRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// @Summary SignUp
// @Description SignUp
// @Tags users
// @Accept  json
// @Param body body SignUpRequest true "이메일과 비밀번호"
// @Produce  json
// @Success 200 {string} string "회원가입에 성공했습니다."
// @Failure 400 {string} string "이메일 형식이 올바르지 않습니다."
// @Failure 400 {string} string "중복된 이메일입니다."
// @Failure 400 {string} string "비밀번호는 8자 이상이어야 합니다."
// @Failure 400 {string} string "비밀번호 암호화에 실패했습니다."
// @Failure 500 {string} string "회원가입에 실패했습니다."
// @Router /users/signup [post]
func SignUp(c *gin.Context) {
	var input SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	emailRegex := '@'
	matched, err := regexp.MatchString(string(emailRegex), input.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !matched {
		c.JSON(400, gin.H{"error": "이메일 형식이 올바르지 않습니다."})
		return
	}

	//DB에 중복되는 이메일이 있는지 확인
	_, err = services.GetUserByEmail(input.Email)
	if err == nil {
		c.JSON(400, gin.H{"error": "중복된 이메일입니다."})
		return
	}

	//비밀번호가 8자 이상인지 확인
	passwordRegex := `^.{8,}$`

	matchedPassword, err := regexp.MatchString(passwordRegex, input.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if !matchedPassword {
		c.JSON(400, gin.H{"error": "비밀번호는 8자 이상이어야 합니다."})
		return
	}

	//비밀번호 암호화
	//email과 password를 받아서 service.SignUp 함수를 호출하고 결과를 받아서 response
	hashedPassword, salt, err := utils.MakePasswordHash(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "비밀번호 암호화에 실패했습니다."})
		return
	}
	user := models.User{
		Email:     input.Email,
		Password:  hashedPassword,
		Salt:      salt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "회원가입에 실패했습니다."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "회원가입에 성공했습니다."})
}

// @Summary SignIn
// @Description SignIn
// @Tags users
// @Accept  json
// @Param body body SignInRequest true "이메일과 비밀번호"
// @Produce  json
// @Success 200 {string} string "로그인에 성공했습니다."
// @Failure 401 {string} string "가입하지 않은 이메일입니다."
// @Failure 401 {string} string "비밀번호가 올바르지 않습니다."
// @Failure 500 {string} string "토큰 생성에 실패했습니다."
// @Failure 500 {string} string "리프레시 토큰 생성에 실패했습니다."
// @Failure 500 {string} string "리프레시 토큰 저장에 실패했습니다."
// @Router /users/signin [post]
func SignIn(c *gin.Context) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := services.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "가입하지 않은 이메일입니다."})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password, user.Salt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "비밀번호가 올바르지 않습니다."})
		return
	}

	//로그인 성공시 토큰 발급
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "토큰 생성에 실패했습니다."})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "리프레시 토큰 생성에 실패했습니다."})
		return
	}

	err = services.SaveOrUpdateToken(user.ID, refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "리프레시 토큰 저장에 실패했습니다."})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
