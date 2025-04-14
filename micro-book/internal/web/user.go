package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* 定义和User相关的路由
* 可以方便利用包变量特性进行测试
 */
type UserHandler struct {
}

func (*UserHandler) Login(ctx *gin.Context) {

}
func (*UserHandler) Edit(ctx *gin.Context) {

}
func (*UserHandler) Profile(ctx *gin.Context) {

}
func (*UserHandler) Delete(ctx *gin.Context) {

}
func (*UserHandler) Page(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "page.html", nil)
}

func (u *UserHandler) RegisterRoutes(ug *gin.RouterGroup) {
	ug.PUT("/:id", u.Login)
	ug.POST(":id", u.Edit)
	ug.GET("/:id", u.Profile)
	ug.DELETE(":id", u.Delete)
	ug.GET("/login", u.Page)
}
