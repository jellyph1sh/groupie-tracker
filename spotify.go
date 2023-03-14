package groupietracker

import (
	"context"
	"log"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func searchProfile(profileName string) *spotify.SearchResult {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     "04275833f211440d97072cbb65fd896b",
		ClientSecret: "d641de49fc534932a851798696dc92b4",
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	results, err := client.Search(ctx, profileName, spotify.SearchTypeArtist)
	if err != nil {
		log.Fatal(err)
	}
	return results
}
