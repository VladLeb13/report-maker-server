package controller

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//BaseAuth - аутентификация (MiddleWare)
func BaseAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessions := sessions.Default(ctx)

		r := ctx.Request

		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		session := sessions.Get("auth-session")
		if session == nil {
			//log.Println("NO auth")
			if r.RequestURI == "/login" || r.RequestURI == "/logining" || r.RequestURI == "/upload" {
				return
			} else {
				ctx.Redirect(http.StatusFound, "/login")
			}
		}
	}
}
