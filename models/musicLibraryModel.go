package models


type MusicLibrary struct {
	Tracks    *MusicTracks
	Playlists *Playlists
}

func NewMusicLibrary()*MusicLibrary{
	return &MusicLibrary{
		Tracks: NewMusicTracks(),
		Playlists: NewPlaylists(),
	}
}
