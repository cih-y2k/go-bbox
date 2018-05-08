# go-bbox

[![Build Status](https://travis-ci.org/umpc/go-bbox.svg?branch=master)](https://travis-ci.org/umpc/go-bbox)
[![Coverage Status](https://codecov.io/github/umpc/go-bbox/badge.svg?branch=master)](https://codecov.io/github/umpc/go-bbox?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/umpc/go-bbox)](https://goreportcard.com/report/github.com/umpc/go-bbox)
[![GoDoc](https://godoc.org/github.com/umpc/go-bbox?status.svg)](https://godoc.org/github.com/umpc/go-bbox)

This package is an implementation of the geospatial algorithms located [here](https://web.archive.org/web/20180508002202/http://janmatuschek.de/LatitudeLongitudeBoundingCoordinates#UsingIndex).

Bronshtein, Semendyayev, Musiol, Mühlig: Handbook of Mathematics. Springer, Berlin.

This port uses Earth's equatorial radius, 6,378,137 meters, for compatibility
with online mapping services such as Google Maps, Bing Maps, and Mapbox.

The 180th meridian is handled by the caller searching within two bounding boxes.

```sh
go get -u github.com/umpc/go-bbox
```

## Example usage

```go
package main

import (
  "fmt"

  "github.com/umpc/go-bbox"
)

func main() {
  // Finds the min/max points for a rectangular area, which will be 276.49km
  // southwest and northeast of the center point.
  bboxes := bbox.New(276.494742, bbox.Point{
    Latitude:  -14.2436432,
    Longitude: -178.1795257,
  })
  for _, bbox := range bboxes {
    fmt.Printf("%+v\n", bbox)
  }
}
```

## Example output

```
{Min:{Latitude:-16.727437727172838 Longitude:179.2578499517392} Max:{Latitude:-11.759848672827163 Longitude:180}}
{Min:{Latitude:-16.727437727172838 Longitude:-180} Max:{Latitude:-11.759848672827163 Longitude:-175.61690135173924}}
```
