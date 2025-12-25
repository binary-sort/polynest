package geometry

import "math"

type Polygon struct {
	Points []Point
}

func (p Polygon) Area() float64 {
	area := 0.0
	n := len(p.Points)

	for i := 0; i < n-1; i++ {
		x1, y1 := p.Points[i].X, p.Points[i].Y
		x2, y2 := p.Points[i+1].X, p.Points[i+1].Y

		area += (x1*y2 - y1*x2)
	}

	return math.Abs(area) / 2
}

func (p Polygon) IsValid() bool {
	return len(p.Points) >= 4 && p.Area() > 0
}

func (p Polygon) Normalize() Polygon {
	bb := p.BoundingBox()

	dx := -bb.Min.X
	dy := -bb.Min.Y

	var (
		pts []Point
	)

	for _, pt := range p.Points {
		pts = append(pts, Point{
			X: pt.X + dx,
			Y: pt.Y + dy,
		})
	}

	return Polygon{Points: pts}
}

func (p Polygon) Offset(spacing float64) Polygon {
	if spacing <= 0 {
		return p
	}

	var pts []Point
	for _, pt := range p.Points {
		pts = append(pts, Point{
			X: pt.X + spacing,
			Y: pt.Y + spacing,
		})
	}

	return Polygon{Points: pts}
}
