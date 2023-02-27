package groupietracker

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	url      string
	fileName string
	API      string
}

func loadTemplate(pageName string) *template.Template {
	return template.Must(template.ParseFiles("./templates/" + pageName + ".html"))
}

func loadPage(mux *http.ServeMux, page *Page) {
	template := loadTemplate(page.fileName)
	mux.HandleFunc(page.url, func(w http.ResponseWriter, r *http.Request) {
		data := GetApi(page.API)
		template.Execute(w, data)
		r.ParseForm()
		if len(r.Form) != 0 {
			fmt.Println(r.Form["username"])
		}
	})
}

func loadEvents(mux *http.ServeMux, pages []Page) {
	for i := 0; i < len(pages); i++ {
		loadPage(mux, &pages[i])
	}
}

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	pages := []Page{{"/", "index", "topfive"}, {"/artists", "artists", "artists"}, {"/concerts", "concerts", "concerts"}}
	loadEvents(mux, pages)
	fmt.Println("URL: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
