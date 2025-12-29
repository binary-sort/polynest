package model

import "github.com/binary-sort/polynest.git/geometry"

func NormalizePart(p Part, spacing float64) Part {
	bb := p.BBox
	dx := -bb.Min.X + spacing
	dy := -bb.Min.Y + spacing

	var polys []geometry.Polygon

	for _, poly := range p.Polygons {
		polys = append(polys, poly.Translate(dx, dy))
	}

	p.Polygons = polys
	p.BBox = geometry.BoundingBoxOfPolygons(polys)
	return p
}
