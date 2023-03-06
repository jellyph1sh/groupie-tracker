package groupietracker

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func loadTemplate(pageName string) *template.Template {
	return template.Must(template.ParseFiles("./templates/" + pageName + ".html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	template := loadTemplate("index")
	data := GetTopFive()
	template.Execute(w, data)
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	template := loadTemplate("artists")
	data := UnMarshallArtists(GetArtists())
	r.ParseForm()
	if len(r.Form) != 0 {
		// on tri
	}
	template.Execute(w, data)
}

func concertsHandler(w http.ResponseWriter, r *http.Request) {
	template := loadTemplate("concerts")
	data := GetConcerts()
	r.ParseForm()
	if len(r.Form) != 0 {
		// on tri
	}
	template.Execute(w, data)
}

/* {
	data := GetApi(page.API)
	r.ParseForm()
	if len(r.Form) != 0 {
		if r.FormValue("sort") != "" {
			data = GetSort(r.FormValue("sort"))
		} else if r.FormValue("pagination") != "" {
			pagination := r.FormValue("pagination")
			pagiInt := 0
			for _, char := range pagination {
				pagiInt *= 10
				pagiInt += int(byte(char) - 48)
			}
			for i := 0; i < pagiInt; i++ {
				var pagiData []any
				pagiData = append(pagiData, data[0:pagiInt])
			}
		} else {
			// définir un id de switch.
		}
	}
	template.Execute(w, data)
})*/

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts/"))))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/artists", artistsHandler)
	mux.HandleFunc("/concerts", concertsHandler)

	fmt.Println("URL: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
