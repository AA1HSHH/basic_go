package main

import (
	"basic-go/webook/internal/repository"
	"basic-go/webook/internal/repository/dao"
	"basic-go/webook/internal/service"
	"basic-go/webook/internal/web"
	"basic-go/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()
	initUser(server, db)
	server.Run(":8080")
}

func initDB() *gorm.DB {
	dsn := "root:root@tcp(43.142.48.34:13316)/webook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUser(server *gin.Engine, db *gorm.DB) {
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	c := web.NewUserHandler(us)
	c.RegisterRouters(server)
}
func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("ssid", store))
	login := &middleware.LoginMiddlewareBuilder{}
	server.Use(login.CheckLogin())
	return server
}
