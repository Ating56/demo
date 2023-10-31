package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(context *gin.Context, httpStatus int, code int, data gin.H, message string) {
	context.JSON(httpStatus, gin.H{"code": code, "data": data, "message": message})
}

func Success(context *gin.Context, data gin.H, message string) {
	Response(context, http.StatusOK, 200, data, message)
}

func Fail(context *gin.Context, data gin.H, message string) {
	Response(context, http.StatusBadRequest, 400, data, message)
}
