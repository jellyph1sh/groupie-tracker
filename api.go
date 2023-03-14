package groupietracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Artists []struct {
	SpotifyID    string
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	}
}

type Dates []struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	}
}

type Relation struct {
	Index []struct {
		ID             int                    `json:"id"`
		DatesLocations map[string]interface{} `json:"datesLocations"`
	} `json:"index"`
}

type Concerts struct {
	Artists  Artists
	Relation Relation
}

/*---------------------- Artist API ----------------------*/

func GetArtists() []byte {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return body
}

func UnMarshallArtists(data []byte) Artists {
	var tab Artists
	err := json.Unmarshal([]byte(data), &tab)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < len(tab); i++ {
		spotifyArtist := searchProfile(tab[i].Name)
		tab[i].SpotifyID = spotifyArtist.Artists.Artists[0].ID.String()
	}
	return tab
}

func GetTopFive() []Artist {
	tab := RandNumber()
	var temp Artist
	var result []Artist
	for numb := 0; numb < len(tab); numb++ {
		url := "https://groupietrackers.herokuapp.com/api/artists/"
		url += strconv.Itoa(tab[numb])
		req, _ := http.NewRequest("GET", url, nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal([]byte(body), &temp)
		result = append(result, temp)
	}
	return result
}

func RandNumber() []int {
	var tabRandNumb []int
	for i := 0; i <= 4; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(1)
		x := rand.Intn(50)
		tabRandNumb = append(tabRandNumb, x)
	}
	return tabRandNumb
}

/*---------------------- Artist API ----------------------*/

/*---------------------- Location API ----------------------*/

func GetLocations() []byte {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return body
}

func UnMarshallLocations(data []byte) Locations {
	var tab Locations
	err := json.Unmarshal(data, &tab)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tab
}

/*---------------------- Location API ----------------------*/

/*---------------------- Relation API ----------------------*/

func GetRelation() []byte {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return body
}

func UnMarshallRelation(data []byte) Relation {
	var tab Relation
	err := json.Unmarshal(data, &tab)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tab
}

func SetDisplayDate(relation *Relation) Relation {
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

/*---------------------- Relation API ----------------------*/

/*---------------------- Dates API ----------------------*/

func GetDates() []byte {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return body
}

func UnMarshallDates(data []byte) Dates {
	var tab Dates
	err := json.Unmarshal(data, &tab)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tab
}

/*---------------------- Dates API ----------------------*/

func GetConcerts() Concerts {
	relation := UnMarshallRelation(GetRelation())
	SetDisplayDate(&relation)
	artists := UnMarshallArtists(GetArtists())
	concerts := Concerts{Artists: artists, Relation: relation}
	return concerts
}

/*---------------------- Sorts ----------------------*/

func GetSort(sortName string) Artists {
	data := UnMarshallArtists(GetArtists())
	switch sortName {
	case "alphabet":
		return quickSort(data, 0, len(data)-1, "partition_alphabet")
	case "date":
		return quickSort(data, 0, len(data)-1, "partition_creationDate")
	case "members":
		return quickSort(data, 0, len(data)-1, "partition_Members")
	}
	return nil
}

func quickSort(list Artists, leftIndex int, rightIndex int, nameSort string) Artists {
	var pivotIndex int // number to determine where to split the array
	if leftIndex < rightIndex {
		switch nameSort {
		case "partition_alphabet":
			pivotIndex = partition_Alphabet(list, leftIndex, rightIndex)
		case "partition_creationDate":
			pivotIndex = partition_CreationDate(list, leftIndex, rightIndex)
		case "partition_Members":
			pivotIndex = partition_Members(list, leftIndex, rightIndex)
		}
		quickSort(list, leftIndex, pivotIndex-1, nameSort)  // sort left side of the array
		quickSort(list, pivotIndex+1, rightIndex, nameSort) // sort right side of the array
	}
	return list
}

/*---------------------- Creation Date Sort ----------------------*/

func partition_CreationDate(list Artists, left int, right int) int {
	pivot := list[right]
	i := left - 1
	for j := left; j < right; j++ {
		if list[j].CreationDate < pivot.CreationDate {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

/*---------------------- Creation Date Sort ----------------------*/

/*---------------------- Alphabet Sort ----------------------*/

func partition_Alphabet(list Artists, left int, right int) int {
	pivot := list[right]
	i := left - 1
	for j := left; j < right; j++ {
		if list[j].Name < pivot.Name {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

/*---------------------- Alphabet Sort ----------------------*/

/*---------------------- Members Sort ----------------------*/

func partition_Members(list Artists, left int, right int) int {
	pivot := list[right]
	i := left - 1
	for j := left; j < right; j++ {
		if len(list[j].Members) < len(pivot.Members) {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

/*---------------------- Members Sort ----------------------*/

/*---------------------- Sorts ----------------------*/
