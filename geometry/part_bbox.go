package geometry

func BoundingBoxOfPolygons(polys []Polygon) BoundingBox {
	if len(polys) == 0 {
		return BoundingBox{}
	}

	bb := polys[0].BoundingBox()

	for _, p := range polys[1:] {
		pbb := p.BoundingBox()

		if pbb.Min.X < bb.Min.X {
			bb.Min.X = pbb.Min.X
		}
		if pbb.Min.Y < bb.Min.Y {
			bb.Min.Y = pbb.Min.Y
		}
		if pbb.Max.X > bb.Max.X {
			bb.Max.X = pbb.Max.X
		}
		if pbb.Max.Y > bb.Max.Y {
			bb.Max.Y = pbb.Max.Y
		}
	}

	return bb
}
