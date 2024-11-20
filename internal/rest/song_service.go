package rest

import (
	"errors"
	"online_song_library/internal/models"
	"online_song_library/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SongService struct {
	DB *gorm.DB
}

func NewSongService(db *gorm.DB) repository.SongRepository {
	return &SongService{DB: db}
}

func (s *SongService) GetPaginated() ([]*models.Song, error) {
	logrus.Debug("Get songs DB in GetPaginated")
	var Songs []*models.Song

	tx := s.DB.Find(&Songs)
	if tx.Error != nil {
		logrus.Error("Failed to get Songs")
		return nil, tx.Error
	}

	return Songs, nil
}

func (s *SongService) GetTextPaginated(id int) (*models.Song, error) {
	logrus.Debug("Get song DB in GetTextPaginated")
	var song models.Song
	tx := s.DB.First(&song, id)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logrus.WithField("id", id).Error("Song not found")
		return nil, tx.Error
	}

	if tx.Error != nil {
		logrus.WithField("id", id).Error("Failed to get song")
		return nil, tx.Error
	}

	return &song, nil
}

func (s *SongService) Create(Song *models.Song) error {
	logrus.Debug("Add song DB in Create")
	tx := s.DB.Create(&Song)
	if tx.Error != nil {
		logrus.Error("Failed to create Song")
		return tx.Error
	}

	return nil
}

func (s *SongService) Update(Song *models.Song) error {
	logrus.Debug("Update song DB in Update")
	tx := s.DB.Model(&Song).Where("id = ?", Song.ID)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logrus.WithField("id", Song.ID).Error("Song not found")
		return tx.Error
	}

	if tx.Error != nil {
		logrus.WithField("id", Song.ID).Error("Failed to update Song")
		return tx.Error
	}

	return nil
}

func (s *SongService) Delete(id int) error {
	logrus.Debug("Delete song DB in Delete")
	var Song models.Song
	tx := s.DB.Delete(&Song, id)

	if tx.Error != nil {
		logrus.WithField("id", id).Error("Failed to delete Song")
		return tx.Error
	}

	return nil
}
