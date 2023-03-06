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

func artistsPage(mux *http.ServeMux) {
	pagiId := 0
	pagiInt := 0
	isSorted := false
	template := loadTemplate("artists")
	mux.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		data := UnMarshallArtists(GetArtists())
		r.ParseForm()
		if len(r.Form) != 0 {
			if r.FormValue("sort") != "" || isSorted {
				data = GetSort(r.FormValue("sort"))
				isSorted = true
			}
			if r.FormValue("pagination") != "" {
				pagination := r.FormValue("pagination")
				pagiId = 0
				pagiInt = 0
				for _, char := range pagination {
					pagiInt *= 10
					pagiInt += int(byte(char) - 48)
				}
				for i := 0; i < pagiInt; i++ {
					data = data[pagiId:pagiInt]
				}
			} else if r.FormValue("switch") != "" {
				if r.FormValue("switch") == "next" {
					if pagiId+pagiInt < 52 {
						pagiId += pagiInt
					}
				} else {
					if pagiId-pagiInt >= 0 {
						pagiId -= pagiInt
					}
				}
				data = data[pagiId : pagiId+pagiInt]
			}
		} else {
			isSorted = false
		}
		template.Execute(w, data)
	})
}

func concertsHandler(w http.ResponseWriter, r *http.Request) {
	template := loadTemplate("concerts")
	data := GetConcerts()
	r.ParseForm()
	template.Execute(w, data)
}

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts/"))))

	artistsPage(mux)
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/concerts", concertsHandler)

	fmt.Println("URL: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
