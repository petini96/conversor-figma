package main

import (
	"flag"
	"fmt"

	"github.com/petini96/conversor-figma/calculator"
)

func main() {
	widthPx := flag.Float64("w", 0.0, "Width in pixels")
	heightPx := flag.Float64("h", 0.0, "Height in pixels")
	leftPx := flag.Float64("l", 0.0, "Left in pixels")
	topPx := flag.Float64("t", 0.0, "Top in pixels")

	flag.Parse()

	if *widthPx != 0.0 {
		var widthModule calculator.WidthModule
		widthPercentage, err := calculator.Convert(widthPx, widthModule)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		calculator.PrintCSS(&widthPercentage, widthModule)
	}

	if *heightPx != 0.0 {
		var heightModule calculator.HeightModule
		heightPxPercentage, err := calculator.Convert(heightPx, heightModule)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		calculator.PrintCSS(&heightPxPercentage, heightModule)
	}

	if *leftPx != 0.0 {
		var leftModule calculator.LeftModule
		leftPxPercentage, err := calculator.Convert(leftPx, leftModule)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		calculator.PrintCSS(&leftPxPercentage, leftModule)
	}

	if *topPx != 0.0 {
		var topModule calculator.TopModule
		topPxPercentage, err := calculator.Convert(topPx, topModule)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		calculator.PrintCSS(&topPxPercentage, topModule)
	}

}
