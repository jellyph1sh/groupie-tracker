package groupietracker

import (
	"strings"
)

func GetSearch(search string) Concerts {
	data := GetConcerts()
	var newArtists Artists
	var newIndex []struct {
		ID             int                    `json:"id"`
		DatesLocations map[string]interface{} `json:"datesLocations"`
	}
	for i := 0; i < len(data.Artists); i++ {
		if strings.Contains(strings.ToUpper(data.Artists[i].Name), strings.ToUpper(search)) {
			newArtists = append(newArtists, data.Artists[i])
			for j := 0; j < len(data.Relation.Index); j++ {
				if data.Relation.Index[j].ID == data.Artists[i].Id {
					newIndex = append(newIndex, data.Relation.Index[j])
				}
			}
		}
	}
	if len(newArtists) > 0 && len(newIndex) > 0 {
		data.Artists = newArtists
		data.Relation.Index = newIndex
	}
	return data
}
