package services

import "math/rand"

type Location struct {
	Pickup  string
	Dropoff string
	Mileage string
}

func GetMileage(pickup, dropoff string) float32 {
	min := 2.5
	max := 50.00
	randomFloat := min + (max-min) * rand.Float64()
	return float32(randomFloat)
}
