package svg

func PrepareShapes(
	shapes []Shape,
	spacing float64,
) []Shape {
	var result []Shape

	for _, s := range shapes {
		p := s.Polygon

		if !p.IsValid() {
			continue
		}

		p = p.Normalize()
		p = p.Offset(spacing)

		result = append(result, Shape{
			Polygon: p,
		})
	}

	return result
}
