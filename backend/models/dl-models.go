package models

import "time"

type DLHero struct {
	ID    string
	Name  string
	Image string
}

type DLSHOCEntity struct {
	IDs      []string
	Names    []string
	Color    []string
	Images   []string
	PickRate float32
	WinRate  float32
}

type DLSHOCResponse struct {
	Status    string         `json:"status"`
	Timestamp time.Time      `json:"timestamp"`
	Data      []DLSHOCEntity `json:"data"`
}
