package hndl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/parallelcointeam/pci/hlp"
	"github.com/parallelcointeam/pci/mod"
	"github.com/parallelcointeam/pci/rps"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// Handlers
var templates = make(map[string]*template.Template)

var JDB, _ = scribble.New("./jdb/", nil)

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	//templates["404"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/404.gohtml", "tpl/hlp/search.gohtml"))

	templates["login"] = template.Must(template.ParseFiles("tpl/login.gohtml"))
	templates["index"] = template.Must(template.ParseFiles("tpl/index.gohtml"))

}

// for GET
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	//	var body, _ = hlp.LoadFile("tpl/login.html")
	//	fmt.Fprintf(response, body)

	renderTemplate(w, "login", "login", nil)

}

// for POST
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if !hlp.IsEmpty(name) && !hlp.IsEmpty(pass) {
		// Database check for user data!
		_userIsValid := rps.UserIsValid(name, pass)

		if _userIsValid {
			SetCookie(name, response)
			redirectTarget = "/index"
		} else {
			redirectTarget = "/"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// // for GET
// func AdminPageHandler(response http.ResponseWriter, request *http.Request) {
// 	var body, _ = hlp.LoadFile("tpl/admin.html")

// 	lngJDB, err := JDB.ReadAll("lang")
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}

// 	langs := []mod.Lang{}
// 	for _, lng := range lngJDB {
// 		lang := mod.Lang{}
// 		if err := json.Unmarshal([]byte(lng), &lang); err != nil {
// 			fmt.Println("Error", err)
// 		}
// 		langs = append(langs, lang)
// 	}

// 	fmt.Fprintf(response, body)
// }

// for GET
func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		//		var indexBody, _ = hlp.LoadFile("tpl/index.html")
		//		fmt.Fprintf(response, indexBody, userName)
		lngJDB, err := JDB.ReadAll("lang")
		if err != nil {
			fmt.Println("Error", err)
		}
		langs := []mod.Lang{}
		for _, lng := range lngJDB {
			lang := mod.Lang{}
			if err := json.Unmarshal([]byte(lng), &lang); err != nil {
				fmt.Println("Error", err)
			}
			langs = append(langs, lang)
		}
		renderTemplate(w, "index", "index", langs)

	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// for POST
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	ClearCookie(response)
	http.Redirect(response, request, "/", 302)
}

// for POST
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	lang := r.FormValue("lang")
	title := r.FormValue("title")
	welcome := r.FormValue("welcome")
	intro := r.FormValue("intro")
	about := r.FormValue("about")

	var LNG mod.Lang = mod.Lang{
		Title:   title,
		Welcome: welcome,
		Intro:   intro,
		About:   about,
	}

	JDB.Write("lang", lang, LNG)
}

// Cookie

func SetCookie(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
