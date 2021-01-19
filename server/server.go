package server

import (
	"context"
	"errors"
	"log"

	"report-maker-server/server/controller"
	"report-maker-server/server/receiver"
	"report-maker-server/tools"

	"github.com/gin-gonic/gin"
)

func Serve() (err error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Server start in localhost:8080 ....")

	router := gin.New()

	newctx := tools.AppContex{Context: context.Background()}

	router.Use(appContext(&newctx))

	authorized := router.Group("/")
	authorized.Use(controller.BaseAuth())
	{
		authorized.GET("/home", controller.Home)
		authorized.GET("/reports", controller.Reports)
		authorized.GET("/login", controller.Login)

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
