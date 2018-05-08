// This package is an implementation of the geospatial algorithms located here:
// https://web.archive.org/web/20180508002202/http://janmatuschek.de/LatitudeLongitudeBoundingCoordinates#UsingIndex
// Bronshtein, Semendyayev, Musiol, Mühlig: Handbook of Mathematics. Springer, Berlin.
// This port uses Earth's equatorial radius for compatibility with Google Maps,
// Bing Maps, and Mapbox.
// Ported by Justin Lowery (ISC License)

package bbox

import "math"

func toRadians(deg float64) float64 { return deg * math.Pi / 180 }
func toDegrees(rad float64) float64 { return rad * 180 / math.Pi }

func normalizeMeridian(lon float64) float64 {
	return math.Mod(lon+3*math.Pi, 2*math.Pi) - math.Pi
}

// Point represents a physical point on Earth, referenced by latitude and longitude coordinates.
type Point struct {
	Latitude,
	Longitude float64
}

func (point Point) toRadians() Point {
	point.Latitude = toRadians(point.Latitude)
	point.Longitude = toRadians(point.Longitude)

	return point
}

func (point Point) toDegrees() Point {
	point.Latitude = toDegrees(point.Latitude)
	point.Longitude = toDegrees(point.Longitude)

	return point
}

func (point Point) normalizeMeridian() Point {
	point.Longitude = normalizeMeridian(point.Longitude)

	return point
}

// BBox represents minimum and maximum bounding box point coordinates.
type BBox struct {
	Min,
	Max Point
}

func (bbox BBox) toDegrees() BBox {
	bbox.Min = bbox.Min.toDegrees()
	bbox.Max = bbox.Max.toDegrees()

	return bbox
}

func (bbox BBox) normalizeMeridian() BBox {
	bbox.Min = bbox.Min.normalizeMeridian()
	bbox.Max = bbox.Max.normalizeMeridian()

	return bbox
}

// New calculates and returns one or more bounding boxes using a coordinate point
// and radius in kilometers.
func New(radius float64, point Point) []BBox {
	point = point.toRadians()

	const equatorialRadius = 6378137
	angularRadius := 1000 * radius / equatorialRadius

	bbox := BBox{
		Min: Point{
			Latitude: point.Latitude - angularRadius,
		},
		Max: Point{
			Latitude: point.Latitude + angularRadius,
		},
	}

	latT := math.Asin(math.Sin(point.Latitude) / math.Cos(angularRadius))
	deltaLon := math.Acos((math.Cos(angularRadius) - math.Sin(latT)*math.Sin(point.Latitude)) / math.Cos(latT) * math.Cos(point.Latitude))

	bbox.Min.Longitude = point.Longitude - deltaLon
	bbox.Max.Longitude = point.Longitude + deltaLon

	if bbox.Max.Latitude > math.Pi/2 {
		// Handle the North Poll.
		bbox.Min.Longitude = -math.Pi
		bbox.Max.Latitude = math.Pi / 2
		bbox.Max.Longitude = math.Pi
	}

	if bbox.Min.Latitude < -math.Pi/2 {
		// Handle the South Poll.
		bbox.Min.Longitude = -math.Pi / 2
		bbox.Min.Latitude = -math.Pi
		bbox.Max.Longitude = math.Pi
	}

	bboxes := make([]BBox, 0, 1)
	if bbox.Min.Longitude < -math.Pi {
		// Handle wraparound if minimum longitude is less than -180 degrees.
		bboxes = append(bboxes, BBox{
			Min: Point{
				Latitude:  bbox.Min.Latitude,
				Longitude: bbox.Min.Longitude + 2*math.Pi,
			},
			Max: Point{
				Latitude:  bbox.Max.Latitude,
				Longitude: math.Pi,
			},
		})
		bboxes = append(bboxes, BBox{
			Min: Point{
				Latitude:  bbox.Min.Latitude,
				Longitude: -math.Pi,
			},
			Max: Point{
				Latitude:  bbox.Max.Latitude,
				Longitude: bbox.Max.Longitude,
			},
		})
	} else if bbox.Max.Longitude > math.Pi {
		// Handle wraparound if maximum longitude is greater than 180 degrees.
		bboxes = append(bboxes, BBox{
			Min: Point{
				Latitude:  bbox.Min.Latitude,
				Longitude: bbox.Min.Longitude,
			},
			Max: Point{
				Latitude:  bbox.Max.Latitude,
				Longitude: -math.Pi,
			},
		})
		bboxes = append(bboxes, BBox{
			Min: Point{
				Latitude:  bbox.Min.Latitude,
				Longitude: -math.Pi,
			},
			Max: Point{
				Latitude:  bbox.Max.Latitude,
				Longitude: bbox.Max.Longitude - 2*math.Pi,
			},
		})
	} else {
		bboxes = append(bboxes, bbox)
	}

	for i, bbox := range bboxes {
		bboxes[i] = bbox.normalizeMeridian()
		bboxes[i] = bbox.toDegrees()
	}

	return bboxes
}
