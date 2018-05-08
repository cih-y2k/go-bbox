// This package is an implementation of the geospatial algorithms located here:
// https://web.archive.org/web/20180508002202/http://janmatuschek.de/LatitudeLongitudeBoundingCoordinates#UsingIndex
// Bronshtein, Semendyayev, Musiol, Mühlig: Handbook of Mathematics. Springer, Berlin.
//
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

func calcAngularRadius(radius float64) float64 {
	const equatorialRadius = 6378137
	return 1000 * radius / equatorialRadius
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

func (point *Point) calcLatT(angularRadius float64) float64 {
	return math.Asin(math.Sin(point.Latitude) / math.Cos(angularRadius))
}

func (point *Point) calcDeltaLon(angularRadius, latT float64) float64 {
	return math.Acos((math.Cos(angularRadius) - math.Sin(latT)*math.Sin(point.Latitude)) / math.Cos(latT) * math.Cos(point.Latitude))
}

func (point Point) normalizeMeridian() Point {
	point.Longitude = normalizeMeridian(point.Longitude)

	return point
}

func (point Point) toDegrees() Point {
	point.Latitude = toDegrees(point.Latitude)
	point.Longitude = toDegrees(point.Longitude)

	return point
}

// BBox represents minimum and maximum bounding box point coordinates.
type BBox struct {
	Min,
	Max Point
}

func (bbox *BBox) applyAngularRadius(angularRadius float64, point Point) *BBox {
	bbox.Min.Latitude = point.Latitude - angularRadius
	bbox.Max.Latitude = point.Latitude + angularRadius

	return bbox
}

func (bbox *BBox) applyDeltaLon(deltaLon float64, point Point) *BBox {
	bbox.Min.Longitude = point.Longitude - deltaLon
	bbox.Max.Longitude = point.Longitude + deltaLon

	return bbox
}

func (bbox *BBox) handleNorthPoll() *BBox {
	if bbox.Max.Latitude > math.Pi/2 {
		bbox.Min.Longitude = -math.Pi
		bbox.Max.Latitude = math.Pi / 2
		bbox.Max.Longitude = math.Pi
	}
	return bbox
}

func (bbox *BBox) handleSouthPoll() *BBox {
	if bbox.Min.Latitude < -math.Pi/2 {
		bbox.Min.Latitude = -math.Pi / 2
		bbox.Min.Longitude = -math.Pi
		bbox.Max.Longitude = math.Pi
	}
	return bbox
}

func (bbox *BBox) handleMeridian180() []BBox {
	// Handle wraparound if minimum longitude is less than -180 degrees.
	if bbox.Min.Longitude < -math.Pi {
		return []BBox{
			{
				Min: Point{
					Latitude:  bbox.Min.Latitude,
					Longitude: bbox.Min.Longitude + 2*math.Pi,
				},
				Max: Point{
					Latitude:  bbox.Max.Latitude,
					Longitude: math.Pi,
				},
			},
			{
				Min: Point{
					Latitude:  bbox.Min.Latitude,
					Longitude: -math.Pi,
				},
				Max: Point{
					Latitude:  bbox.Max.Latitude,
					Longitude: bbox.Max.Longitude,
				},
			},
		}
	}
	// Handle wraparound if maximum longitude is greater than 180 degrees.
	if bbox.Max.Longitude > math.Pi {
		return []BBox{
			{
				Min: Point{
					Latitude:  bbox.Min.Latitude,
					Longitude: bbox.Min.Longitude,
				},
				Max: Point{
					Latitude:  bbox.Max.Latitude,
					Longitude: -math.Pi,
				},
			},
			{
				Min: Point{
					Latitude:  bbox.Min.Latitude,
					Longitude: -math.Pi,
				},
				Max: Point{
					Latitude:  bbox.Max.Latitude,
					Longitude: bbox.Max.Longitude - 2*math.Pi,
				},
			},
		}
	}
	return []BBox{*bbox}
}

func (bbox BBox) normalizeMeridian() BBox {
	bbox.Min = bbox.Min.normalizeMeridian()
	bbox.Max = bbox.Max.normalizeMeridian()

	return bbox
}

func (bbox BBox) toDegrees() BBox {
	bbox.Min = bbox.Min.toDegrees()
	bbox.Max = bbox.Max.toDegrees()

	return bbox
}

// New calculates and returns one or more bounding boxes using a coordinate point
// and radius in kilometers.
func New(radius float64, pointVal Point) []BBox {
	angularRadius := calcAngularRadius(radius)
	point := pointVal.toRadians()

	bbox := new(BBox)
	bbox = bbox.applyAngularRadius(angularRadius, point)

	latT := point.calcLatT(angularRadius)
	deltaLon := point.calcDeltaLon(angularRadius, latT)

	bbox = bbox.applyDeltaLon(deltaLon, point)

	bbox = bbox.handleNorthPoll()
	bbox = bbox.handleSouthPoll()

	bboxes := bbox.handleMeridian180()
	for i, bbox := range bboxes {
		bboxes[i] = bbox.normalizeMeridian()
		bboxes[i] = bbox.toDegrees()
	}

	return bboxes
}
