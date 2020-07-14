package geokit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance1(t *testing.T) {
	distance, _ := Distance(&LatLng{Lat: 41.3083, Lng: -72.9279}, &LatLng{Lat: -33.8688, Lng: 151.2093}, nil)
	assert.Equal(t, distance, 16082.81120656383, "Distance between New Haven and Sydney should be 16082.81120656383 km")
}

func TestDistance2(t *testing.T) {
	unit := "miles"
	distance, _ := Distance(&LatLng{Lat: 40.7128, Lng: -74.006}, &LatLng{Lat: 37.7749, Lng: -122.4194}, &unit)
	assert.Equal(t, distance, 2568.4458439997047, "Distance between New York and San Francisco should be 2568.4458439997047 miles")
}

func TestDistanceFail(t *testing.T) {
	_, err := Distance(&LatLng{Lat: 91, Lng: 0}, &LatLng{Lat: 0, Lng: 0}, nil)
	assert.Error(t, err, "Distance between 91, 0 and 0, 0 with should throw an Error")
}

func TestDistanceFail2(t *testing.T) {
	_, err := Distance(&LatLng{Lat: 0, Lng: 0}, &LatLng{Lat: -91, Lng: 0}, nil)
	assert.Error(t, err, "Distance between 0, 0 and -91, 0 with should throw an Error")
}

func TestDistanceFail3(t *testing.T) {
	_, err := Distance(&LatLng{Lat: 180, Lng: 0}, &LatLng{Lat: 0, Lng: 0}, nil)
	assert.Error(t, err, "Distance between 181, 0 and 0, 0 with should throw an Error")
}

func TestDistanceFail4(t *testing.T) {
	_, err := Distance(&LatLng{Lat: 0, Lng: 0}, &LatLng{Lat: -181, Lng: 0}, nil)
	assert.Error(t, err, "Distance between 0, 0 and -181, 0 with should throw an Error")
}
