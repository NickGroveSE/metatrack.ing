package models

import "time"

type DLHero struct {
	ID    int32
	Name  string
	Image string
	Color string
}

type DLHeroStats struct {
	Wins     int64
	Matches  int64
	PickRate float32
	WinRate  float32
}

type DLHeroEntity struct {
	Heroes []DLHero
	Stats  DLHeroStats
}

type DLDataResponse struct {
	Status    string         `json:"status"`
	Timestamp time.Time      `json:"timestamp"`
	Data      []DLHeroEntity `json:"data"`
}

type DLIndividualComboDataResponse struct {
	Status    string       `json:"status"`
	Timestamp time.Time    `json:"timestamp"`
	Data      DLHeroEntity `json:"data"`
}
