package models

import "strings"

type MusicTracks struct {
	Tracks map[string]*MusicTrackItem
}

func NewMusicTracks(helpers Helpers) *MusicTracks {
	return &MusicTracks{
		Tracks: make(map[string]*MusicTrackItem),
	}
}

func (m *MusicTracks) AddNewMusicTrack(mt *MusicTrackItem) {
	m.Tracks[mt.ID] = mt
}

func (m *MusicTracks) RemoveMusicTrack(mt *MusicTrackItem) {
	if _, ok := m.Tracks[mt.ID]; ok {
		delete(m.Tracks, mt.ID)
	}
}

func (m *MusicTracks) UpdateMusicTrack(mt *MusicTrackItem, title, artist, album, genre string, releaseYear, duration int) {
	if _, ok := m.Tracks[mt.ID]; ok {
		mt.Title = title
		mt.Artist = artist
		mt.Album = album
		mt.Genre = genre
		mt.ReleaseYear = releaseYear
		mt.Duration = duration
	}
}

func (m *MusicTracks) GetTrackByName(trackName string) *MusicTrackItem {
	for _, track := range m.Tracks {
		if track.Title == trackName {
			return track
		}
	}
	return nil
}

func (m *MusicTracks) GetTrackByID(id string) *MusicTrackItem {
	for _, track := range m.Tracks {
		if track.ID == id {
			return track
		}
	}
	return nil
}

func (m *MusicTracks) SearchForTracksByTitle(title string) []*MusicTrackItem {
	results := []*MusicTrackItem{}
	for _, track := range m.Tracks {
		if strings.Contains(strings.ToLower(track.Title), strings.ToLower(title)) {
			results = append(results, track)
		}
	}
	return results
}

func (m *MusicTracks) DisplayAll() {
	for _, track := range m.Tracks {
		track.DisplayInfo()
	}
}