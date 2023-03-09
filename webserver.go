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

func concertsPage(mux *http.ServeMux) {
	pagiId := 0
	pagiInt := 0
	template := loadTemplate("concerts")
	mux.HandleFunc("/concerts", func(w http.ResponseWriter, r *http.Request) {
		data := GetConcerts()
		r.ParseForm()
		if len(r.Form) != 0 {
			if r.FormValue("search") != "" {
				fmt.Println(r.FormValue("search"))
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
					data.Artists = data.Artists[pagiId:pagiInt]
					data.Relation.Index = data.Relation.Index[pagiId:pagiInt]
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
				data.Artists = data.Artists[pagiId : pagiId+pagiInt]
				data.Relation.Index = data.Relation.Index[pagiId : pagiId+pagiInt]
			}
		}
		template.Execute(w, data)
	})
}

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts/"))))

	mux.HandleFunc("/", homeHandler)
	artistsPage(mux)
	concertsPage(mux)

	fmt.Println("URL: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
