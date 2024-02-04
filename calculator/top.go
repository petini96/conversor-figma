package calculator

import (
	"fmt"
)

type TopModule struct {
}

func (h TopModule) convertToPercentage(heightPx *float64) (float64, error) {
	resultY := (*heightPx * 100.0) / baseHeight
	return resultY, nil
}

func (h TopModule) printCSS(heightPercentage *float64) {
	fmt.Printf("top: %g%%;\n", *heightPercentage)
}
