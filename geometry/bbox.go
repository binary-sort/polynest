package geometry

type BoundingBox struct {
	Min Point
	Max Point
}

func (p Polygon) BoundingBox() BoundingBox {
	if len(p.Points) == 0 {
		return BoundingBox{}
	}

	minX, minY := p.Points[0].X, p.Points[0].Y
	maxX, maxY := minX, minY

	for _, pt := range p.Points {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}

	return BoundingBox{
		Min: Point{minX, minY},
		Max: Point{maxX, maxY},
	}
}
