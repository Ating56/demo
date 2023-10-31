package middleware

import (
	"github.com/gin-gonic/gin"
	"my_blog_back/common"
	"my_blog_back/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取authorization header
		tokenString := context.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			context.Abort()
			return
		}

		// 因为 “Bearer ”一共占7位，所以从第7位截取
		tokenString = tokenString[7:]

		// 调用ParseToken方法，在tokenString中解析出claims
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid { // 如果有error或token无效
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			context.Abort()
			return
		}

		// 验证通过后获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 验证用户是否存在
		if userId == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			context.Abort()
			return
		}

		// 用户存在，将user信息写入上下文
		context.Set("user", user)
		context.Next()
	}
}
