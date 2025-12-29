package nesting

import "github.com/binary-sort/polynest.git/geometry"

type PlacedShape struct {
	Polygon geometry.Polygon
	X       float64
	Y       float64
	BBox    geometry.BoundingBox
}
