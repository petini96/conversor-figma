package calculator

import "fmt"

type LeftModule struct {
}

func (w LeftModule) convertToPercentage(widthPx *float64) (float64, error) {
	resultX := (*widthPx * 100.0) / baseWidth
	return resultX, nil
}

func (w LeftModule) printCSS(widthPx *float64) {
	fmt.Printf("left: %g%%;\n", *widthPx)
}
