package svg

import (
	"fmt"
	"os"
	"strings"

	"github.com/binary-sort/polynest.git/geometry"
	"github.com/binary-sort/polynest.git/nesting"
)

func WriteSVG(
	outputPath string,
	width, height float64,
	placed []nesting.PlacedShape,
) error {

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// SVG header
	fmt.Fprintf(f,
		`<svg xmlns="http://www.w3.org/2000/svg" width="%f" height="%f" viewBox="0 0 %f %f">`,
		width, height, width, height,
	)
	fmt.Fprintln(f)

	// Optional background
	fmt.Fprintf(f,
		`<rect width="%f" height="%f" fill="white" stroke="black"/>`,
		width, height,
	)
	fmt.Fprintln(f)

	// Draw each polygon
	for i, p := range placed {
		fmt.Fprintf(
			f,
			`<polygon points="%s" fill="none" stroke="black" stroke-width="1"/>`,
			formatPoints(p.Polygon),
		)
		fmt.Fprintln(f)

		// Optional label (debug)
		fmt.Fprintf(
			f,
			`<text x="%f" y="%f" font-size="10" fill="red">%d</text>`,
			p.BBox.Min.X+2, p.BBox.Min.Y+12, i,
		)
		fmt.Fprintln(f)
	}

	fmt.Fprintln(f, `</svg>`)
	return nil
}

func formatPoints(p geometry.Polygon) string {
	var sb strings.Builder

	for _, pt := range p.Points {
		sb.WriteString(fmt.Sprintf("%f,%f ", pt.X, pt.Y))
	}

	return strings.TrimSpace(sb.String())
}
