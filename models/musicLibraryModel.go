package models

import (
	"fmt"
	"strings"
)

type MusicLibraryItem struct {
	PlaylistID string
	TrackID    string
}

type MusicLibrary struct {
	MusicTracks    *MusicTracks
	Playlists      *Playlists
	MusicLibraries []*MusicLibraryItem
}

type Helpers interface {
	RandomID() string
}

func NewMusicLibrary(helpers Helpers) *MusicLibrary {
	return &MusicLibrary{
		MusicTracks:    NewMusicTracks(helpers),
		Playlists:      NewPlaylists(helpers),
		MusicLibraries: make([]*MusicLibraryItem, 0),
	}
}

func (ml *MusicLibrary) AddTrackToPlaylist(trackName, playlistName string) {
	track := ml.MusicTracks.GetTrackByName(trackName)
	playlist := ml.Playlists.GetPlaylistByName(playlistName)
	ml.MusicLibraries = append(ml.MusicLibraries, &MusicLibraryItem{PlaylistID: playlist.ID, TrackID: track.ID})
}

func (ml *MusicLibrary) GetAllTrackFromPlaylist(playlistName string) []*MusicTrackItem {
	playlist := ml.Playlists.GetPlaylistByName(playlistName)
	tracks := []*MusicTrackItem{}
	for _, item := range ml.MusicLibraries {
		if item.PlaylistID == playlist.ID {
			track := ml.MusicTracks.GetTrackByID(item.TrackID)
			tracks = append(tracks, track)
		}
	}
	return tracks
}

func (ml *MusicLibrary) DisplayAllTrackFromPlaylist(playlistName string) {
	playlist := ml.Playlists.GetPlaylistByName(playlistName)
	fmt.Println("Playlist:", playlist.Name)
	fmt.Println("Tracks:")
	tracks := ml.GetAllTrackFromPlaylist(playlistName)
	for i, track := range tracks {
		fmt.Println("Track", i+1)
		track.DisplayInfo()
	}
}

func (ml *MusicLibrary) SearchTracksAndPlaylistsByTitle(searchTerm string) map[string][]interface{} {
	results := make(map[string][]interface{})

	// Search for tracks
	for _, track := range ml.MusicTracks.Tracks {
	    if strings.Contains(strings.ToLower(track.Title), strings.ToLower(searchTerm)) {
	        results["Tracks"] = append(results["Tracks"], track)
	    }
	}

	// Search for playlists that contain the tracks
	for _, libraryItem := range ml.MusicLibraries {
		if _, ok := results["Tracks"]; ok {
			// Check if this library item's track ID is in the results
			for _, track := range results["Tracks"] {
				if track.(*MusicTrackItem).ID == libraryItem.TrackID {
					// Get the playlist from the playlists map
					playlist := ml.Playlists.GetPlaylistByID(libraryItem.PlaylistID)
					results["Playlists"] = append(results["Playlists"], playlist)
					results["TracksInPlaylists"] = append(results["TracksInPlaylists"], track)
				}
			}
		}
	}

	return results
}

func (ml *MusicLibrary) DisplayTrackAndLibraryByTitle(searchTerm string) {
	results := ml.SearchTracksAndPlaylistsByTitle(searchTerm)
	fmt.Println("Tracks:")
	for _, track := range results["Tracks"] {
		fmt.Println(track.(*MusicTrackItem).Title)
	}
	fmt.Println("Playlists:")
	for i, playlist := range results["Playlists"] {
		fmt.Println("-", playlist.(*PlaylistItem).Name)
		tracksInPlaylist := results["TracksInPlaylists"][i].(*MusicTrackItem)
		fmt.Println("  -", tracksInPlaylist.Title)
	}
}