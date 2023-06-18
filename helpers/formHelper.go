package helpers

import (
	"fmt"
	"mms/models"
	"time"
)

func (h *Helper) GetMusicTrackDetail() *models.MusicTrackItem {
	track := models.NewMusicTrackItem(h)

	fmt.Println("\n---- Add a New Music Track ----")

	track.Title = h.GetInputStringWithConstraint("Enter Title (leave blank for default value): ", "default title", 1, 100)

	track.Artist = h.GetInputStringWithConstraint("Enter Artist (leave blank for default value): ", "default artist", 2, 100)

	track.Album = h.GetInputStringWithConstraint("Enter Album (leave blank for default value): ", "default album", 1, 100)

	track.Genre = h.GetInputStringWithConstraint("Enter Genre (leave blank for default value): ", "default genre", 1, 100)

	track.ReleaseYear = h.GetInputIntWithConstraint("Enter Release Year(1900-current year) (enter -1 for default value): ", time.Now().Year(), 1900, time.Now().Year())

	track.Duration = h.GetInputIntWithConstraint("Enter Duration (in seconds) (enter -1 for default value): ", 0, 0, 10000)

	return track
}

func (h *Helper) GetUpdatedMusicTrackDetail(track *models.MusicTrackItem) models.MusicTrackItem {

	fmt.Println("\n---- Update Music Track Details ----")

	updatedTrack := *track

	newTitle := h.GetInputStringWithConstraint("Enter new Title (leave blank to keep current value): ", track.Title, 1, 100)
	if newTitle != "" {
		updatedTrack.Title = newTitle
	}

	newArtist := h.GetInputStringWithConstraint("Enter new Artist (leave blank to keep current value): ", track.Artist, 2, 100)
	if newArtist != "" {
		updatedTrack.Artist = newArtist
	}

	newAlbum := h.GetInputStringWithConstraint("Enter new Album (leave blank to keep current value): ", track.Album, 1, 100)
	if newAlbum != "" {
		updatedTrack.Album = newAlbum
	}

	newGenre := h.GetInputStringWithConstraint("Enter new Genre (leave blank to keep current value): ", track.Genre, 1, 100)
	if newGenre != "" {
		updatedTrack.Genre = newGenre
	}

	newReleaseYear := h.GetInputIntWithConstraint("Enter new Release Year (type -1 to keep current value): ", track.ReleaseYear, 1900, time.Now().Year())
	if newReleaseYear != -1 {
		updatedTrack.ReleaseYear = newReleaseYear
	}

	newDuration := h.GetInputIntWithConstraint("Enter new Duration (in seconds) (type -1 to keep current value): ", track.Duration, 0, 10000)
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
	playlists.DisplayAll()

	choice := h.GetInputInt("Enter playlist ID: ")

	return playlists.GetPlaylistByID(fmt.Sprint(choice))
}

func (h *Helper) SearchForTrackAndPlaylist(ml *models.MusicLibrary) {
	fmt.Println("\n---- Search For Track and playlist ----")
	fmt.Println("1. Search by title")
	fmt.Println("2. Search by artist")
	fmt.Println("3. Search by album")
	fmt.Println("4. Search by genre")
	fmt.Println("5. Go back")

	choice := h.GetInputInt("Enter your choice: ")
	h.ClearConsole()

	switch choice {
	case 1:
		//set "default value"+fmt.Sprint(time.Now()) as default value to avoid unwanted match
		title := h.GetInputStringWithConstraint("Enter track's title: ", "default value"+fmt.Sprint(time.Now()), 2, 100)
		result := ml.SearchTracksAndPlaylistsByTitle(title)
		ml.DisplayTrackAndLibrary(result)
	case 2:
		//set "default value"+fmt.Sprint(time.Now()) as default value to avoid unwanted match
		title := h.GetInputStringWithConstraint("Enter artist: ", "default value"+fmt.Sprint(time.Now()), 1, 100)
		result := ml.SearchTracksAndPlaylistsByArtist(title)
		ml.DisplayTrackAndLibrary(result)
	case 3:
		//set "default value"+fmt.Sprint(time.Now()) as default value to avoid unwanted match
		title := h.GetInputStringWithConstraint("Enter album: ", "default value"+fmt.Sprint(time.Now()), 1, 100)
		result := ml.SearchTracksAndPlaylistsByAlbum(title)
		ml.DisplayTrackAndLibrary(result)
	case 4:
		//set "default value"+fmt.Sprint(time.Now()) as default value to avoid unwanted match
		title := h.GetInputStringWithConstraint("Enter genre: ", "default value"+fmt.Sprint(time.Now()), 1, 100)
		result := ml.SearchTracksAndPlaylistsByGenre(title)
		ml.DisplayTrackAndLibrary(result)
	case 5:
		h.ClearConsole()
		return
	}
}
