package svg

import (
	"errors"

	"github.com/binary-sort/polynest.git/geometry"
)

func ParsePath(d string) ([]geometry.Polygon, error) {
	tokens, err := tokenizePath(d)

	if err != nil {
		return nil, err
	}

	var polys []geometry.Polygon
	var curr []geometry.Point

	var cx, cy float64
	var sx, sy float64

	i := 0

	for i < len(tokens) {
		t := tokens[i]

		if t.typ != tokenCommand {
			return nil, errors.New("expected command")
		}

		cmd := t.val
		i++

		switch cmd {

		case "M", "m":
			if i+1 > len(tokens) {
				return nil, errors.New("invalid M command")
			}

			x := parseFloat(tokens[i].val)
			y := parseFloat(tokens[i+1].val)
			i += 2

			if cmd == "m" {
				x += cx
				y += cy
			}

			if len(curr) > 0 {
				curr = nil
			}

			cx, cy = x, y
			sx, sy = x, y
			curr = append(curr, geometry.Point{X: x, Y: y})

		case "L", "l":
			x := parseFloat(tokens[i].val)
			y := parseFloat(tokens[i+1].val)
			i += 2

			if cmd == "l" {
				x += cx
				y += cy
			}

			cx, cy = x, y
			curr = append(curr, geometry.Point{X: x, Y: y})

		case "H", "h":
			x := parseFloat(tokens[i].val)
			i++

			if cmd == "h" {
				x += cx
			}

			cx = x
			curr = append(curr, geometry.Point{X: cx, Y: cy})

		case "V", "v":
			y := parseFloat(tokens[i].val)
			i++

			if cmd == "v" {
				y += cy
			}

			cx = y
			curr = append(curr, geometry.Point{X: cx, Y: cy})

		case "Z", "z":
			curr = append(curr, geometry.Point{X: sx, Y: sy})
			if len(curr) >= 4 {
				polys = append(polys, geometry.Polygon{Points: curr})
			}
			curr = nil

		default:
			return nil, errors.New("unsupported svg path command: " + cmd)
		}
	}

	return polys, nil
}
