package geokit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validGeohashes = []string{"4", "d62dtu", "000000000000"}
var invalidGeohashes = []string{"", "aaa"}

func TestDecodeHash1(t *testing.T) {
	geo, _ := DecodeHash("drk4urzw2c")
	assert.Equal(t, geo, &LatLng{
		Lat: 41.30830138921738,
		Lng: -72.9278951883316,
	}, "Geohash 'drk4urzw2c' should yield coordinates 41.30830138921738, -72.9278951883316")
}

func TestDecodeHash2(t *testing.T) {
	geo, _ := DecodeHash("r3gx2f77b")
	assert.Equal(t, geo, &LatLng{
		Lat: -33.86881113052368,
		Lng: 151.2093186378479,
	}, "Geohash 'r3gx2f77b' should yield coordinates -33.86881113052368, 151.2093186378479")
}

func TestDecodeHashFail(t *testing.T) {
	_, err := DecodeHash("aaa")
	assert.Error(t, err, "Throws error when given invalid Geohash")
}

func TestHash(t *testing.T) {
	hash, err := Hash(&LatLng{Lat: 41.3083, Lng: -72.9279}, nil)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, *hash, "drk4urzw2c", "Geohash coordinates 41.3083, -72.9279 should yield 'drk4urzw2c'")
}

func TestHash0(t *testing.T) {
	hash, err := Hash(&LatLng{Lat: 0, Lng: 0}, nil)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, *hash, "7zzzzzzzzz", "Geohash coordinates 0, 0 should yield '7zzzzzzzzz'")
}

func TestHash1(t *testing.T) {
	precision := 9
	hash, err := Hash(&LatLng{Lat: -33.8688, Lng: 151.2093}, &precision)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, *hash, "r3gx2f77b", "Geohash coordinates -33.8688, 151.2093 with a precision of 9 should yield 'r3gx2f77b'")
}

func TestHashFail(t *testing.T) {
	precision := 25
	_, err := Hash(&LatLng{Lat: -33.8688, Lng: 151.2093}, &precision)

	assert.Error(t, err, "Should fail at precision greater than 22")
}

func TestHashFail2(t *testing.T) {
	precision := -1
	_, err := Hash(&LatLng{Lat: -33.8688, Lng: 151.2093}, &precision)

	assert.Error(t, err, "Should fail at precision less than 0")
}

func TestHash2(t *testing.T) {
	_, err := Hash(&LatLng{Lat: 91, Lng: 0}, nil)
	assert.Error(t, err, "Geohash coordinates 91, 0 with should throw an Error")
}

func TestValidateGeoHash(t *testing.T) {
	for _, hash := range validGeohashes {
		_, err := ValidateHash(hash)
		assert.NoError(t, err, "Validatehash() does not throw errors given valid geohashes")
	}
}

func TestValidateGeoHashFail(t *testing.T) {
	for _, hash := range invalidGeohashes {
		_, err := ValidateHash(hash)
		assert.Error(t, err, "Validatehash() throws errors given invalid geohashes")
	}
}
