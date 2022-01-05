package entities

type Calculator struct {
	Method string    `json:"method"`
	Data   []float64 `json:"data"`
}

type CalculatorUsecase interface {
	ThreeSigmas(*Calculator) error
}
