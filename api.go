package groupietracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
