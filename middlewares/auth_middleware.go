package middlewares

import (
	"net/http"
	"onycom/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization 헤더에서 토큰 가져오기
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization 헤더가 필요합니다."})
			c.Abort()
			return
		}

		// 토큰 형식 검사 (Bearer 토큰)
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 Authorization 헤더 형식입니다."})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 토큰 검증
		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			// 토큰 만료된 경우 DB에 저장된 리프레시 토큰 확인
			if err == utils.ErrTokenExpired {
				// DB에서 userID를 기반으로 리프레시 토큰 확인
				err = utils.CheckRefreshTokenInDB(userID)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "리프레시 토큰이 만료되었거나 유효하지 않습니다."})
					c.Abort()
					return
				}

				// 리프레시 토큰이 유효하면 새로운 액세스 토큰 발급
				newAccessToken, err := utils.GenerateToken(userID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "새로운 액세스 토큰 발급에 실패했습니다."})
					c.Abort()
					return
				}

				// 새 액세스 토큰을 응답에 추가
				c.JSON(http.StatusOK, gin.H{
					"access_token": newAccessToken,
				})
				return
			}

			// 그 외의 경우 유효하지 않은 토큰 처리
			c.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 토큰입니다."})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
