package main

import (
	"clanrece.com/EchoPong/controller"
	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	r := gin.Default()

	userGroup := r.Group("user")
	{
		userGroup.POST("/login", controller.LoginIn)
		userGroup.GET("/detail", controller.UserDetail)
		userGroup.GET("/all", controller.ListUsers)
	}

	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	GormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	gplus.Init(GormDB)

	r.Run(":8000")
}
