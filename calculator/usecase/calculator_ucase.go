package usecase

import (
	"context"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/marcellribeiro/odoscope_attrrange3/entities"
)

type calculatorUsecase struct {
	Datasets entities.Datasets `json:"datasets"`
}

func NewCalculatorUsecase(d entities.Datasets) entities.CalculatorUseCase {
	return &calculatorUsecase{
		Datasets: d,
	}
}

func (a *calculatorUsecase) CalculatorService(ctx context.Context) (entities.OutliersResult, error) {
	if a.Datasets.OutliersDetectionMethod == "3-sigmas" {
		return a.threeSigmasMethod(ctx)
	}
	return entities.OutliersResult{}, fmt.Errorf("unknown method")
}

func (a *calculatorUsecase) threeSigmasMethod(ctx context.Context) (entities.OutliersResult, error) {
	var (
		wg     sync.WaitGroup
		eventc = make(chan *entities.Result)
		errc   = make(chan error, 1)

		result = &entities.OutliersResult{
			SiteId:                  a.Datasets.SiteId,
			TimeAgo:                 a.Datasets.TimeAgo,
			TimeStep:                a.Datasets.TimeStep,
			OutliersDetectionMethod: a.Datasets.OutliersDetectionMethod,
			CheckTimeStart:          time.Now().Format("2006-01-02 15:04:05"),
			DateStart:               time.Now().Format("2006-01-02 15:04:05"),
		}
	)

	go func() {
		wg.Wait()
		close(eventc)
		close(errc)
	}()

	for i := range a.Datasets.Attributes {
		for j := range a.Datasets.Attributes[i].Metrics {
			wg.Add(1)

			go func(i, j int) {
				defer wg.Done()

				events, err := a.calcMetrics(MetricParams{
					AttributeName:            a.Datasets.Attributes[i].Name,
					OutliersMultiplier:       a.Datasets.OutliersDetection.OutliersMultiplier,
					StrongOutliersMultiplier: a.Datasets.OutliersDetection.StrongOutliersMultiplier,
					Metrics:                  a.Datasets.Attributes[i].Metrics[j],
				})
				if err != nil {
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				}

				select {
				case <-ctx.Done():
				case eventc <- events:
				}
			}(i, j)

		}
	}

	for events := range eventc {
		result.Result.Alarms = append(result.Result.Alarms, events.Alarms...)
		result.Result.Warnings = append(result.Result.Warnings, events.Warnings...)
	}

	for err := range errc {
		if err != nil {
			return entities.OutliersResult{}, err
		}
	}

	sort.Stable(entities.Alarms(result.Result.Alarms))
	sort.Stable(entities.Warnings(result.Result.Warnings))
	result.CheckTimeEnd = time.Now().Format("2006-01-02 15:04:05")
	result.DateEnd = time.Now().Format("2006-01-02 15:04:05")

	return *result, nil
}

type MetricParams struct {
	AttributeName            string
	OutliersMultiplier       int
	StrongOutliersMultiplier int
	Metrics                  entities.Metrics
}

func (a *calculatorUsecase) calcMetrics(m MetricParams) (*entities.Result, error) {
	sigma, err := a.standardDeviation(m.Metrics)
	if err != nil {
		return nil, err
	}

	return a.findAlerts(m, sigma), nil
}

func (a *calculatorUsecase) findAlerts(m MetricParams, sigma float64) *entities.Result {
	sort.Stable(entities.MetricsList(m.Metrics.Values))

	var (
		warningThreshold = float64(m.OutliersMultiplier) * sigma
		alertThreshold   = float64(m.StrongOutliersMultiplier) * sigma
		events           = &entities.Result{
			Warnings: []entities.Warning{},
			Alarms:   []entities.Alarm{},
		}

		currentAlarm   *entities.Alarm
		currentWarning *entities.Warning
	)

	for i, value := range m.Metrics.Values {
		if value.Value > alertThreshold {
			if currentAlarm == nil {
				currentAlarm = entities.AddAlarm(m.AttributeName, m.Metrics.Name, value)
			}
			if i == len(m.Metrics.Values)-1 {
				currentAlarm.OutlierPeriodEnd = value.Date
				events.Alarms = append(events.Alarms, *currentAlarm)
			}
		} else if currentAlarm != nil {
			currentAlarm.OutlierPeriodEnd = value.Date
			events.Alarms = append(events.Alarms, *currentAlarm)
			currentAlarm = nil
		}

		if value.Value > warningThreshold {
			if currentWarning == nil {
				currentWarning = entities.AddWarning(m.AttributeName, m.Metrics.Name, value)
			}
			if i == len(m.Metrics.Values)-1 {
				currentWarning.OutlierPeriodEnd = value.Date
				events.Warnings = append(events.Warnings, *currentWarning)
			}
		} else if currentWarning != nil {
			currentWarning.OutlierPeriodEnd = value.Date
			events.Warnings = append(events.Warnings, *currentWarning)
			currentWarning = nil
		}
	}

	return events
}

func (a *calculatorUsecase) standardDeviation(metrics entities.Metrics) (float64, error) {
	var m, s, prevM float64

	for i := 1; i <= len(metrics.Values); i++ {
		x := metrics.Values[i-1].Value
		prevM = m
		m += (x - m) / float64(i)
		s += (x - m) * (x - prevM)
	}

	return math.Sqrt(s / float64(len(metrics.Values)-1)), nil
}
