package calculator

import (
	"fmt"
)

var (
	baseHeight = 667.0
)

type HeightModule struct {
}

func (h HeightModule) convertToPercentage(heightPx *float64) (float64, error) {
	resultY := (*heightPx * 100.0) / baseHeight
	return resultY, nil
}

func (h HeightModule) printCSS(heightPercentage *float64) {
	fmt.Printf("height: %g%%;\n", *heightPercentage)
}
