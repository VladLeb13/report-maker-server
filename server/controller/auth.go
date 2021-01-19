package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var (
	sessionStore = sessions.NewCookieStore([]byte("vl") /*, []byte("qwertyui-qwertyu")*/)
	block1       = []byte("vl")
	block2       = []byte("qwertyui-qwertyu")
	s            = securecookie.New(block1, block2)
)

//BaseAuth - аутентификация (MiddleWare)
func BaseAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		w := ctx.Writer
		r := ctx.Request

		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		cookie, _ := r.Cookie("auth-session")

		if cookie == nil {
			//log.Println("NO auth")
			if r.RequestURI == "/login" || r.RequestURI == "/logining" || r.RequestURI == "/upload" {
				return
			} else {
				http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			}

		} else {
			/*
				Добавить проверку сессии по имени пользоватлея и id в базе
			*/
			//fmt.Println(cookie.String())
			value := map[string]string{}
			sCookie, err := r.Cookie("cookies")
			err = s.Decode("cookies", sCookie.Value, &value)
			if err != nil {
				log.Println(err)
			} else {
				for key, val := range value {
					log.Println("Ключ: ", key, "\n", "Значение: ", val)
				}
			}
			session, err := sessionStore.Get(r, "auth-session")
			if session == nil {
				log.Println("Session not found")
			} else {
				var user = session.Values["uname"]
				log.Println("user:", user)

				flashes := session.Flashes()

				if len(flashes) > 0 {
					log.Println("flash:", flashes[0])
				}

			}

		}
	}
}
