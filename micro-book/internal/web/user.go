package web

import (
	"micro-book/internal/domain"
	"micro-book/internal/service"
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(svc *service.UserService) *UserHandler {
	EmailRegexp := regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
		regexp.None,
	)
	PasswordRegexp := regexp.MustCompile(
		`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]{8,}$`,
		regexp.None,
	)

	return &UserHandler{
		svc:            svc,
		EmailRegexp:    EmailRegexp,
		PasswordRegexp: PasswordRegexp,
	}
}

/**
* 定义和User相关的路由
* 可以方便利用包变量特性进行测试
 */
type UserHandler struct {
	svc            *service.UserService
	EmailRegexp    *regexp.Regexp
	PasswordRegexp *regexp.Regexp
}

func (u *UserHandler) Signup(ctx *gin.Context) {
	type SignupRequest struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignupRequest
	if err := ctx.Bind(&req); err != nil {
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusInternalServerError, "两次输入的密码不相同")
		return
	}

	ok, err := u.EmailRegexp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "邮箱不合规")
		return
	}

	err = u.svc.SignupService(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.DuplicateUserEmailError {
		// ctx.String(http.StatusOK, "邮箱已被注册")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "邮箱已被注册",
		})
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "ok")
}

func (u *UserHandler) Signin(ctx *gin.Context) {
	type SinginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req SinginRequest

	if err := ctx.Bind(&req); err != nil {
		return
	}

	user, err := u.svc.SigninService(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.InvavidUserOrPasswordError {
		ctx.String(http.StatusOK, "用户名或密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	session := sessions.Default(ctx)
	session.Set("userId", user.Id)
	session.Save()
	ctx.String(http.StatusOK, "登录成功")
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
	ug.PUT("/:id", u.Signup)
	ug.POST("/", u.Signin)
	ug.POST("/:id", u.Edit)
	ug.GET("/:id", u.Profile)
	ug.DELETE("/:id", u.Delete)
	ug.GET("/login", u.Page)
}
