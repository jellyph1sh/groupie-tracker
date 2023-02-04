package groupietracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type artists struct {
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

func TestAPI() []byte { //Fonction pour récupérer les infos sur l'artiste/groupe
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body // renvoie un tableau de bytes
}

func UnMarshall(data []byte) { // Fonction pour UnMarshall le résultat de la requête API
	var tab []artists
	err := json.Unmarshal([]byte(data), &tab)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(tab) // Pour l'instant il est simplement print (dans le futur il seras renvoyé pour la page HTML)
}
