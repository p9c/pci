package rts

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/parallelcointeam/pci/mod"
)

func AmpHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// lang := vars["lang"]
	// page := vars["page"]
	site := mod.Site{}
	if err := JDB.Read("site", "meta", &site); err != nil {
		fmt.Println("Error", err)
	}
	home := mod.Home{}
	if err := JDB.Read("en", "home", &home); err != nil {
		fmt.Println("Error", err)
	}
	data := struct {
		Site mod.Site `json:"site"`
		Home mod.Home `json:"home"`
	}{
		site,
		home,
	}
	//	tmpl, _ := template.New("").ParseFiles("./tpl/amp/index.gohtml", "./tpl/amp/lyt/home.gohtml", "./tpl/amp/inc/amp.gohtml", "./tpl/amp/inc/nav.gohtml", "./tpl/amp/inc/amp-css-home.gohtml", "./tpl/amp/inc/footer.gohtml")
	tmpl, _ := template.New("").ParseFiles("./tpl/amp/index.gohtml", "./tpl/amp/lyt/home.gohtml", "./tpl/amp/inc/nav.gohtml", "./tpl/amp/inc/amp.gohtml", "./tpl/amp/inc/amp-css-home.gohtml", "./tpl/amp/inc/footer.gohtml")
	tmpl.ExecuteTemplate(w, "home", data)
}
