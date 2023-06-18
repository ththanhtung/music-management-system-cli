package models

import (
	"errors"
	"fmt"
	"mms/database"
	"strings"
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

func (p *Playlists) AddNewPlaylist(pl *PlaylistItem) error{
	if p.Playlists == nil {
		p.Playlists = make(map[string]*PlaylistItem)
	}

	// only create new if the playlist has not existed
	_, ok := p.Playlists[pl.ID]
	if !ok {
		// a user cannot create 2 or more playlist with the same name
		if !p.CheckDuplicatePlaylist(pl.Name){
			p.Playlists[pl.ID] = pl
			err := database.SaveMapToDB[map[string]*PlaylistItem](playlistURI, p.Playlists)
			if err != nil {
				fmt.Printf("Error saving music tracks to DB: %v\n", err)
			}
			return nil
		}
	}
	return errors.New("Duplicate playlist")
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

func (p *Playlists) CheckDuplicatePlaylist(playlistName string) bool{
	for _, playlist := range p.Playlists{
		if strings.EqualFold(strings.ToLower(playlist.Name), strings.ToLower(playlistName)){
			return true
		}
	}
	return false
}