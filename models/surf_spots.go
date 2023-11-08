package models

import "gorm.io/gorm"

type SurfSpots struct {
	ID         uint8  `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	WaveHeight string `json:"waveHeight"`
	WavePower  int    `json:"wavePower"`
	SkillLevel string `json:"skillLevel"`
}

func MigrateSpots(db *gorm.DB) error {
	err := db.AutoMigrate(&SurfSpots{})
	if err != nil {
		return err
	}
	return nil
}
