package helpers

import (
	"fmt"
	"mms/models"
)

func (h *Helper) GetMusicTrackDetail() *models.MusicTrackItem {
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

func (h *Helper) GetUpdatedMusicTrackDetail(track *models.MusicTrackItem) models.MusicTrackItem {

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

func (h *Helper) GetPlaylistDetail() *models.PlaylistItem {
	fmt.Println("\n---- Create a Playlist ----")
	playlist := models.NewPlaylistItem(h)

	playlist.Name = h.GetInputString("Enter Playlist Name: ")

	return playlist
}

func (h *Helper) SelectPlaylist(playlists *models.Playlists) *models.PlaylistItem {
	fmt.Println("\n---- Select a Playlist ----")
	for _, playlist := range playlists.Playlists {
		fmt.Println(playlist.ID, "-", playlist.Name)
	}

	choice := h.GetInputInt("Enter playlist ID: ")

	return playlists.GetPlaylistByID(fmt.Sprint(choice))
}

func (h *Helper) SearchForTrack() {
	fmt.Println("\n---- Search For Track ----")
	fmt.Println("1. Search by title")
	fmt.Println("2. Search by artist")
	fmt.Println("3. Search by album")
	fmt.Println("4. Search by genre")

	choice := h.GetInputInt("Enter your choice")

	if choice > 0 && choice <= 4 {
		
	}
}
