package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"my_blog_back/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	// 需事先引入mysql
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	// 自动创建users表
	err2 := db.AutoMigrate(&model.User{})
	if err2 != nil {
		return nil
	}

	// DB实例赋值，即连接数据库
	DB = db
	return db
}

// GetDB 获取DB实例
func GetDB() *gorm.DB {
	return DB
}
