package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/binary-sort/polynest.git/geometry"
	"github.com/binary-sort/polynest.git/nesting"
	"github.com/binary-sort/polynest.git/svg"
)

func main() {
	input := flag.String("input", "", "Input SVG file")
	output := flag.String("output", "nested.svg", "Output SVG file")
	width := flag.Float64("width", 1000, "Sheet width")
	height := flag.Float64("height", 1000, "Sheet height")
	spacing := flag.Float64("spacing", 5, "Spacing between parts")

	flag.Parse()

	if *input == "" {
		fmt.Println("Error: input SVG is required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("SVG Nesting Tool")
	fmt.Println("Input:", *input)
	fmt.Println("Output:", *output)
	fmt.Println("Sheet:", *width, "x", *height)
	fmt.Println("Spacing:", *spacing)

	shapes, err := svg.ParseSVG(*input)
	if err != nil {
		panic(err)
	}

	shapes = svg.PrepareShapes(shapes, *spacing)

	packer := nesting.NewPacker(*width, *height, 5)

	var polys []geometry.Polygon
	for _, s := range shapes {
		polys = append(polys, s.Polygon)
	}

	placed := packer.Pack(polys)

	fmt.Println("Prepared shapes:", len(shapes))
	fmt.Println("Placed shapes:", len(placed))

	err = svg.WriteSVG(*output, *width, *height, placed)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output written to:", *output)

}
