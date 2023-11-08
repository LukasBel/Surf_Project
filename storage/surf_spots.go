package storage

type SurfSpots struct {
	ID         uint8  `gorm:"primaryKey;default:0" json:"id"`
	Type       string `json:"type"`
	WaveHeight string `json:"waveHeight"`
	WavePower  int    `json:"wavePower"`
	SkillLevel string `json:"skillLevel"`
}
