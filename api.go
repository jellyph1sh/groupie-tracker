package groupietracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func getAPIData(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

/*------------------------------------------------------*/
/*						Artists:						*/
/*------------------------------------------------------*/
func unmarshalArtistsData(data []byte) Artists {
	if data == nil {
		return nil
	}
	var artists Artists
	err := json.Unmarshal([]byte(data), &artists)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return artists
}

func GetArtistsData() Artists {
	return unmarshalArtistsData(getAPIData("https://groupietrackers.herokuapp.com/api/artists"))
}

/*------------------------------------------------------*/
/*						Relation:						*/
/*------------------------------------------------------*/
func unmarshalRelation(data []byte) Relation {
	if data == nil {
		return Relation{}
	}
	var relation Relation
	err := json.Unmarshal(data, &relation)
	if err != nil {
		fmt.Println(err)
		return Relation{}
	}
	return relation
}

func GetRelationData() Relation {
	return unmarshalRelation(getAPIData("https://groupietrackers.herokuapp.com/api/relation"))
}

/*------------------------------------------------------*/
/*						Locations:						*/
/*------------------------------------------------------*/
/*
func unmarshalLocations(data []byte) Locations {
	var locations Locations
	err := json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println(err)
		return Locations{}
	}
	return locations
}

func GetLocationsData() Locations {
	return unmarshalLocations(getAPIData("https://groupietrackers.herokuapp.com/api/locations"))
}
*/
/*------------------------------------------------------*/
/*						Dates:							*/
/*------------------------------------------------------*/
/*
func unmarshalDates(data []byte) Dates {
	var dates Dates
	err := json.Unmarshal(data, &dates)
	if err != nil {
		fmt.Println(err)
		return Dates{}
	}
	return dates
}

func getDatesData() Dates {
	return unmarshalDates(getAPIData("https://groupietrackers.herokuapp.com/api/dates"))
}
*/
/*------------------------------------------------------*/
/*						Concerts:						*/
/*------------------------------------------------------*/
func setDisplayDate(relation *Relation) Relation {
	for i := 0; i < len(relation.Index); i++ {
		for place := range relation.Index[i].DatesLocations {
			var newPlace = ""
			for li := 0; li < len(place); li++ {
				if li == 0 {
					if rune(place[li]) < 97 {
						newPlace += string(place[li])
					} else {
						newPlace += string(byte(rune(place[li]) - 32))
					}
				} else if place[li] == byte('-') || place[li] == byte('_') {
					newPlace += " "
				} else if place[li-1] == byte('-') || place[li-1] == byte('_') {
					newPlace += string(byte(rune(place[li]) - 32))
				} else {
					newPlace += string(place[li])
				}
			}
			relation.Index[i].DatesLocations[newPlace] = relation.Index[i].DatesLocations[place]
			delete(relation.Index[i].DatesLocations, place)
		}
	}
	return *relation
}

func GetConcerts() Concerts {
	relation := GetRelationData()
	if relation.Index == nil {
		return Concerts{}
	}
	setDisplayDate(&relation)
	artists := GetArtistsData()
	if artists == nil {
		return Concerts{}
	}
	return Concerts{Artists: artists, Relation: relation}
}

/*------------------------------------------------------*/
/*						Top Five:						*/
/*------------------------------------------------------*/
func generateRandomNumbers() []int {
	var blacklist []int
	var randomNumbers []int

	rand.Seed(time.Now().UnixNano())

	for len(randomNumbers) < 5 {
		generated := false
		for !generated {
			found := false
			number := rand.Intn(52) + 1
			for i := 0; i < len(blacklist); i++ {
				if blacklist[i] == number {
					found = true
				}
			}
			if !found {
				randomNumbers = append(randomNumbers, number)
				blacklist = append(blacklist, number)
				generated = true
			}
		}
	}
	return randomNumbers
}

func GetTopFive() []Artist {
	url := "https://groupietrackers.herokuapp.com/api/artists/"
	randNbs := generateRandomNumbers()
	var topFive []Artist
	for x := 0; x < len(randNbs); x++ {
		var artist Artist
		modifiedURL := url
		modifiedURL += strconv.Itoa(randNbs[x])
		data := getAPIData(modifiedURL)
		if data == nil {
			return nil
		}
		err := json.Unmarshal([]byte(data), &artist)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		topFive = append(topFive, artist)
	}
	return topFive
}
