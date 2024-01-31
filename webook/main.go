package main

import (
	"basic-go/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	hdl := web.NewUserHandler()

	server := gin.Default()
	//server.Use(cors.New(cors.Config{
	//	AllowCredentials: true,
	//	AllowHeaders:     []string{"Content-Type", "Authorization"},
	//	AllowOriginFunc: func(origin string) bool {
	//		return true
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"authorization", "content-type"},
		AllowMethods:     []string{"POST"},
		AllowOrigins:     []string{"http://localhost:3000"},
		MaxAge:           12 * time.Hour,
	}))

	hdl.RegisterRouters(server)

	server.Run(":8080")
}
