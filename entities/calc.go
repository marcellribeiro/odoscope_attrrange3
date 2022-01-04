package entities

type Calculator struct {
	Method string    `json:"method"`
	Data   []float64 `json:"data"`
	Config Config    `json:"config"`
}

type CalculatorUsecase interface {
	ThreeSigmas(*Calculator) error
}
