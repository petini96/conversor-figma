package calculator

type BaseMeasure struct {
	width  int
	height int
}

type BaseScreen struct {
	BaseMeasure
}

func NewBaseScreen(width int, height int) *BaseScreen {
	return &BaseScreen{
		BaseMeasure{
			width:  width,
			height: height,
		},
	}
}
