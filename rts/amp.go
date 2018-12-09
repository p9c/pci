package rts

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/pci/mod"
)

func AmpHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]
	page := vars["page"]
	site := mod.Site{}
	if err := JDB.Read("site", "meta", &site); err != nil {
		fmt.Println("Error", err)
	}
	var pageData interface{}
	if err := JDB.Read(lang, page, &pageData); err != nil {
		fmt.Println("Error", err)
	}
	data := struct {
		Site mod.Site    `json:"site"`
		Lang string      `json:"lang"`
		Page interface{} `json:"page"`
	}{
		site,
		lang,
		pageData,
	}
	//	tmpl, _ := template.New("").ParseFiles("./tpl/amp/index.gohtml", "./tpl/amp/lyt/home.gohtml", "./tpl/amp/inc/amp.gohtml", "./tpl/amp/inc/nav.gohtml", "./tpl/amp/inc/amp-css-home.gohtml", "./tpl/amp/inc/footer.gohtml")
	tmpl, _ := template.New("").ParseFiles(
		"./tpl/icons/logo.gohtml",
		"./tpl/icons/icons.gohtml",
		"./tpl/amp/lyt/base.gohtml",
		"./tpl/amp/"+page+".gohtml",
		"./tpl/amp/inc/nav.gohtml",
		"./tpl/amp/inc/amp.gohtml",
		"./tpl/amp/inc/amp-basecss.gohtml",
		"./tpl/amp/inc/amp-basecssplgs.gohtml",
		"./tpl/amp/inc/amp-css.gohtml",
		"./tpl/amp/inc/footer.gohtml")
	//tmpl, _ := template.New("").ParseFiles("./tpl/amp/lyt/base.gohtml")
	fmt.Println("fdfdfdffdfdfdffdfdfdffdfdfdffdfdfdffdfdfdffdfdfdffdfdfdffdfdfdf", page)
	tmpl.ExecuteTemplate(w, "base", data)
}
