package gt

import (
	"fmt"
	API "gt/webapp/API"
)

/*
The function below is meant to load list of files-
that are parsed for use of html/template library-
at a global scope.
*/

var HtmlTmpl []string // global variables to be used by other functions
var APIcall []API.Artists

func Init() {
	fmt.Println("Initializing Global Variable") // XXX
	HtmlTmpl = []string{
		"../webapp/static/base.html",
		"../webapp/static/details.html",
		"../webapp/static/error.html",
		// Add new html / template names here
	}
	fmt.Println("Global Variable initialized") // XXX
	APIcall = API.LoadArtist()   //used to unmarshal full data into APIcall
	allLocations := API.Locations()
	allDates := API.Dates()
	allRelations := API.Relations()
	for i, _ := range APIcall {   
	APIcall[i].Locations = allLocations[i]
	APIcall[i].Dates = allDates[i]
	APIcall[i].Relations = allRelations[i]
	}
}
