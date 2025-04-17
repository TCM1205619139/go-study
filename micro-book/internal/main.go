package main

import (
	"micro-book/internal/repository"
	"micro-book/internal/repository/dao"
	"micro-book/internal/service"
	"micro-book/internal/web"
	"micro-book/internal/web/middlewares"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	db := initDatabase()
	user := initUser(db)

	server.Use(middlewares.NewLoginMiddlewareBuilder().
		IgnoreRequest(http.MethodPut, "/user").
		Build())

	user.RegisterRoutes(server.Group("/user"))
	server.Run(":8080")
}

func initDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:13306)/webook"), &gorm.Config{})
	if err != nil {
		panic("数据库初始化错误")
	}
	err = dao.InitTable(db)
	if err != nil {
		panic("表初始化错误")
	}
	return db
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	usvc := service.NewUserService(ur)
	user := web.NewUserHandler(usvc)

	return user
}
