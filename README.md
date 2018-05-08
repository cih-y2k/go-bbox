# go-bbox

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
	point := bbox.Point{
		Latitude:  0,
		Longitude: -180,
	}

  // Finds the min/max points for a rectangular area containing a 32km radius.
	bboxes := bbox.New(point, 32)
  for _, bbox := range bboxes {
    fmt.Printf("%+v\n", bbox)
  }
}
```

## Example output

```
{Min:{Latitude:-0.28746089091824684 Longitude:179.71253910908138} Max:{Latitude:0.28746089091824684 Longitude:180}}
{Min:{Latitude:-0.28746089091824684 Longitude:-180} Max:{Latitude:0.28746089091824684 Longitude:-179.71253910908138}}
```
