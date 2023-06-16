package models

type Playlists struct {
	Playlists map[string]*PlaylistItem
}

func NewPlaylists(helpers Helpers) *Playlists {
	return &Playlists{
		Playlists: make(map[string]*PlaylistItem),
	}
}

func (p *Playlists) AddNewPlaylist(pl *PlaylistItem) {
	p.Playlists[pl.ID] = pl
}

func (p *Playlists) RemovePlaylist(pl *PlaylistItem) {
	if _, ok := p.Playlists[pl.ID]; ok {
		delete(p.Playlists, pl.ID)
	}
}

func (p *Playlists) GetPlaylistByName(playlistName string) *PlaylistItem{
	for _, playlist := range p.Playlists{
		if playlist.Name == playlistName{
			return playlist
		}
	}
	return nil
}

func (p *Playlists) GetPlaylistByID(id string) *PlaylistItem{
	for _, playlist := range p.Playlists{
		if playlist.ID == id{
			return playlist
		}
	}
	return nil
}