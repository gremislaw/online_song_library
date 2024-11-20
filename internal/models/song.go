package models

import "time"

// @Description Song model
type Song struct {
	// ID of the song
	ID uint `gorm:"primaryKey" json:"id"`
	// Group name
	Group string `json:"group"`
	// Song title
	Title string `json:"song" gorm:"column:song"`
	// Song text
	Text string `json:"text"`
	// Link to the song
	Link string `json:"link"`
	// Release date of the song
	ReleaseDate string `json:"releaseDate"`
	// Created date of the song
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	// Updated date of the song
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// @Description SongDetail model
type SongDetail struct {
	// Release date of the song
	Text string `json:"text"`
	// Text of the song
	Link string `json:"link"`
	// Link to the song
	ReleaseDate string `json:"releaseDate"`
}
