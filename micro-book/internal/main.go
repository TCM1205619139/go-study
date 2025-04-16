package main

import (
	"micro-book/internal/repository"
	"micro-book/internal/repository/dao"
	"micro-book/internal/service"
	"micro-book/internal/web"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()

	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:13306)/webook"), &gorm.Config{})
	if err != nil {
		panic("数据库初始化错误")
	}
	err = dao.InitTable(db)
	if err != nil {
		panic("表初始化错误")
	}
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	usvc := service.NewUserService(ur)
	user := web.NewUserHandler(usvc)
	user.RegisterRoutes(server.Group("/user"))

	server.Run(":8080")
}
