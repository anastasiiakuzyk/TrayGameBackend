package game

type Map struct {
	Players *Players `json:"players"`
	Width uint	`json:"width"`
	Height uint `json:"height"`
	Helpers Helpers
}

type SyncMap struct {
	Players []*Player `json:"players"`
	Width uint	`json:"width"`
	Height uint `json:"height"`
	Helpers []Helper `json:"helpers"`
}

func (m *Map) ToSyncMap() SyncMap {
	return SyncMap{
		Players: m.Players.ToArray(),
		Width: m.Width,
		Height: m.Height,
		Helpers: m.Helpers.ToArray(),
	}
}

func NewDefaultMap() *Map {
	return &Map{
		Players: NewPlayers(),
		Width:   1280,
		Height:  720,
	}
}