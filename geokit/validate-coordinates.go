package geokit

import (
	"fmt"
	"strings"
)

// ValidateCoordinates : Validates coordinates and returns a boolean if valid, or throws an error if invalid.
func ValidateCoordinates(coordinates *LatLng) (bool, error) {
	var err []string

	latitude := coordinates.Lat
	longitude := coordinates.Lng

	if int(latitude) < -90 || int(latitude) > 90 {
		err = append(err, "Latitude must be within the range [-90, 90].")
	} else if int(longitude) < -180 || int(longitude) > 180 {
		err = append(err, "Longitude must be within the range [-180, 180]")
	}

	if len(err) > 0 {
		return false, fmt.Errorf("Invalid coordinates %s", strings.Join(err, " "))
	}

	return true, nil
}
