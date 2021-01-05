package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//TODO: context: config templates file
	path := "/home/worker/Studing/report-maker-server/src/server/templates/"

	tmpl, err := template.ParseFiles(path + "login.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(w, nil)
	}

}

//Logining -Авторизация в системе
func Logining(w http.ResponseWriter, r *http.Request) {

	rlogin := r.FormValue("login")
	rpass := r.FormValue("pass")

	var authResult bool
	var err error
	if authResult, err = FindInBase(rlogin, rpass); err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 301)
	}

	if authResult {
		//TODO: context: cookeis name
		session, err := sessionStore.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["uname"] = &rlogin
		session.ID = uuid.New().String()
		session.Options = &sessions.Options{
			//todo: set normal max ege
			MaxAge: 120,
		}
		var message string = "flash-message"
		session.AddFlash(&message)

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		value := map[string]string{}
		value["secure"] = "cooookiess"
		if encoded, err := s.Encode("cookies", &value); err == nil {
			httpCookie := &http.Cookie{
				Name:  "cookies",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, httpCookie)

		}

		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)

	} else {
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
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
	//var db *sql.DB
	//db = data.GetSqlite()
	//defer db.Close()
	//
	//result := db.QueryRow("select login,password from users where login = $1", rlogin)
	//var login, pass string
	//err = result.Scan(&login, &pass)
	//if err != nil {
	//	return
	//}
	//
	//if pass != "" {
	//	bytePass := []byte(pass)
	//	err = bcrypt.CompareHashAndPassword(bytePass, []byte(rpass))
	//	if err != nil {
	//		return
	//	}
	//	authResult = true
	//}

}
