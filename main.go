package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	spotifydl "github.com/debug-ing/spotifydl/src"

	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()
	var trackID string
	var playlistID string
	var albumID string
	var spotifyURL string

	var rootCmd = &cobra.Command{
		Use:   "spotifydl",
		Short: "spotifydl is a awesome music downloader!",
		Long:  `spotifydl lets you download albums and playlists and tags them for you.`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				_ = cmd.Help()
			}

			spotifyURL = args[0]

			if len(spotifyURL) == 0 {
				fmt.Println("=> Spotify URL required.")
				_ = cmd.Help()
				return
			}

			splitURL := strings.Split(spotifyURL, "/")

			if len(splitURL) < 2 {
				fmt.Println("=> Please enter the url copied from the spotify client.")
				os.Exit(1)
			}

			spotifyID := splitURL[len(splitURL)-1]
			if strings.Contains(spotifyID, "?") {
				spotifyID = strings.Split(spotifyID, "?")[0]
			}

			if strings.Contains(spotifyURL, "album") {
				albumID = spotifyID
				spotifydl.DownloadAlbum(ctx, albumID)
			} else if strings.Contains(spotifyURL, "playlist") {
				playlistID = spotifyID
				spotifydl.DownloadPlaylist(ctx, playlistID)
			} else if strings.Contains(spotifyURL, "track") {
				trackID = spotifyID
				spotifydl.DownloadSong(ctx, trackID)
			} else {
				fmt.Println("=> Only Spotify Album/Playlist/Track URL's are supported.")
				_ = cmd.Help()
			}

		},
	}

	rootCmd.SetUsageTemplate(`spotifydl [spotify_url]`)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
