package models


type PlaylistItem struct {
	ID     string
	Name   string
}

type Playlists struct {
	Playlists map[string]*PlaylistItem
}

func NewPlaylists()*Playlists{
	return &Playlists{
		Playlists: make(map[string]*PlaylistItem),
	}
}