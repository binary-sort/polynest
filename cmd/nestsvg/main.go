package main

import (
	"flag"
	"fmt"
	"os"
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

	// TODO:
	// 1. Parse SVG
	// 2. Convert paths to polygons
	// 3. Run nesting algorithm
	// 4. Write output SVG
}
