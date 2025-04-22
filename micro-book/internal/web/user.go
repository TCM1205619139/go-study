package web

import (
	"micro-book/internal/domain"
	"micro-book/internal/service"
	"net/http"
	"strconv"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	jwt "github.com/golang-jwt/jwt/v5"

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
	BirthdayRegexp := regexp.MustCompile(
		`^((19|20)\d{2})-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`,
		regexp.None,
	)
	NickNameRegexp := regexp.MustCompile(
		`^(?:[\p{L}\p{N}_]){1,32}$`,
		regexp.None,
	)
	DescriptionRegexp := regexp.MustCompile(
		`^(?:[\p{L}\p{N}_]){1,128}$`,
		regexp.None,
	)

	return &UserHandler{
		svc:               svc,
		EmailRegexp:       EmailRegexp,
		PasswordRegexp:    PasswordRegexp,
		BirthdayRegexp:    BirthdayRegexp,
		NickNameRegexp:    NickNameRegexp,
		DescriptionRegexp: DescriptionRegexp,
	}
}

/**
* 定义和User相关的路由
* 可以方便利用包变量特性进行测试
 */
type UserHandler struct {
	svc               *service.UserService
	EmailRegexp       *regexp.Regexp
	PasswordRegexp    *regexp.Regexp
	BirthdayRegexp    *regexp.Regexp
	NickNameRegexp    *regexp.Regexp
	DescriptionRegexp *regexp.Regexp
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
	type SigninRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req SigninRequest

	if err := ctx.Bind(&req); err != nil {
		return
	}

	user, err := u.svc.SigninService(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.InvalidUserOrPasswordError {
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
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
		"id":  user.Id,
	})
}

func (u *UserHandler) SigninJWT(ctx *gin.Context) {
	type SigninRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req SigninRequest

	if err := ctx.Bind(&req); err != nil {
		return
	}

	user, err := u.svc.SigninService(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.InvalidUserOrPasswordError {
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": user.Email,
		"Id":    user.Id,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("12345678123456781234567812345678"))

	ctx.Header("x-jwt-token", tokenString)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
		"id":  user.Id,
	})
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	type EditRequest struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		NickName    string `json:"nickName"`
		Birthday    string `json:"birthday"`
		Description string `json:"description"`
	}
	id, ok := ctx.Params.Get("id")
	if !ok {
		return
	}
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}
	var req EditRequest
	if err := ctx.Bind(&req); err != nil {
		return
	}

	ok, err = u.EmailRegexp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "邮箱不合规")
		return
	}
	ok, err = u.BirthdayRegexp.MatchString(req.Birthday)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "生日信息不合规")
		return
	}
	ok, err = u.NickNameRegexp.MatchString(req.NickName)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "昵称不合规")
		return
	}
	ok, err = u.DescriptionRegexp.MatchString(req.Description)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "个人描述信息不合规")
		return
	}

	user, err := u.svc.EditService(ctx, intId, domain.User{
		Email:       req.Email,
		Password:    req.Password,
		NickName:    req.NickName,
		Birthday:    req.Birthday,
		Description: req.Description,
	})
	if err == service.InvalidUserEmailError {
		ctx.String(http.StatusOK, "邮箱不存在")
	}
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
		"user": gin.H{
			"Email":       user.Email,
			"NickName":    user.NickName,
			"Birthday":    user.Birthday,
			"Description": user.Description,
		},
	})
}
func (u *UserHandler) Profile(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		return
	}
	type ProfileRequest struct {
		Email string `json:"email"`
	}
	var req ProfileRequest
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.ProfileService(ctx, id)
	if err == service.InvalidUserEmailError {
		ctx.String(http.StatusOK, "未查询到相关信息")
	}
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"email":       user.Email,
		"nickName":    user.NickName,
		"birthday":    user.Birthday,
		"description": user.Description,
	})
}
func (*UserHandler) Delete(ctx *gin.Context) {

}

func (u *UserHandler) RegisterRoutes(ug *gin.RouterGroup) {
	ug.PUT("", u.Signup)
	ug.POST("", u.SigninJWT)
	ug.POST("/:id", u.Edit)
	ug.GET("/:id", u.Profile)
	ug.DELETE("/:id", u.Delete)
}
