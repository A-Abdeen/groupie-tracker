package gt

import (
	"fmt"
	// API "gt/webapp/API"
	// "html/template"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request){
if r.Method != "GET" {
	ErrorHandler(w, r, http.StatusMethodNotAllowed)
	return
}
// Verify Request Pattern (Path)
if r.URL.Path != "/search/" {
	ErrorHandler(w, r, http.StatusNotFound)
	return
}
typedData := r.FormValue("search")
fmt.Println(string(typedData))
}