package groupietracker

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func searchProfile(profileName string) *spotify.SearchResult {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     "",
		ClientSecret: "",
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		fmt.Printf("Couldn't get token: %v", err)
		return nil
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	results, err := client.Search(ctx, profileName, spotify.SearchTypeArtist)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return results
}
