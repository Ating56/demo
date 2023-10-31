package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"my_blog_back/common"
	"my_blog_back/router"
	"os"
)

func main() {
	InitConfig()
	// 创建服务
	ginServer := gin.Default()

	// 连接数据库，调用common中的InitDB
	common.InitDB()

	// 访问地址，接收请求
	ginServer = router.CollectRoute(ginServer)

	// 使用viper修改端口
	//port := viper.GetString("server.port")
	//if port != "" {
	//	panic(ginServer.Run(":" + port))
	//}

	// 创建端口
	panic(ginServer.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
