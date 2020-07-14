package geokit

import (
	"math"
	"strings"
)

// Distance : Calculates the distance, in kilometers, between two coordinates.
func Distance(start *LatLng, end *LatLng, unit *string) (float64, error) {
	if unit == nil || *unit != "miles" {
		_u := "km"
		unit = &_u
	}
	_, err := ValidateCoordinates(start)
	if err != nil {
		return -1, err
	}

	_, err = ValidateCoordinates(end)
	if err != nil {
		return -1, err
	}

	radius := float64(6371)
	if strings.ToLower(*unit) == "miles" {
		radius = float64(3963)
	}

	dLat := ToRad(end.Lat - start.Lat)
	dLon := ToRad(end.Lng - start.Lng)

	lat1 := ToRad(start.Lat)
	lat2 := ToRad(end.Lat)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return radius * c, nil
}
