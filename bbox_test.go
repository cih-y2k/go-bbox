package bbox

import "testing"

func TestNew(t *testing.T) {
	bboxes := New(0, Point{})
	if bboxes == nil {
		t.Fatal("bboxes == nil")
	}
}

func testBBoxes(t *testing.T, radius float64, point Point, bboxesAssertionVal []BBox, pointWraps bool) {
	bboxes := New(radius, point)
	if !pointWraps {
		if len(bboxes) != 1 {
			t.Fatalf("!pointWraps && len(bboxes) != 1: %d\n",
				len(bboxes),
			)
		}
	}
	if pointWraps {
		if len(bboxes) != 2 {
			t.Fatalf("pointWraps && len(bboxes) != 2: %d\n",
				len(bboxes),
			)
		}
	}
	for i, bbox := range bboxes {
		if bbox.Min.Latitude != bboxesAssertionVal[i].Min.Latitude {
			t.Fatalf("bboxes[%d].Min.Latitude != %v: %v\n",
				i,
				bboxesAssertionVal[i].Min.Latitude,
				bbox.Min.Latitude,
			)
		}
		if bbox.Min.Longitude != bboxesAssertionVal[i].Min.Longitude {
			t.Fatalf("bboxes[%d].Min.Longitude != %v: %v\n",
				i,
				bboxesAssertionVal[i].Min.Longitude,
				bbox.Min.Longitude,
			)
		}
		if bbox.Max.Latitude != bboxesAssertionVal[i].Max.Latitude {
			t.Fatalf("bboxes[%d].Max.Latitude != %v: %v\n",
				i,
				bboxesAssertionVal[i].Max.Latitude,
				bbox.Max.Latitude,
			)
		}
		if bbox.Max.Longitude != bboxesAssertionVal[i].Max.Longitude {
			t.Fatalf("bboxes[%d].Max.Longitude != %v: %v\n",
				i,
				bboxesAssertionVal[i].Max.Longitude,
				bbox.Max.Longitude,
			)
		}
		if bbox.Min.Latitude > bbox.Max.Latitude {
			t.Fatalf("bbox.Min.Latitude > bbox.Max.Latitude: %v > %v\n",
				bbox.Min.Latitude,
				bbox.Max.Latitude,
			)
		}
		if bbox.Min.Longitude > bbox.Max.Longitude {
			t.Fatalf("bbox.Min.Longitude > bbox.Max.Longitude: %v > %v\n",
				bbox.Min.Longitude,
				bbox.Max.Longitude,
			)
		}
	}
}

func TestNYC(t *testing.T) {
	const radius = 123.123654
	point := Point{
		Latitude:  40.7491902,
		Longitude: -74.0057076,
	}
	bboxesAssertionVal := []BBox{
		{
			Min: Point{
				Latitude:  39.643151597751555,
				Longitude: -75.46574887344373,
			},
			Max: Point{
				Latitude:  41.85522880224843,
				Longitude: -72.54566632655624,
			},
		},
	}
	testBBoxes(t, radius, point, bboxesAssertionVal, false)
}

func TestLondon(t *testing.T) {
	const radius = 322.14
	point := Point{
		Latitude:  51.5073482,
		Longitude: -0.1452675,
	}
	bboxesAssertionVal := []BBox{
		{
			Min: Point{
				Latitude:  48.613515343737376,
				Longitude: -4.797770051087011,
			},
			Max: Point{
				Latitude:  54.40118105626262,
				Longitude: 4.507235051087011,
			},
		},
	}
	testBBoxes(t, radius, point, bboxesAssertionVal, false)
}

func TestMontevideo(t *testing.T) {
	const radius = 56.0
	point := Point{
		Latitude:  -34.8283457,
		Longitude: -56.3119767,
	}
	bboxesAssertionVal := []BBox{
		{
			Min: Point{
				Latitude:  -35.33140225910693,
				Longitude: -56.924816336439505,
			},
			Max: Point{
				Latitude:  -34.32528914089306,
				Longitude: -55.69913706356049,
			},
		},
	}
	testBBoxes(t, radius, point, bboxesAssertionVal, false)
}

func TestToloke(t *testing.T) {
	const radius = 276.494742
	point := Point{
		Latitude:  -14.2436432,
		Longitude: -178.1795257,
	}
	bboxesAssertionVal := []BBox{
		{
			Min: Point{
				Latitude:  -16.727437727172838,
				Longitude: 179.2578499517392,
			},
			Max: Point{
				Latitude:  -11.759848672827163,
				Longitude: 180,
			},
		},
		{
			Min: Point{
				Latitude:  -16.727437727172838,
				Longitude: -180,
			},
			Max: Point{
				Latitude:  -11.759848672827163,
				Longitude: -175.61690135173924,
			},
		},
	}
	testBBoxes(t, radius, point, bboxesAssertionVal, true)
}

func TestSuva(t *testing.T) {
	const radius = 453.2345
	point := Point{
		Latitude:  -18.1236158,
		Longitude: 178.427969,
	}
	bboxesAssertionVal := []BBox{
		{
			Min: Point{
				Latitude:  -22.195090586402696,
				Longitude: 174.1435668119667,
			},
			Max: Point{
				Latitude:  -14.052141013597307,
				Longitude: 180,
			},
		},
		{
			Min: Point{
				Latitude:  -22.195090586402696,
				Longitude: -180,
			},
			Max: Point{
				Latitude:  -14.052141013597307,
				Longitude: -177.28762881196675,
			},
		},
	}
	testBBoxes(t, radius, point, bboxesAssertionVal, true)
}

func TestSouthPole(t *testing.T) {
	const radius = 194.645
	point := Point{
		Latitude:  -88.6349537,
		Longitude: 51.3556355,
	}
	bboxesAssertionVal := []BBox{
		{
			Min: Point{
				Latitude:  -90,
				Longitude: -180,
			},
			Max: Point{
				Latitude:  -86.88642791522555,
				Longitude: 180,
			},
		},
	}
	// Point does not trigger a typical 180 deg wraparound because it crosses
	// the south pole.
	testBBoxes(t, radius, point, bboxesAssertionVal, false)
}

func TestNorthPole(t *testing.T) {
	const radius = 85.245
	point := Point{
		Latitude:  89.6349537,
		Longitude: 51.3556355,
	}
	bboxesAssertionVal := []BBox{
		{
			Min: Point{
				Latitude:  88.86918483605231,
				Longitude: -180,
			},
			Max: Point{
				Latitude:  90,
				Longitude: 180,
			},
		},
	}
	// Point does not trigger a typical 180 deg wraparound because it crosses
	// the north pole.
	testBBoxes(t, radius, point, bboxesAssertionVal, false)
}
