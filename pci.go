package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/pci/rts"
)

var r = mux.NewRouter()

func main() {

	r.PathPrefix("/s/").Handler(http.StripPrefix("/s/", http.FileServer(http.Dir("./tpl/static/"))))
	// r.PathPrefix("/amp/").Handler(http.StripPrefix("/amp/", http.FileServer(http.Dir("./tpl/amp/"))))

	r.HandleFunc("/", rts.LoginPageHandler) // GET
	r.HandleFunc("/admin/", rts.AdminHomeHandler).Methods("GET")
	r.HandleFunc("/admin/{lang}", rts.AdminLangHandler).Methods("GET")
	r.HandleFunc("/admin/{lang}/{page}", rts.AdminPageHandler).Methods("GET")
	r.HandleFunc("/admin/", rts.AdminHandler).Methods("POST")
	r.HandleFunc("/login", rts.LoginHandler).Methods("POST")

	r.HandleFunc("/logout", rts.LogoutHandler).Methods("POST")

	r.HandleFunc("/api/{lang}", rts.ApiHandler)

	r.HandleFunc("/amp/{lang}", rts.AmpHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
