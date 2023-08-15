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
	var DisplayDetails API.DisplayDetails
	DisplayDetails.Id = APIcall[idNumber].Id
	DisplayDetails.Image = APIcall[idNumber].Image
	DisplayDetails.Name = APIcall[idNumber].Name
	DisplayDetails.Member = APIcall[idNumber].Member
	DisplayDetails.Creationdate = APIcall[idNumber].Creationdate
	DisplayDetails.FirstAlbum = APIcall[idNumber].FirstAlbum
	DisplayDetails.Relations = API.Relations(idNumber)
	DisplayDetails.Locations = API.Locations(idNumber)
	DisplayDetails.Dates = API.Dates(idNumber)
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		fmt.Println(err.Error()) // XXX
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "details.html", DisplayDetails)
}
