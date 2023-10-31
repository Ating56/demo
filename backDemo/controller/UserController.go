package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"my_blog_back/common"
	"my_blog_back/dto"
	"my_blog_back/model"
	"my_blog_back/response"
	"net/http"
)

func Register(context *gin.Context) {
	DB := common.GetDB()

	requestUser := model.User{}
	err := context.Bind(&requestUser)
	if err != nil {
		return
	}
	// 获取参数
	name := requestUser.Name
	password := requestUser.Password

	// 数据验证
	if len(password) < 6 {
		// gin.H 返回给前端响应数据
		// 422状态码，接受表单数据，但是数据校验不过，无法处理提交过来的数据
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于等于6")
		return
	}
	if len(name) == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名长度必须大于0")
		return
	}

	//新建用户
	//密码加密存储
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Password: string(hashPassword),
	}
	DB.Create(&newUser)

	// 注册成功也直接发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v !", err)
		return
	}

	// 返回结果
	response.Success(context, gin.H{"token": token}, "注册成功")
}

func Login(context *gin.Context) {
	DB := common.GetDB()

	requestUser := model.User{}
	err := context.Bind(&requestUser)
	if err != nil {
		return
	}
	// 获取参数
	name := requestUser.Name
	password := requestUser.Password

	// 数据验证
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于等于6")
		return
	}
	if len(name) == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名长度必须大于0")
		return
	}

	// 判断密码是否正确
	var user model.User
	DB.Where("name = ?", name).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(context, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v !", err)
		return
	}

	// 返回结果
	response.Success(context, gin.H{"token": token}, "登录成功")
}

func Info(context *gin.Context) {
	user, _ := context.Get("user")

	response.Success(context, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}
