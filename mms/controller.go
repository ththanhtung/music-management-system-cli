package mms

import (
	"fmt"
	"mms/helpers"
	"mms/models"
	"strings"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Init() {
	c.loop()
}

func (c *Controller) loop() {
	helpers := helpers.NewHelper()

	helpers.ClearConsole()

	library := models.NewMusicLibrary(helpers)

	for {
		fmt.Println("\n---- Music Library Management System ----")
		fmt.Println("1. Add a new music track")
		fmt.Println("2. View details of a music track")
		fmt.Println("3. Update details of a music track")
		fmt.Println("4. Delete a music track")
		fmt.Println("5. Display all music tracks")
		fmt.Println("6. Create a playlist")
		fmt.Println("7. Add a track to a playlist")
		fmt.Println("8. Search music tracks and playlists")
		fmt.Println("9. Exit")

		choice := helpers.GetInputInt("Enter your choice: ")

		helpers.ClearConsole()

		switch choice {
		case 1:
			// Display form to take track info
			track := helpers.GetMusicTrackDetail()

			// Add track to library
			if err := library.MusicTracks.AddNewMusicTrack(track); err != nil {
				helpers.ClearConsole()
				fmt.Println(err.Error())
			} else {
				helpers.ClearConsole()
				track.DisplayInfo()
				fmt.Println("Music track added successfully!")
			}

		case 2:
			// Display form to take track id
			search := helpers.GetInputString("Enter music track ID: ")

			// Get track
			track := library.MusicTracks.GetTrackByID(search)
			helpers.ClearConsole()
			if track == nil {
				fmt.Println("No music track found!")
			} else {
				// Display track info
				track.DisplayInfo()
			}
		case 3:
			// Display form to take track id
			search := helpers.GetInputString("Enter music track ID: ")

			// Get track
			track := library.MusicTracks.GetTrackByID(search)
			helpers.ClearConsole()
			if track == nil {
				fmt.Println("No music track found!")
			} else {
				// Display form to take updated track info
				updatedTrack := helpers.GetUpdatedMusicTrackDetail(track)

				// Update track info
				if err := library.MusicTracks.UpdateMusicTrack(track, updatedTrack.Title, updatedTrack.Artist, updatedTrack.Album, updatedTrack.Genre, updatedTrack.ReleaseYear, updatedTrack.Duration); err != nil {
					helpers.ClearConsole()
					fmt.Println(err.Error())
				} else {
					helpers.ClearConsole()
					fmt.Println("Music track updated successfully!")
				}
			}
		case 4:
			// Display form to take track id
			search := helpers.GetInputString("Enter music track ID: ")

			// Get track
			track := library.MusicTracks.GetTrackByID(search)
			helpers.ClearConsole()
			if track == nil {
				fmt.Println("No music track found!")
			} else {
				// Remove track
				if err := library.MusicTracks.RemoveMusicTrack(track); err != nil {
					helpers.ClearConsole()
					fmt.Println(err.Error())
				} else {
					helpers.ClearConsole()
					fmt.Println("Music track deleted successfully!")
				}
			}
		case 5:
			helpers.ClearConsole()
			library.MusicTracks.DisplayAll()
		case 6:
			// Display from to take playlist info
			playlist := helpers.GetPlaylistDetail()

			// Add playlist 
			if err := library.Playlists.AddNewPlaylist(playlist); err != nil {
				helpers.ClearConsole()
				fmt.Println(err.Error())
			} else {
				helpers.ClearConsole()
				playlist.DisplayInfo()
				fmt.Println("Playlist created successfully!")
			}
		case 7:
			// Display some minimum info (id - title - artist) so that the user can know the id of the track they want to add to a playlist
			library.MusicTracks.DisplayMinimumInfoAll()
			search := helpers.GetInputInt("Enter music track ID you want to add: ")
			helpers.ClearConsole()

			// Get track
			track := library.MusicTracks.GetTrackByID(fmt.Sprint(search))
			if track == nil {
				fmt.Println("No music track found!")
			} else {
				playlist := helpers.SelectPlaylist(library.Playlists)
				if playlist == nil {
					helpers.ClearConsole()
					fmt.Println("No playlist found!")
				} else {
					// Add track to playlist
					if err := library.AddTrackToPlaylist(track.Title, playlist.Name); err != nil {
						helpers.ClearConsole()
						fmt.Println(err.Error())
					} else {
						helpers.ClearConsole()
						fmt.Println(strings.ToUpper(track.Title) + " added to " + strings.ToUpper(playlist.Name) + " playlist successfully!")
					}
				}
			}
		case 8:
			helpers.ClearConsole()
			helpers.SearchForTrackAndPlaylist(library)
		case 9:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice! Please enter a valid option.")
		}
	}
}
