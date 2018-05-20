# go-bbox

[![Build Status](https://travis-ci.org/umpc/go-bbox.svg?branch=master)](https://travis-ci.org/umpc/go-bbox)
[![Coverage Status](https://codecov.io/github/umpc/go-bbox/badge.svg?branch=master)](https://codecov.io/github/umpc/go-bbox?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/umpc/go-bbox)](https://goreportcard.com/report/github.com/umpc/go-bbox)
[![GoDoc](https://godoc.org/github.com/umpc/go-bbox?status.svg)](https://godoc.org/github.com/umpc/go-bbox)

This package is an implementation of a geospatial bounding box algorithm located [here](https://web.archive.org/web/20180508002202/http://janmatuschek.de/LatitudeLongitudeBoundingCoordinates#UsingIndex).

```sh
go get -u github.com/umpc/go-bbox
```

* Earth's equatorial radius of 6,378,137 meters is used for compatibility
with online mapping services such as:
  * Google Maps
  * Bing Maps
  * Mapbox
* Bounds that cross the antimeridian are represented using two bounding boxes.
* Bounds that cross the poles are represented using a single bounding box that has:
  * a min longitude of -180 degrees
  * a max longitude of 180 degrees

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

## Benchmark results

```sh
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/umpc/go-bbox
BenchmarkNewEmpty-8             10000000               181 ns/op
BenchmarkNYC-8                   5000000               263 ns/op
BenchmarkLondon-8                5000000               267 ns/op
BenchmarkMontevideo-8            5000000               268 ns/op
BenchmarkToloke-8                5000000               320 ns/op
BenchmarkSuva-8                  5000000               322 ns/op
BenchmarkNorthPole-8            10000000               200 ns/op
BenchmarkSouthPole-8            10000000               202 ns/op
PASS
ok      github.com/umpc/go-bbox 15.109s
```

* CPU: Intel Core i7-4790k
* Memory: DDR3-1600

## References:

* Bronshtein, Semendyayev, Musiol, Mühlig: Handbook of Mathematics. Springer, Berlin. ISBN-13: 978-3817120079
