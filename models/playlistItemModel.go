package models

import "fmt"

type PlaylistItem struct {
	ID   string
	Name string
}

func NewPlaylistItem(helpers Helpers) *PlaylistItem {
	return &PlaylistItem{
		ID: helpers.RandomID(),
	}
}

func (p *PlaylistItem) DisplayInfo() {
	fmt.Println(p.ID, "-", p.Name)
}