package providers

import (
	"math/rand"
	"time"
)

// GetCompanyToGoMars  returns the name and multiplier of company
func GetCompanyToGoMars() (string, float32, int, int, int) {

	//add one second for seed to work
	time.Sleep(time.Second)
	rand.Seed(time.Now().UnixNano())

	companyIndex := rand.Intn(3) + 1
	speed := rand.Intn(15) + 16
	tripPrice := getPrice(speed)
	oneWay := rand.Intn(2)

	switch {
	case companyIndex == 1:
		return "Space Adventure", 1.08, speed, tripPrice, oneWay
	case companyIndex == 2:
		return "SpaceX", 1.02, speed, tripPrice, oneWay
	case companyIndex == 3:
		return "Virgin Galactic", 1.05, speed, tripPrice, oneWay
	default:
		return "", 0, 0, 0, 0
	}
}

func getPrice(speed int) int {

	if speed >= 16 && speed < 21 {
		return 36
	} else if speed >= 21 && speed < 26 {
		return 41
	} else if speed >= 26 && speed < 30 {
		return 48
	} else {
		return 50
	}
}
