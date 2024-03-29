package models

import (
	"errors"
	"fmt"
	"log"
	"mms/database"
	"strings"
)

type MusicTracks struct {
	Tracks map[string]*MusicTrackItem
}

var tracksURI string = "./database/tracks.json"

func NewMusicTracks(helpers Helpers) *MusicTracks {
	tracks := &MusicTracks{}
	tracksFromDB, err := database.ReadMapFromDB[MusicTrackItem](tracksURI)
	if err != nil {
		log.Fatalf(err.Error())
	}
	tracks.Tracks = tracksFromDB
	return tracks
}

func (m *MusicTracks) AddNewMusicTrack(mt *MusicTrackItem) error {
	if m.Tracks == nil {
		m.Tracks = make(map[string]*MusicTrackItem)
	}
	m.Tracks[mt.ID] = mt
	err := database.SaveMapToDB[map[string]*MusicTrackItem](tracksURI, m.Tracks)
	if err != nil {
		fmt.Printf("Error saving music tracks to DB: %v\n", err)
		return errors.New("error saving music tracks to DB")
	}
	return nil
}

func (m *MusicTracks) RemoveMusicTrack(mt *MusicTrackItem) error {
	if _, ok := m.Tracks[mt.ID]; ok {
		delete(m.Tracks, mt.ID)
	}
	err := database.SaveMapToDB[map[string]*MusicTrackItem](tracksURI, m.Tracks)
	if err != nil {
		fmt.Printf("Error saving music tracks to DB: %v\n", err)
		return errors.New("error saving music tracks to DB")
	}
	return nil
}

func (m *MusicTracks) UpdateMusicTrack(mt *MusicTrackItem, title, artist, album, genre string, releaseYear, duration int) error {
	if _, ok := m.Tracks[mt.ID]; ok {
		mt.Title = title
		mt.Artist = artist
		mt.Album = album
		mt.Genre = genre
		mt.ReleaseYear = releaseYear
		mt.Duration = duration
	}
	err := database.SaveMapToDB[map[string]*MusicTrackItem](tracksURI, m.Tracks)
	if err != nil {
		fmt.Printf("Error saving music tracks to DB: %v\n", err)
		return errors.New("error saving music tracks to DB")
	}
	return nil
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

func (m *MusicTracks) DisplayMinimumInfoAll() {
	fmt.Println("ID - Title - Artist")
	for _, track := range m.Tracks {
		track.DisplayMinimumInfo()
	}
}
