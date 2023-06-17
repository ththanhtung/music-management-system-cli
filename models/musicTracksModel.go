package models

import (
	"fmt"
	"mms/database"
	"strings"
)

type MusicTracks struct {
	Tracks map[string]*MusicTrackItem
}

func NewMusicTracks(helpers Helpers) *MusicTracks {
	tracks := &MusicTracks{}
	tracksFromDB, _ := database.ReadFromDB[MusicTrackItem]("./database/tracks.json")
	tracks.Tracks = tracksFromDB
	return tracks
}

func (m *MusicTracks) AddNewMusicTrack(mt *MusicTrackItem) {
	if m.Tracks == nil {
		m.Tracks = make(map[string]*MusicTrackItem)
	}
	m.Tracks[mt.ID] = mt
	err := database.SaveToDB[map[string]*MusicTrackItem]("./database/tracks.json", m.Tracks)
	if err != nil {
		fmt.Printf("Error saving music tracks to DB: %v\n", err)
	}
}

func (m *MusicTracks) RemoveMusicTrack(mt *MusicTrackItem) {
	if _, ok := m.Tracks[mt.ID]; ok {
		delete(m.Tracks, mt.ID)
	}
	err := database.SaveToDB[map[string]*MusicTrackItem]("./database/tracks.json", m.Tracks)
	if err != nil {
		fmt.Printf("Error saving music tracks to DB: %v\n", err)
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
	err := database.SaveToDB[map[string]*MusicTrackItem]("./database/tracks.json", m.Tracks)
	if err != nil {
		fmt.Printf("Error saving music tracks to DB: %v\n", err)
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

func (m *MusicTracks) loadMusicTracks(trackList []*MusicTrackItem) map[string]*MusicTrackItem {
	musicTracks := make(map[string]*MusicTrackItem, 0)
	for _, track := range trackList {
		musicTracks[track.ID] = track
	}
	return musicTracks
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

func (m *MusicTracks) SearchForTracksByArtist(artist string) []*MusicTrackItem {
	results := []*MusicTrackItem{}
	for _, track := range m.Tracks {
		if strings.Contains(strings.ToLower(track.Artist), strings.ToLower(artist)) {
			results = append(results, track)
		}
	}
	return results
}

func (m *MusicTracks) SearchForTracksByAlbum(album string) []*MusicTrackItem {
	results := []*MusicTrackItem{}
	for _, track := range m.Tracks {
		if strings.Contains(strings.ToLower(track.Album), strings.ToLower(album)) {
			results = append(results, track)
		}
	}
	return results
}

func (m *MusicTracks) SearchForTracksByGenry(genre string) []*MusicTrackItem {
	results := []*MusicTrackItem{}
	for _, track := range m.Tracks {
		if strings.Contains(strings.ToLower(track.Genre), strings.ToLower(genre)) {
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
