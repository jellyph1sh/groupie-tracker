package groupietracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Artists []struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Locations struct {
	Index []struct {
		ID        int
		Locations []string
	}
}

type Dates struct {
	Index []struct {
		ID    int
		Dates []string
	}
}

type Relation struct {
	Index []struct {
		ID             int
		DatesLocations map[string][]string // prend comme clé la ville puis comme valeur un tableau de dates
	}
}

/*---------------------- Artist API ----------------------*/

func GetArtist() []byte {
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
	return body // renvoie un tableau de bytes
}

func UnMarshallArtists(data []byte) Artists {
	var tab Artists
	err := json.Unmarshal([]byte(data), &tab)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(tab)
	return tab
}

/*---------------------- Artist API ----------------------*/

/*---------------------- Location API ----------------------*/

func GetLocation() []byte {
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
	fmt.Println(tab)
	return tab
}

/*---------------------- Location API ----------------------*/

/*---------------------- Relation API ----------------------*/

func GetRelation() []byte {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
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

func UnMarshallRelation(data []byte) Relation {
	var tab Relation
	err := json.Unmarshal(data, &tab)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(tab)
	return tab
}

/*---------------------- Relation API ----------------------*/

/*---------------------- Dates API ----------------------*/

func GetDate() []byte {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(string(body))
	// fmt.Println()
	return body
}

func UnMarshallDate(data []byte) Dates {
	var tab Dates
	err := json.Unmarshal(data, &tab)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(tab)
	return tab
}

/*---------------------- Dates API ----------------------*/
