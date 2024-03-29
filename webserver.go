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
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	template := loadTemplate("index")
	data := GetTopFive()
	if data == nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	template.Execute(w, data)
}

func artistInfoHandler(w http.ResponseWriter, r *http.Request, artist Artist) {
	template := loadTemplate("artist-info")
	template.Execute(w, artist)
}

func artistsPage(mux *http.ServeMux) {
	pagiId := 0
	pagiInt := 0
	isSorted := false
	template := loadTemplate("artists")
	mux.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/artists" {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		data := GetArtistsData()
		if data == nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		r.ParseForm()
		if len(r.Form) != 0 {
			if r.FormValue("info") != "" {
				for i := 0; i < len(data); i++ {
					if r.FormValue("info") == data[i].Name {
						spotifyArtist := searchProfile(data[i].Name)
						if spotifyArtist == nil {
							errorHandler(w, r, http.StatusInternalServerError)
							return
						}
						artist := Artist{
							Id:           data[i].Id,
							SpotifyID:    spotifyArtist.Artists.Artists[0].ID.String(),
							Image:        data[i].Image,
							Name:         data[i].Name,
							Members:      data[i].Members,
							CreationDate: data[i].CreationDate,
							FirstAlbum:   data[i].FirstAlbum,
							Locations:    data[i].Locations,
							ConcertDates: data[i].ConcertDates,
							Relations:    data[i].Relations,
						}
						artistInfoHandler(w, r, artist)
						return
					}
				}
			}
			if r.FormValue("sort") != "" || isSorted {
				data = GetSort(r.FormValue("sort"), data)
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
		if r.URL.Path != "/concerts" {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		data := GetConcerts()
		if data.Artists == nil || data.Relation.Index == nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		r.ParseForm()
		if len(r.Form) != 0 {
			if r.FormValue("search") != "" {
				searchRes := GetSearch(r.FormValue("search"))
				if len(searchRes.Artists) != 0 {
					data = searchRes
				}
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

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	if status == http.StatusNotFound {
		template := loadTemplate("error404")
		template.Execute(w, nil)
	} else if status == http.StatusInternalServerError {
		template := loadTemplate("error500")
		template.Execute(w, nil)
	}
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
