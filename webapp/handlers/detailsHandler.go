package gt

import (
	"fmt"
	API "gt/webapp/API"
	"html/template"
	"net/http"
	"strconv"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DetailsHandler is called.") // XXX
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/details" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	idNumber, _ := strconv.Atoi(r.FormValue("idNumber"))
	idNumber--
	var DisplayDetails API.Artists
	DisplayDetails = APIcall[idNumber]
	DisplayDetails.Relations = API.Relations(idNumber)
	DisplayDetails.Locations = API.Locations(idNumber)
	DisplayDetails.Dates = API.Dates(idNumber)
	fmt.Println(len(DisplayDetails.Member))
	DisplayDetails.Member = API.Arrangestring(DisplayDetails.Member)
	fmt.Println(DisplayDetails)
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		fmt.Println(err.Error()) // XXX
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "details.html", DisplayDetails)
}
