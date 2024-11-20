package db

import (
	"errors"
	"fmt"
	"online_song_library/internal/config"
	"online_song_library/internal/models"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	logrus.Debug("Init DB")
	cfg, err := config.Init()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", cfg.POSTGRES_USER,
		cfg.POSTGRES_PASSWORD, cfg.POSTGRES_DB, cfg.POSTGRES_HOST, cfg.POSTGRES_PORT)

	var DB *gorm.DB
	for DB, err = connect(connStr); err != nil; {
		DB, err = connect(connStr)
		if err != nil {
			logrus.Warnf("%v... Retrying to connect...", err)
			time.Sleep(5 * time.Second)
		}
	}

	err = autoMigrate(DB)
	if err != nil {
		return nil, err
	}

	return DB, nil
}

func connect(dsn string) (*gorm.DB, error) {
	logrus.Debug("Connecting to DB in Connect")
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Failed to connect to database: " + err.Error())
	}
	return DB, nil
}

func autoMigrate(db *gorm.DB) error {
	logrus.Debug("Migrating DB in autoMigrate")
	err := db.AutoMigrate(&models.Song{})
	if err != nil {
		return errors.New("Failed to migrate database: " + err.Error())
	}
	return nil
}
