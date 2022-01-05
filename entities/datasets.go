package entities

import (
	"context"
	"time"
)

type Datasets struct {
	SiteId                  string            `string:"siteId"`
	TimeAgo                 int               `int:"TimeAgo"`
	TimeStep                int               `int:"TimeStep"`
	OutliersDetectionMethod string            `string:"OutliersDetectionMethod"`
	MetricesList            []string          `slicetructure:"MetricesList"`
	MinVisitorsPerTimeStep  int               `int:"MinVisitorsPerTimeStep"`
	OutliersDetection       OutliersDetection `slicetructure:"OutliersDetection"`
	Attributes              []Attributes      `slicetructure:"Attributes"`
}

type OutliersDetection struct {
	OutliersMultiplier       int `int:"OutliersMultiplier"`
	StrongOutliersMultiplier int `int:"StrongOutliersMultiplier"`
}

type Attributes struct {
	Name    string    `string:"Name"`
	Metrics []Metrics `slicetructure:"Metrics"`
}

type Metrics struct {
	Name   string          `string:"Name"`
	Values []MetricsValues `slicetructure:"Values"`
}

type MetricsValues struct {
	Id    string    `string:"id"`
	Value float64   `float64:"value"`
	Date  time.Time `string:"date"`
}

type MetricsList []MetricsValues

func (m MetricsList) Len() int {
	return len(m)
}

func (m MetricsList) Less(i, j int) bool {
	return m[i].Date.Before(m[j].Date)
}

func (m MetricsList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type CalculatorUseCase interface {
	CalculatorService(ctx context.Context) (OutliersResult, error)
}
