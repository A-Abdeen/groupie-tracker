package gt

import (
	API "gt/webapp/API"
	"strings"
)

func Tosearch(typedData string) []API.Artists {
	typedDataInt := API.Atoi(typedData)
	typedData = strings.ToUpper(typedData)
	var dataToReturn []API.Artists
	var toDisplay string
	ifMatching := false  
	for i, oneArtist := range APIcall {
		if strings.ToUpper(oneArtist.Name) == typedData {
			toDisplay = "Search matches name " + oneArtist.Name
			ifMatching = true
		}
		for _, oneMember := range oneArtist.Member {
			if strings.ToUpper(oneMember) == typedData {
				toDisplay = "Search matches member name " + oneMember
				ifMatching = true
			}
		}
		if typedDataInt == oneArtist.Creationdate {
			toDisplay = "Search matches creation date " + typedData
			ifMatching = true
		}
		firstAlbum := API.Atoi(oneArtist.FirstAlbum)
		if typedDataInt == firstAlbum {
			toDisplay = "Search matches First Album date " + typedData
			ifMatching = true
		}
		for _, oneLocation := range oneArtist.Locations {
			if strings.ToUpper(oneLocation) == typedData {
				toDisplay = "Search matches Location " + oneLocation
				ifMatching = true			
			}
		}
		for _, oneDate := range oneArtist.Dates {
			oneDateInt := API.Atoi(oneDate)
			if oneDateInt == typedDataInt {
				toDisplay = "Search matches Date " + oneDate
				ifMatching = true
			}
		}
		if ifMatching {
		ArtistDetails := APIcall[i]
		ArtistDetails.SearchResult = toDisplay
		dataToReturn = append(dataToReturn, ArtistDetails)
	}
	ifMatching = false
	}
	return dataToReturn
}
