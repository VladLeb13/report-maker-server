package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"report-maker-server/config"
	"report-maker-server/tools"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(ctx *gin.Context) {
	app_ctx := ctx.MustGet("app-context").(*tools.AppContex)
	cnf := app_ctx.Context.Value("config").(config.Config)

	path := cnf.Template_path

	tmpl, err := template.ParseFiles(path + "login.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(ctx.Writer, nil)
	}

}

//Logining -Авторизация в системе
func Logining(ctx *gin.Context) {
	session := sessions.Default(ctx)

	r := ctx.Request
	w := ctx.Writer

	rlogin := r.FormValue("login")
	rpass := r.FormValue("pass")

	var authResult bool
	var err error
	if authResult, err = FindInBase(rlogin, rpass); err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 301)
	}

	if authResult {
		session.Options(struct {
			Path     string
			Domain   string
			MaxAge   int
			Secure   bool
			HttpOnly bool
			SameSite http.SameSite
		}{Path: "/", Domain: "", MaxAge: 380, Secure: false, HttpOnly: false, SameSite: 0})
		session.Set("auth-session", uuid.New().String())
		err := session.Save()
		if err != nil {
			log.Println("Error session save", err)
		}
		ctx.Redirect(http.StatusFound, "/home")

	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}

}

func FindInBase(rlogin string, rpass string) (authResult bool, err error) {

	if rlogin == "vl" && rpass == "vl" {
		authResult = true
	}

	if rlogin == "client" && rpass == "clientpass" {
		authResult = true
	}
	return

}
