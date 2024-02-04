package calculator

import "fmt"

var (
	baseWidth = 375.0
)

type WidthModule struct {
}

// Converts a px width to percentage
func (w WidthModule) convertToPercentage(widthPx *float64) (float64, error) {
	resultX := (*widthPx * 100.0) / baseWidth
	return resultX, nil
}

func (w WidthModule) printCSS(widthPx *float64) {
	fmt.Printf("width: %g%%;\n", *widthPx)
}
