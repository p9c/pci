package rts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parallelcointeam/pci/mod"

	"github.com/gorilla/mux"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]
	getLang := mod.Home{}
	if err := JDB.Read("lang", lang, &getLang); err != nil {
		fmt.Println("Error", err)
	}
	jsonLang, err := json.Marshal(getLang)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(jsonLang))
}
