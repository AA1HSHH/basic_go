package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddlewareBuilder struct {
}

func (*LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		url := context.Request.URL.Path
		if url == "/users/signup" || url == "/users/login" {
			return
		}

		sess := sessions.Default(context)
		if sess.Get("userId") == nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
