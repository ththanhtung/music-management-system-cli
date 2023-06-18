package models

import "fmt"

type MusicTrackItem struct {
	ID          string
	Title       string
	Artist      string
	Album       string
	Genre       string
	ReleaseYear int
	Duration    int
}

func NewMusicTrackItem(helpers Helpers) *MusicTrackItem {
	return &MusicTrackItem{
		ID: helpers.RandomID(),
	}
}

func (m *MusicTrackItem) DisplayInfo() {
	fmt.Println("\n---- Music Track Details ----")
	fmt.Println("ID:", m.ID)
	fmt.Println("Title:", m.Title)
	fmt.Println("Artist:", m.Artist)
	fmt.Println("Album:", m.Album)
	fmt.Println("Genre:", m.Genre)
	fmt.Println("Release Year:", m.ReleaseYear)
	fmt.Println("Duration:", m.Duration, "seconds")
}

func (m *MusicTrackItem) DisplayMinimumInfo() {
	fmt.Println(m.ID + " - " + m.Title + " - " + m.Artist)
}
