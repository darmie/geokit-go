package geokit

import "math"

func ToRad(degrees float64) float64 {
	return (degrees * math.Pi) / 180
}
