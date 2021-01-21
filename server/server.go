package server

import (
	"context"
	"errors"
	"log"
	"report-maker-server/server/receiver"

	"report-maker-server/server/controller"
	"report-maker-server/tools"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Serve() (err error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Server start in localhost:8080 ....")

	router := gin.Default()

	newctx := tools.AppContex{Context: context.Background()}

	router.Use(appContext(&newctx))
	router.Use(gin.Logger())

	store := cookie.NewStore([]byte("secret"))

	authorized := router.Group("/")
	authorized.Use(sessions.Sessions("auth-session", store))
	authorized.Use(controller.BaseAuth())
	{
		authorized.GET("/home", controller.Home)
		authorized.POST("/home", controller.Home)

		authorized.GET("/reports", controller.Reports)
		authorized.POST("/reports", controller.Home)

		authorized.GET("/login", controller.Login)
		authorized.POST("/login", controller.Login)

		authorized.POST("/logining", controller.Logining)
		authorized.POST("/upload", receiver.Upload)

	}

	router.Run(":8080")

	return errors.New("Server shutdown")
}

func appContext(app_context *tools.AppContex) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app-context", app_context)
	}
}
