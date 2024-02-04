package calculator

type IMeasure interface {
	convertToPercentage(value *float64) (float64, error)
	printCSS(value *float64)
}

func Convert(value *float64, iMeasure IMeasure) (float64, error) {
	return iMeasure.convertToPercentage(value)
}

func PrintCSS(value *float64, iMeasure IMeasure) {
	iMeasure.printCSS(value)
}
