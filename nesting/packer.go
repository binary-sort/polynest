package nesting

import (
	"sort"

	"github.com/binary-sort/polynest.git/geometry"
)

type Packer struct {
	SheetWidth  float64
	SheetHeight float64
	Step        float64
}

func NewPacker(w, h, step float64) *Packer {
	return &Packer{
		SheetWidth:  w,
		SheetHeight: h,
		Step:        step,
	}
}

func (p *Packer) Pack(polys []geometry.Polygon) []PlacedShape {
	// sort largest area first
	sort.Slice(polys, func(i, j int) bool {
		return polys[i].Area() > polys[j].Area()
	})

	var placed []PlacedShape

	for _, poly := range polys {
		bb := poly.BoundingBox()
		pw := bb.Max.X - bb.Min.X
		ph := bb.Max.Y - bb.Min.Y

	placedLoop:
		for y := 0.0; y+ph <= p.SheetHeight; y += p.Step {
			for x := 0.0; x+pw <= p.SheetWidth; x += p.Step {

				tb := geometry.BoundingBox{
					Min: geometry.Point{X: x, Y: y},
					Max: geometry.Point{X: x + pw, Y: y + ph},
				}

				if collides(tb, placed) {
					continue
				}

				placed = append(placed, PlacedShape{
					Polygon: poly.Translate(x, y),
					X:       x,
					Y:       y,
					BBox:    tb,
				})

				break placedLoop
			}
		}
	}

	return placed
}

func collides(bb geometry.BoundingBox, placed []PlacedShape) bool {
	for _, p := range placed {
		if bb.Intersects(p.BBox) {
			return true
		}
	}
	return false
}
