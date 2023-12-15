package models

import (
	"gorm.io/gorm"
	"time"
)

type SurfSpots struct {
	ID         uint8     `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	WaveHeight string    `json:"waveHeight"`
	WavePower  int       `json:"wavePower"`
	SkillLevel string    `json:"skillLevel"`
	CreatedAt  time.Time `json:"createdAt"`
}

func MigrateSpots(db *gorm.DB) error {
	err := db.AutoMigrate(&SurfSpots{})
	if err != nil {
		return err
	}
	return nil
}
