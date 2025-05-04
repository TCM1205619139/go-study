package middlewares

import (
	"fmt"
	"micro-book/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type IgnoreRequest struct {
	method string
	path   string
}

type LoginMiddlewareBuilder struct {
	requests []IgnoreRequest
}

func (middleware *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, r := range middleware.requests {
			if r.method == ctx.Request.Method && r.path == ctx.Request.URL.Path {
				return
			}
		}
		// session := sessions.Default(ctx)
		// id := session.Get("userId")
		authorization := ctx.GetHeader("Authorization")

		if authorization == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := &web.JWTUserClaims{}

		token, err := jwt.ParseWithClaims(
			authorization,
			claims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte("12345678123456781234567812345678"), nil
			},
		)
		if err != nil || !token.Valid || claims.Id == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Println("token", token.Raw)
		fmt.Println("token", token.Valid)
		// if id == nil {
		// 	ctx.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		ctx.Set("claims", claims)
	}
}

func (middleware *LoginMiddlewareBuilder) IgnoreRequest(method string, path string) *LoginMiddlewareBuilder {
	middleware.requests = append(middleware.requests, IgnoreRequest{
		method: method,
		path:   path,
	})

	return middleware
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}
