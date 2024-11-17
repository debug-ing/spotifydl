package spotifydl

import (
	"context"
	"log"

	"github.com/debug-ing/spotifydl/src/config"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// UserData is a struct to hold all variables
type UserData struct {
	UserClient      *spotify.Client
	TrackList       []spotify.FullTrack
	SimpleTrackList []spotify.SimpleTrack
	YoutubeIDList   []string
}

// InitAuth starts Authentication
func InitAuth() *spotify.Client {
	ctx := context.Background()
	env := config.LoadConfig()
	if env.ClientID == "" || env.ClientSecret == "" {
		panic("No .env file found please config the .env file")
	}
	// Please do not misuse :)
	config := &clientcredentials.Config{
		ClientID:     env.ClientID,
		ClientSecret: env.ClientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	return client
}
