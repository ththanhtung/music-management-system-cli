package models

import (
	"fmt"
	"mms/database"
	"strings"
)

type MusicLibrary struct {
	MusicTracks    *MusicTracks
	Playlists      *Playlists
	MusicLibraries []*MusicLibraryItem
}

type Helpers interface {
	RandomID() string
}

var librariesURI string = "./database/libraries.json"

func NewMusicLibrary(helpers Helpers) *MusicLibrary {
	musicLibraby := &MusicLibrary{}
	musicLibraby.MusicTracks = NewMusicTracks(helpers)
	musicLibraby.Playlists = NewPlaylists(helpers)
	librariesFromDB,_:=database.ReadSliceFromDB[MusicLibraryItem](librariesURI)
	musicLibraby.MusicLibraries = librariesFromDB
	return musicLibraby
}

func (ml *MusicLibrary) AddTrackToPlaylist(trackName, playlistName string) {
	track := ml.MusicTracks.GetTrackByName(trackName)
	playlist := ml.Playlists.GetPlaylistByName(playlistName)
	ml.MusicLibraries = append(ml.MusicLibraries, &MusicLibraryItem{PlaylistID: playlist.ID, TrackID: track.ID})

	// assign like this to avoid nil pointer
	if ml.MusicLibraries == nil {
		ml.MusicLibraries = make([]*MusicLibraryItem, 0)
	}
	database.SaveSliceToDB[MusicLibraryItem](librariesURI, ml.MusicLibraries)
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

func (ml *MusicLibrary) SearchTracksAndPlaylistsByArtist(searchTerm string) map[string][]interface{} {
	results := make(map[string][]interface{})

	// Search for tracks
	for _, track := range ml.MusicTracks.Tracks {
	    if strings.Contains(strings.ToLower(track.Artist), strings.ToLower(searchTerm)) {
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

func (ml *MusicLibrary) SearchTracksAndPlaylistsByAlbum(searchTerm string) map[string][]interface{} {
	results := make(map[string][]interface{})

	// Search for tracks
	for _, track := range ml.MusicTracks.Tracks {
	    if strings.Contains(strings.ToLower(track.Album), strings.ToLower(searchTerm)) {
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

func (ml *MusicLibrary) SearchTracksAndPlaylistsByGenre(searchTerm string) map[string][]interface{} {
	results := make(map[string][]interface{})

	// Search for tracks
	for _, track := range ml.MusicTracks.Tracks {
	    if strings.Contains(strings.ToLower(track.Genre), strings.ToLower(searchTerm)) {
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
					_, ok := results["Playlists"]
					if !ok {
						results["Playlists"] = append(results["Playlists"], playlist)
					}
					results["TracksInPlaylists"] = append(results["TracksInPlaylists"], track)
				}
			}
		}
	}

	return results
}

func (ml *MusicLibrary) DisplayTrackAndLibrary(results map[string][]interface{}) {
	fmt.Println("Tracks:")
	for _, track := range results["Tracks"] {
		fmt.Println("-",track.(*MusicTrackItem).Title)
	}
	fmt.Println("Playlists:")
	for i, playlist := range results["Playlists"] {
		fmt.Println("-", playlist.(*PlaylistItem).Name)
		tracksInPlaylist := results["TracksInPlaylists"][i].(*MusicTrackItem)
		fmt.Println("  -", tracksInPlaylist.Title)
	}
}