package main

import (
	"fmt"
	"main/models"
	"math/rand"
)

func WavePowerResponse(spot *models.SurfSpots) (int, error) {
	swellPeriod := rand.Intn(20)

	fmt.Printf("\nSwell Size: %v, Swell Period: %v, Wave Power: %v joules",
		spot.WaveHeight, swellPeriod, spot.WavePower*swellPeriod)
	return spot.WavePower * swellPeriod, nil
}

// this function does not work
func RankSpots(spots *[]models.SurfSpots) error {
	spots = &[]models.SurfSpots{}
	//sort.Slice(&spots, func(i, j int) bool { return spots[i].WavePower < spots[j].WavePower })
	fmt.Println(spots)
	return nil
}
