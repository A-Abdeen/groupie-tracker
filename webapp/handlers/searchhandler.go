package gt

import (
	"fmt"
	API "gt/webapp/API"
	"html/template"
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
var idToReturn int
var dataToReturn API.SearchReturn
// var matchingData string
for _, oneArtist := range APIcall {
if oneArtist.Name == typedData {
	idToReturn = oneArtist.Id
	// matchingData = typedData + "=" + oneArtist.Name
}
	}
oneArtist := APIcall[idToReturn]
dataToReturn = API.SearchReturn{Name:oneArtist.Name, Image:oneArtist.Image}
t, err := template.ParseFiles(HtmlTmpl...)
if err != nil {
	ErrorHandler(w, r, http.StatusInternalServerError)
	return
}
fmt.Println(dataToReturn)
fmt.Println(string(typedData))
w.WriteHeader(http.StatusOK)
t.ExecuteTemplate(w, "search.html", dataToReturn)

}