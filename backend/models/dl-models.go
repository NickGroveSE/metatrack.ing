package models

import "time"

type DLHero struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"minimap_image_webp"`
}

type DLSHOCEntity struct {
	IDs      []int
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
