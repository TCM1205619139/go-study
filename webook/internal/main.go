package main

import (
	"micro-book/config"
	"micro-book/internal/repository"
	"micro-book/internal/repository/dao"
	"micro-book/internal/service"
	"micro-book/internal/web"
	"micro-book/internal/web/middlewares"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	// store, err := redis.NewStore(10, "tcp", config.Config.Redis.Addr, "", "")
	// if err != nil {
	// 	panic("redis初始化错误")
	// }
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"X-Jwt-Token"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	server.Use(sessions.Sessions("mysession", store))
	db := initDatabase()
	user := initUser(db)

	server.Use(middlewares.NewLoginMiddlewareBuilder().
		IgnoreRequest(http.MethodPut, "/user").
		IgnoreRequest(http.MethodPost, "/user").
		IgnoreRequest(http.MethodGet, "/user/test").
		Build())

	user.RegisterRoutes(server.Group("/user"))
	server.Run(":8081")
}

func initDatabase() *gorm.DB {
	// db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:13306)/webook"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN), &gorm.Config{})
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
