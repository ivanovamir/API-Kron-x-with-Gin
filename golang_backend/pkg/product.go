package pkg

import (
	"math"
)

const (
	domain = "https://api.kron-x.ru/data/photos/"
)

func Round(x float64) float64 {
	var rounder float64
	pow := math.Pow(10, float64(2))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
