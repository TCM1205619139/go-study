package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
		session := sessions.Default(ctx)
		id := session.Get("mysession")
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Printf("id: %s", id)
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
