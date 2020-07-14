package geokit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validCoordinates = []*LatLng{
	&LatLng{Lat: 0, Lng: 0},
	&LatLng{Lat: -90, Lng: 180},
	&LatLng{Lat: 90, Lng: -180},
	&LatLng{Lat: 23, Lng: 74},
	&LatLng{Lat: 47.235124363, Lng: 127.2379654226},
	&LatLng{Lat: 0},
}

var invalidCoordinates = []*LatLng{
	&LatLng{Lat: -91, Lng: 0},
	&LatLng{Lat: 91, Lng: 0},
	&LatLng{Lat: 0, Lng: 181},
	&LatLng{Lat: 0, Lng: -181},
	&LatLng{Lng: -181},
}

func TestValidateCoordinates(t *testing.T) {
	for _, coord := range validCoordinates {
		_, err := ValidateCoordinates(coord)
		assert.NoError(t, err, "ValidateCoordinates() does not throw errors given valid coordinates")
	}
}

func TestValidateCoordinatesFail(t *testing.T) {
	for _, coord := range invalidCoordinates {
		_, err := ValidateCoordinates(coord)
		assert.Error(t, err, "throws errors given invalid coordinates")
	}
}
