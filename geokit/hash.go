package geokit

import (
	"errors"
	"fmt"
	"strings"
)

const BASE32 = "0123456789bcdefghjkmnpqrstuvwxyz"

func Hash(coordinates *LatLng, precision *int) (*string, error) {
	if precision == nil {
		_p := 10
		precision = &_p
	}

	_, err := ValidateCoordinates(coordinates)
	if err != nil {
		return nil, err
	}
	if int(*precision) <= 0 {
		return nil, errors.New("Precision must be greater than 0")
	} else if int(*precision) > 22 {
		return nil, errors.New("Precision cannot be greater than 22")
	}

	latRng := []float64{-90, 90}

	lngRng := []float64{-180, 180}

	hash := ""
	hashVal := 0
	bits := 0

	even := true

	for len(hash) < int(*precision) {
		val := coordinates.Lat
		if even {
			val = coordinates.Lng
		}

		_range := latRng
		if even {
			_range = lngRng
		}

		mid := (_range[0] + _range[1]) / 2

		if val > mid {
			hashVal = (hashVal << 1) + 1
			_range[0] = mid
		} else {
			hashVal = (hashVal << 1) + 0
			_range[1] = mid
		}

		even = !even

		if bits < 4 {
			bits++
		} else {
			bits = 0
			hash += string(BASE32[hashVal])
			hashVal = 0
		}

	}
	return &hash, nil
}

// ValidateHash : Validates a Geohash and returns a boolean if valid, or throws an error if invalid.
func ValidateHash(geohash string) (bool, error) {
	var err []string
	if len(geohash) == 0 {
		err = append(err, "geohash cannot be the empty string")
	}

	for _, letter := range geohash {
		if strings.Index(BASE32, string(letter)) == -1 {
			err = append(err, fmt.Sprintf("geohash cannot contain '%s'", string(letter)))
		}
	}

	if len(err) > 0 {
		return false, fmt.Errorf("Invalid geohash '%s':\n %s", geohash, strings.Join(err, "\n"))
	}

	return true, nil
}

// DecodeHash : Decodes a Geohash into its Latitude and Longitude as a `LatLng`
func DecodeHash(hash string) (*LatLng, error) {
	_, err := ValidateHash(hash)
	if err != nil {
		return nil, err
	}

	even := true

	latRng := []float64{-90, 90}

	lngRng := []float64{-180, 180}

	hashChars := strings.Split(hash, "")

	for len(hashChars) > 0 {
		fchar := hashChars[0]
		hashChars = hashChars[1:]
		chunk := strings.Index("0123456789bcdefghjkmnpqrstuvwxyz", strings.ToLower(fchar))
		_maskArr := []int{16, 8, 4, 2, 1}
		for i := 0; i < 5; i++ {
			mask := _maskArr[i]
			_range := latRng
			if even {
				_range = lngRng
			}

			middle := (_range[0] + _range[1]) / 2
			ind := 0
			if (chunk & mask) == 0 {
				ind = 1
			}
			_range[ind] = middle
			even = !even
		}
	}

	return &LatLng{
		Lat: (latRng[0] + latRng[1]) / 2,
		Lng: (lngRng[0] + lngRng[1]) / 2,
	}, nil

}
