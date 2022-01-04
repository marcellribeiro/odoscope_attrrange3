package usecase

import (
	"math"

	"github.com/marcellribeiro/odoscope_attrrange3/entities"
)

type calculatorUsecase struct {
	Method string          `json:"method"`
	Config entities.Config `json:"config"`
	Data   []float64       `json:"data"`
}

func NewCalculatorUsecase(method string, config entities.Config, data []float64) *entities.CalculatorUsecase {
	return &calculatorUsecase{
		Method: method,
		Config: config,
		Data:   data,
	}
}

func (a *calculatorUsecase) ThreeSigmasMethod(data []float64) (entities.OutliersResult, error) {
	sd := a.standardDeviation(data)
	result := entities.OutliersResult{
		SiteId: a.Config.Datasets.SiteId,
	}

	// SiteId                  string `json:"siteId"`
	// TimeAgo                 string `json:"TimeAgo"`
	// TimeStep                string `json:"TimeStep"`
	// OutliersDetectionMethod string `json:"OutliersDetectionMethod"`
	// checkTimeStart          string `json:"checkTimeStart"`
	// checkTimeEnd            string `json:"checkTimeEnd"`
	// DateStart               string `json:"DateStart"`
	// DateEnd                 string `json:"DateEnd"`
	// Result                  Result `json:"Result"`

	// Warnings []Warning `json:"Warnings"`
	// Alarms   []Alarm   `json:"Alarms"`

	// OutlierPeriodStart string `json:"OutlierPeriodStart"`
	// OutlierPeriodEnd   string `json:"OutlierPeriodEnd"`
	// Metric             string `json:"Metric"`
	// Attribute          string `json:"Attribute"`

	return nil, nil
}

func (a *calculatorUsecase) standardDeviation(data []float64) float64 {
	var sum, mean, sd float64
	dataSize := float64(len(data))

	for _, num := range data {
		sum += num
	}

	mean = sum / dataSize

	for _, num := range data {
		sd += math.Pow(num-mean, 2)
	}

	return math.Sqrt(sd / dataSize)
}
