package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world")
	})
	server.GET("/hello/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(http.StatusOK, "hello, "+id)
	})
	server.GET("/order/", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "order, "+id)
	})
	server.GET("/views/*.html", func(ctx *gin.Context) {
		id := ctx.Param(".html")
		ctx.String(http.StatusOK, "views, "+id)
	})
	server.Run(":8080")
}
