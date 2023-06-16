package models


type MusicTrackItem struct {
	ID          string
	Title       string
	Artist      string
	Album       string
	Genre       string
	ReleaseYear int
	Duration    int
}

type MusicTracks struct {
	Tracks map[string]*MusicTrackItem
}

func NewMusicTracks()*MusicTracks{
	return &MusicTracks{
		Tracks: make(map[string]*MusicTrackItem),
	}
}