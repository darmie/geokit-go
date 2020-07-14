An assortment of geolocation related tools, all packaged in one easy to use kit. A direct port from http://github.com/MichaelSolati/geokit

## Methods

### `Distance(start *LatLng, end *LatLng, unit *string)(float64, error)`

Calculates the distance, in kilometers, between two coordinates.

`start` and `end` must be LatLng `{ Lat: 0, Lng: 0 }`.

```go
start := &geokit.LatLng{Lat: 41.3083, Lng: -72.9279};
end := &geokit.LatLng{Lat: -33.8688, Lng: 151.2093};

distance, _ := geokit.Distance(location1, location2, nil); // distance == 16082.811206563834
```

### `Hash(coordinates *LatLng, precision *int)(*string, error)`

Generates Geohash of coordinates.

`coordinates` must be LatLng `{ Lat: 0, Lng: 0 }`.

```go
coordinates := &geokit.LatLng{Lat: 41.3083,  Lng: -72.9279};

hash, _ := geokit.Hash(coordinates, nil); // hash == 'drk4urzw2c'
```

### `DecodeHash(hash string) (*LatLng, error)`

Decodes a Geohash into its Latitude and Longitude as a LatLng.

```go
hash := 'r3gx2f77b';

coordinates, _ := geokit.DecodeHash(hash); // coordinates === &geokit.LatLng{Lat: -33.86881113052368,  Lng: 151.2093186378479}
```

### `ValidateCoordinates(coordinates *LatLng) (bool, error)`

Validates coordinates and returns a boolean if valid, or throws an error if invalid.

`coordinates` must be LatLng `{ lat: 0, lng: 0 }`.

```go
coordinates := &geokit.LatLng{Lat: 41.3083,  Lng: -72.9279};

isValid, _ = geokit.ValidateCoordinates(coordinates); // true
```

### `ValidateHash(hash string) (bool, error)`

Validates a Geohash and returns a boolean if valid, or throws an error if invalid.

```go

hash := 'r3gx2f77b';

isValid = geokit.ValidateHash(hash); // true
```
