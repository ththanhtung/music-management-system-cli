package helpers

import (
	"fmt"
	"mms/models"
)

func (h *Helper) GetInputInt(msg string) int {
	for {
		fmt.Print(msg)
		var input int
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid integer.")
			// Clear the input buffer
			h.ClearConsole()
			continue
		}
		return input
	}
}

func (h *Helper) GetInputString(msg string) string {
	for {
		fmt.Print(msg)
		var input string
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid string.")
			// Clear the input buffer
			h.ClearConsole()
			continue
		}
		return input
	}
}
func(h *Helper) GetInputWithDefaultString(msg, defaultValue string) string {
	for {
		fmt.Print(msg)
		var input string
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid string.")
			// Clear the input buffer
			h.ClearConsole()
			continue
		}
		if input == "" {
			return defaultValue
		}
		return input
	}
}
func(h *Helper) GetInputWithDefaultInt(msg string, defaultValue int) int {
	for {
		fmt.Print(msg)
		var input int
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid string.")
			// Clear the input buffer
			h.ClearConsole()
			continue
		}
		if input == 0 {
			return defaultValue
		}
		return input
	}
}

func(h *Helper) GetMusicTrackDetail() *models.MusicTrackItem {
	track := models.NewMusicTrackItem(h)

	fmt.Println("\n---- Add a New Music Track ----")

	track.Title = h.GetInputString("Enter Title: ")

	track.Artist = h.GetInputString("Enter Artist: ")

	track.Album = h.GetInputString("Enter Album: ")

	track.Genre = h.GetInputString("Enter Genre: ")

	track.ReleaseYear = h.GetInputInt("Enter Release Year: ")

	track.Duration = h.GetInputInt("Enter Duration (in seconds): ")

	return track
}

func(h *Helper) GetUpdatedMusicTrackDetail(track *models.MusicTrackItem) models.MusicTrackItem {

	fmt.Println("\n---- Update Music Track Details ----")

	updatedTrack := *track

	newTitle := h.GetInputWithDefaultString("Enter new Title (leave blank to keep current value): ", track.Title)
	if newTitle != "" {
		updatedTrack.Title = newTitle
	}

	newArtist := h.GetInputWithDefaultString("Enter new Artist (leave blank to keep current value): ", track.Artist)
	if newArtist != "" {
		updatedTrack.Artist = newArtist
	}

	newAlbum := h.GetInputWithDefaultString("Enter new Album (leave blank to keep current value): ", track.Album)
	if newAlbum != "" {
		updatedTrack.Album = newAlbum
	}

	newGenre := h.GetInputWithDefaultString("Enter new Genre (leave blank to keep current value): ", track.Genre)
	if newGenre != "" {
		updatedTrack.Genre = newGenre
	}

	newReleaseYear := h.GetInputWithDefaultInt("Enter new Release Year (leave blank to keep current value): ", track.ReleaseYear)
	if newReleaseYear != -1 {
		updatedTrack.ReleaseYear = newReleaseYear
	}

	newDuration := h.GetInputWithDefaultInt("Enter new Duration (in seconds) (leave blank to keep current value): ", track.Duration)
	if newDuration != -1 {
		updatedTrack.Duration = newDuration
	}

	return updatedTrack
}

func(h *Helper) GetPlaylistDetail() *models.PlaylistItem{
	fmt.Println("\n---- Create a Playlist ----")
	playlist := models.NewPlaylistItem(h)

	playlist.Name = h.GetInputString("Enter Playlist Name: ")

	return playlist
}

func(h *Helper) SelectPlaylist(playlists *models.Playlists) *models.PlaylistItem{
	fmt.Println("\n---- Select a Playlist ----")
	for _, playlist := range playlists.Playlists {
		fmt.Println(playlist.ID, "-", playlist.Name)
	}

	choice := h.GetInputInt("Enter playlist ID: ")

	return playlists.GetPlaylistByID(fmt.Sprint(choice))
}

// func searchMusicTracksAndPlaylists() string{

// }