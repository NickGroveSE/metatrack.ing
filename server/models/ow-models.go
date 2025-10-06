package models

import (
	"time"
)

type OWHero struct {
	Name     string
	Color    string
	Image    string
	PickRate float32
	WinRate  float32
}

type OWDataResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Data      []OWHero  `json:"data"`
}
