package groupietracker

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homeEvent(homePage *template.Template) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homePage.Execute(w, nil)
	})
}

func concertsEvent(concertsPage *template.Template) {
	var concerts = GetConcerts()
	http.HandleFunc("/concerts", func(w http.ResponseWriter, r *http.Request) {
		concertsPage.Execute(w, concerts)
	})
}

func loadTemplates(path string) (*template.Template, *template.Template) {
	var home = template.Must(template.ParseFiles(path + "index.html"))
	var concerts = template.Must(template.ParseFiles(path + "concerts.html"))
	return home, concerts
}

func StartServer() {
	var homePage, concertsPage = loadTemplates("./templates/")
	homeEvent(homePage)
	concertsEvent(concertsPage)

	fmt.Println("URL: http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
