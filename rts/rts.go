package rts

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/parallelcointeam/pci/hlp"
	"github.com/parallelcointeam/pci/mod"
	"github.com/parallelcointeam/pci/rps"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var JDB, _ = scribble.New("./jdb/data/", nil)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("").ParseFiles("./tpl/admin/login.gohtml", "./tpl/admin/base.gohtml")
	tmpl.ExecuteTemplate(w, "base", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if !hlp.IsEmpty(name) && !hlp.IsEmpty(pass) {
		_userIsValid := rps.UserIsValid(name, pass)

		if _userIsValid {
			SetCookie(name, w)
			redirectTarget = "/admin/"
		} else {
			redirectTarget = "/"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func AdminHomeHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		var content interface{}
		if err := JDB.Read("site", "meta", &content); err != nil {
			fmt.Println("Error", err)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/admin.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", content)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
func AdminLangHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		var content interface{}
		if err := JDB.Read("site", "meta", &content); err != nil {
			fmt.Println("Error", err)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/admin.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", content)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]
	page := vars["page"]
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		var home mod.Home
		if err := JDB.Read(lang, page, &home); err != nil {
			fmt.Println("Error", err)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/"+page+".gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", home)

	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearCookie(w)
	http.Redirect(w, r, "/", 302)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	subtitle := r.FormValue("subtitle")
	welcome := r.FormValue("welcome")
	about := r.FormValue("about")
	features := r.FormValue("features")
	feature1 := r.FormValue("feature1")
	feature1txt := r.FormValue("feature1txt")
	feature2 := r.FormValue("feature2")
	feature2txt := r.FormValue("feature2txt")
	feature3 := r.FormValue("feature3")
	feature3txt := r.FormValue("feature3txt")
	feature4 := r.FormValue("feature4")
	feature4txt := r.FormValue("feature4txt")
	gallery := r.FormValue("gallery")
	specs := r.FormValue("specs")
	spec1 := r.FormValue("spec1")
	spec2 := r.FormValue("spec2")
	spec3 := r.FormValue("spec3")
	spec4 := r.FormValue("spec4")
	moto1 := r.FormValue("moto1")
	moto2 := r.FormValue("moto2")
	contact := r.FormValue("contact")
	footer := r.FormValue("footer")

	var HOME mod.Home = mod.Home{
		Title:       title,
		SubTitle:    subtitle,
		Welcome:     welcome,
		About:       about,
		Features:    features,
		Feature1:    feature1,
		Feature1txt: feature1txt,
		Feature2:    feature2,
		Feature2txt: feature2txt,
		Feature3:    feature3,
		Feature3txt: feature3txt,
		Feature4:    feature4,
		Feature4txt: feature4txt,
		Gallery:     gallery,
		Specs:       specs,
		Spec1:       spec1,
		Spec2:       spec2,
		Spec3:       spec3,
		Spec4:       spec4,
		Moto1:       moto1,
		Moto2:       moto2,
		Contact:     contact,
		Footer:      footer,
	}
	JDB.Write("lang", "home", HOME)
}

func SetCookie(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func ClearCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func GetUserName(w *http.Request) (userName string) {
	if cookie, err := w.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
