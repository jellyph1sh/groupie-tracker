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

func GetApi(name string) any {
	switch name {
	case "topfive":
		return GetTopFive()
	case "artists":
		return UnMarshallArtists(GetArtists())
	case "concerts":
		return GetConcerts()
	}
	return nil
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

func GetConcerts() *Concerts {
	relation := UnMarshallRelation(GetRelation())
	SetDisplayDate(&relation)
	artists := UnMarshallArtists(GetArtists())
	concerts := Concerts{Artists: artists, Relation: relation}
	return &concerts
}

/*---------------------- Sort ----------------------*/

func GetSort(name string) Artists {
	data := UnMarshallArtists(GetArtists())
	switch name {
	case "alphabet":
		return ExecutePartSort(data, "FindMiddlePart_Alphabet")
	case "date":
		return ExecutePartSort(data, "FindMiddlePart_CreationDate")
	case "members":
		return ExecutePartSort(data, "FindMiddlePart_Members")
	}
	return nil
}

/*---------------------- Creation Date Sort ----------------------*/

func FindMiddlePart_CreationDate(List Artists, first int, last int) int {
	middleNumber := List[last]
	i := first
	j := first
	for j < last {
		if List[j].CreationDate <= middleNumber.CreationDate {
			List[i], List[j] = List[j], List[i]
			i += 1
		}
		j += 1
	}
	List[last], List[i] = List[i], List[last]
	return i
}

/*---------------------- Creation Date Sort ----------------------*/

/*---------------------- Alphabet Sort ----------------------*/

func FindMiddlePart_Alphabet(List Artists, first int, last int) int {
	middleNumber := List[last]
	i := first
	j := first
	for j < last {
		if List[j].Name <= middleNumber.Name {
			List[i], List[j] = List[j], List[i]
			i++
		}
		j++
	}
	List[last], List[i] = List[i], List[last]
	return i
}

/*---------------------- Alphabet Sort ----------------------*/

/*---------------------- Members Sort ----------------------*/

func FindMiddlePart_Members(List Artists, first int, last int) int {
	middleNumber := List[last]
	i := first
	j := first
	for j < last {
		if len(List[j].Members) <= len(middleNumber.Members) {
			List[i], List[j] = List[j], List[i]
			i += 1
		}
		j += 1
	}
	List[last], List[i] = List[i], List[last]
	return i
}

func RecursivePartitionSort(List Artists, first int, last int, nameSort string) {
	var i int
	switch nameSort {
	case "FindMiddlePart_Alphabet":
		i = FindMiddlePart_Alphabet(List, first, last)
	case "FindMiddlePart_CreationDate":
		i = FindMiddlePart_CreationDate(List, first, last)
	case "FindMiddlePart_Members":
		i = FindMiddlePart_Members(List, first, last)
	}
	if first < last {
		RecursivePartitionSort(List, first, i-1, nameSort)
		RecursivePartitionSort(List, i+1, last, nameSort)
	}
}

func ExecutePartSort(List Artists, nameSort string) Artists {
	fmt.Println(len(List) - 1)
	RecursivePartitionSort(List, 0, len(List)-1, nameSort)
	return List
}

/*---------------------- Members Sort ----------------------*/
