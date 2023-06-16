package models

import "fmt"

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
	RandomID()string
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