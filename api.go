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
		ID             int64
		DatesLocations map[string][]string
	}
}

/*---------------------- Artist API ----------------------*/

func GetArtist() []byte { //Fonction pour récupérer les infos sur l'artiste/groupe
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

func UnMarshallArtists(data []byte) Artists { // Fonction pour UnMarshall le résultat de la requête API Artiste
	var tab Artists
	err := json.Unmarshal([]byte(data), &tab)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(tab[0])
	return tab // Pour l'instant il est simplement print (dans le futur il seras renvoyé pour la page HTML)
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

func UnMarshallLocations(data []byte) Locations { // Fonction pour UnMarshall le résultat de la requête API Location
	var tab Locations
	json.Unmarshal(data, &tab)
	fmt.Println(tab)
	return tab
}

/*---------------------- Location API ----------------------*/

/*---------------------- Relation API ----------------------*/

func GetRelation() []byte { //Fonction pour récupérer les infos sur l'artiste/groupe
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
	return body // renvoie un tableau de bytes
}

func UnMarshallRelation(data []byte) Relation { // Fonction pour UnMarshall le résultat de la requête API Artiste
	var tab Relation
	json.Unmarshal(data, &tab)
	fmt.Println(tab)
	return tab
}

/*---------------------- Relation API ----------------------*/
