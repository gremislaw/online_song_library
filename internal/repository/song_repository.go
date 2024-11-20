package repository

import (
	"online_song_library/internal/models"
)

type SongRepository interface {
	GetPaginated() ([]*models.Song, error)
	Create(Song *models.Song) error
	Update(Song *models.Song) error
	Delete(id int) error
}
