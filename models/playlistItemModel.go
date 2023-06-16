package models

type PlaylistItem struct {
	ID   string
	Name string
}

func NewPlaylistItem(helpers Helpers) *PlaylistItem {
	return &PlaylistItem{
		ID: helpers.RandomID(),
	}
}
