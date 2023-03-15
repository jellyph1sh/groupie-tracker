package groupietracker

/*------------------------------------------------------*/
/*					Artists Structures :				*/
/*------------------------------------------------------*/
type Artist struct {
	Id           int `json:"id"`
	SpotifyID    string
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Artists []struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

/*------------------------------------------------------*/
/*					Concerts Structures :				*/
/*------------------------------------------------------*/

type Dates []struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	}
}

type Relation struct {
	Index []struct {
		ID             int                    `json:"id"`
		DatesLocations map[string]interface{} `json:"datesLocations"`
	} `json:"index"`
}

type Concerts struct {
	Artists  Artists
	Relation Relation
}

/*------------------------------------------------------*/
/*						Locations :						*/
/*------------------------------------------------------*/
type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	}
}
