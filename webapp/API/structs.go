package gt

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Member       []string `json:"members"`
	Creationdate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type TmpAllConRel struct {
	Index []struct {
		Relation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type TmpLocations struct {
	Index []struct {
		LocationsDetailed []string `json:"locations"`
		DatesDetailed     string   `json:"dates"`
	} `json:"index"`
}

type TmpDates struct {
	Index []struct {
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Err struct {
	IsErr      bool
	Msg        string
	StatusCode int
}

/*
	Below struct is known as struct composition
	(or struct embedding)
	Allows using single struct to access all resources
	Checkout PassError method below
	Next step: adding stuct methods for specific handlers
*/

type WebHandler struct {
	*Artists
	Locations *TmpLocations
	Dates     *TmpDates
	Relations *TmpAllConRel
	*Err
}

type DisplayDetails struct {
	Id           int      
	Image        string   
	Name         string
	Member       []string `json:"members"`
	Creationdate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    TmpAllConRel
}

func (w WebHandler) PassError(x string, y int) WebHandler {
	errorResponse := WebHandler{
		Err: &Err{
			IsErr:      true,
			Msg:        x,
			StatusCode: y,
		},
	}
	return errorResponse
}

// func (w WebHandler) LoadDetails(id int) WebHandler {
// 	fullDetails := WebHandler{
// 		Artists: &Artists{
// 			Id: Artists[Id], 
// 			// Image:Artists[id].Image,
// 			// Name:Artists[id].Name,   
// 			// Member:Artists[id].Member, 
// 			// Creationdate:Artists[id].Creationdate,      
// 			// FirstAlbum:Artists[id].FirstAlbum,   
// 		},
// 	}
// 	return fullDetails
// }