package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/pci/hndl"
)

var r = mux.NewRouter()

func main() {

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./tpl/css/"))))

	r.HandleFunc("/", hndl.LoginPageHandler) // GET

	r.HandleFunc("/index", hndl.IndexPageHandler).Methods("GET")
	r.HandleFunc("/admin", hndl.AdminHandler).Methods("POST")
	r.HandleFunc("/login", hndl.LoginHandler).Methods("POST")

	//r.HandleFunc("/admin", hndl.AdminPageHandler).Methods("GET")

	r.HandleFunc("/logout", hndl.LogoutHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
