package models

import (
	"fmt"
	"mms/database"
)

type Playlists struct {
	Playlists map[string]*PlaylistItem
}

var playlistURI string = "./database/playlists.json"

func NewPlaylists(helpers Helpers) *Playlists {
	playlists := &Playlists{}
	playlistsFromDB, _ := database.ReadMapFromDB[PlaylistItem](playlistURI)
	playlists.Playlists = playlistsFromDB
	return playlists
}

func (p *Playlists) AddNewPlaylist(pl *PlaylistItem) {
	if p.Playlists == nil {
		p.Playlists = make(map[string]*PlaylistItem)
	}
	// only create new if the playlist has not existed
	_, ok := p.Playlists[pl.ID]
	if !ok {
		p.Playlists[pl.ID] = pl
		err := database.SaveMapToDB[map[string]*PlaylistItem](playlistURI, p.Playlists)
		if err != nil {
			fmt.Printf("Error saving music tracks to DB: %v\n", err)
		}
	}

}

func (p *Playlists) RemovePlaylist(pl *PlaylistItem) {
	if _, ok := p.Playlists[pl.ID]; ok {
		delete(p.Playlists, pl.ID)
	}
	err := database.SaveMapToDB[map[string]*PlaylistItem](playlistURI, p.Playlists)
	if err != nil {
		fmt.Printf("Error saving music tracks to DB: %v\n", err)
	}
}

func (p *Playlists) GetPlaylistByName(playlistName string) *PlaylistItem {
	for _, playlist := range p.Playlists {
		if playlist.Name == playlistName {
			return playlist
		}
	}
	return nil
}

func (p *Playlists) GetPlaylistByID(id string) *PlaylistItem {
	for _, playlist := range p.Playlists {
		if playlist.ID == id {
			return playlist
		}
	}
	return nil
}

func (p *Playlists) DisplayAll(){
	for _, playlist := range p.Playlists {
		playlist.DisplayInfo()
	}
}