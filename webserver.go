package groupietracker

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homeEvent(homePage *template.Template) {
	topFive := GetTopFive()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homePage.Execute(w, topFive)
	})
}

func artistsEvent(artistsPage *template.Template) {
	artists := UnMarshallArtists(GetArtists())
	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		artistsPage.Execute(w, artists)
	})
}

func loadTemplates(path string) (*template.Template, *template.Template) {
	var home = template.Must(template.ParseFiles(path + "index.html"))
	var artists = template.Must(template.ParseFiles(path + "artists.html"))
	return home, artists
}

func StartServer() {
	GetTopFive()
	var homePage, artistsPage = loadTemplates("./templates/")
	homeEvent(homePage)
	artistsEvent(artistsPage)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	fmt.Println("URL: http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
