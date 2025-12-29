package model

import "github.com/binary-sort/polynest.git/geometry"

type Part struct {
	Name     string
	Polygons []geometry.Polygon
	Quantity int

	BBox geometry.BoundingBox
}
