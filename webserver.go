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

func loadTemplates(path string) *template.Template {
	var home = template.Must(template.ParseFiles(path + "index.html"))
	return home
}

func StartServer() {
	var homePage = loadTemplates("./templates/")
	homeEvent(homePage)

	fmt.Println("URL: http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
