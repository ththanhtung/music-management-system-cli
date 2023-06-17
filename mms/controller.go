package mms

import (
	"fmt"
	"mms/helpers"
	"mms/models"
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
			track := helpers.GetMusicTrackDetail()
			library.MusicTracks.AddNewMusicTrack(track)
			helpers.ClearConsole()
			fmt.Println("Music track added successfully!")
		case 2:
			search := helpers.GetInputString("Enter music track name: ")
			track := library.MusicTracks.GetTrackByName(search)
			helpers.ClearConsole()
			if track == nil {
				fmt.Println("No music track found!")
			} else {
				track.DisplayInfo()
			}
		case 3:
			search := helpers.GetInputString("Enter music track name: ")
			track := library.MusicTracks.GetTrackByName(search)
			helpers.ClearConsole()
			if track == nil {
				fmt.Println("No music track found!")
			} else {
				updatedTrack := helpers.GetUpdatedMusicTrackDetail(track)
				library.MusicTracks.UpdateMusicTrack(track, updatedTrack.Title,updatedTrack.Artist,updatedTrack.Album,updatedTrack.Genre,updatedTrack.ReleaseYear, updatedTrack.Duration)
				fmt.Println("Music track updated successfully!")
			}
		case 4:
			search := helpers.GetInputString("Enter music track name: ")
			track := library.MusicTracks.GetTrackByName(search)
			helpers.ClearConsole()
			if track == nil {
				fmt.Println("No music track found!")
			} else {
				library.MusicTracks.RemoveMusicTrack(track)
				fmt.Println("Music track deleted successfully!")
			}
		case 5:
			helpers.ClearConsole()
			library.MusicTracks.DisplayAll()
		case 6:

			playlist := helpers.GetPlaylistDetail()
			library.Playlists.AddNewPlaylist(playlist)
			helpers.ClearConsole()
			fmt.Println("Playlist created successfully!", playlist.ID)
		case 7:
			search := helpers.GetInputString("Enter music track you want to add: ")
			track := library.MusicTracks.GetTrackByName(search)
			helpers.ClearConsole()
			if track == nil {
				fmt.Println("No music track found!")
			} else {
				playlist := helpers.SelectPlaylist(library.Playlists)
				if playlist == nil {
					fmt.Println("No playlist found!")
				} else {
					library.AddTrackToPlaylist(track.Title, playlist.Name)
					fmt.Println("Track added to playlist successfully!")
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