package svg

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/binary-sort/polynest.git/geometry"
	"github.com/binary-sort/polynest.git/model"
)

func ParseSVGAsPart(path string, quantity int) (model.Part, error) {
	shapes, err := ParseSVG(path)

	if err != nil {
		return model.Part{}, err
	}

	var polys []geometry.Polygon

	for _, s := range shapes {
		if s.Polygon.IsValid() {
			polys = append(polys, s.Polygon)
		}
	}

	part := model.Part{
		Name:     filepath.Base(path),
		Polygons: polys,
		Quantity: quantity,
	}

	part.BBox = geometry.BoundingBoxOfPolygons(part.Polygons)
	return part, nil
}

func ParseSVG(path string) ([]Shape, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		return nil, err
	}

	root := doc.Root()
	var shapes []Shape

	for _, el := range root.FindElements(".//polygon") {
		pointsAttr := el.SelectAttrValue("points", "")
		if pointsAttr == "" {
			continue
		}

		polygon, err := parsePolygonPoints(pointsAttr)
		if err != nil {
			return nil, err
		}

		shapes = append(shapes, Shape{
			Polygon: polygon,
		})
	}

	for _, el := range root.FindElements(".//rect") {
		x, _ := strconv.ParseFloat(el.SelectAttrValue("x", "0"), 64)
		y, _ := strconv.ParseFloat(el.SelectAttrValue("y", "0"), 64)
		w, _ := strconv.ParseFloat(el.SelectAttrValue("width", "0"), 64)
		h, _ := strconv.ParseFloat(el.SelectAttrValue("height", "0"), 64)

		poly := geometry.Polygon{
			Points: []geometry.Point{
				{X: x, Y: y},
				{X: x + w, Y: y},
				{X: x + w, Y: y + h},
				{X: x, Y: y + h},
			},
		}

		shapes = append(shapes, Shape{Polygon: poly})
	}

	for _, el := range root.FindElements(".//path") {
		d := el.SelectAttrValue("d", "")
		if d == "" {
			continue
		}

		polys, err := ParsePath(d)
		if err != nil {
			continue // skip unsupported paths
		}

		for _, p := range polys {
			shapes = append(shapes, Shape{Polygon: p})
		}
	}

	return shapes, nil
}

func parsePolygonPoints(points string) (geometry.Polygon, error) {
	fields := strings.FieldsFunc(points, func(r rune) bool {
		return r == ',' || r == ' '
	})

	if len(fields)%2 != 0 {
		return geometry.Polygon{}, fmt.Errorf("invalid polygon points")
	}

	var pts []geometry.Point
	for i := 0; i < len(fields); i += 2 {
		x, _ := strconv.ParseFloat(fields[i], 64)
		y, _ := strconv.ParseFloat(fields[i+1], 64)
		pts = append(pts, geometry.Point{X: x, Y: y})
	}

	return geometry.Polygon{Points: pts}, nil
}
