package entities

import (
	"fmt"
	"time"
)

type OutliersResult struct {
	SiteId                  string `json:"siteId"`
	TimeAgo                 int    `json:"TimeAgo"`
	TimeStep                int    `json:"TimeStep"`
	OutliersDetectionMethod string `json:"OutliersDetectionMethod"`
	CheckTimeStart          string `json:"checkTimeStart"`
	CheckTimeEnd            string `json:"checkTimeEnd"`
	DateStart               string `json:"DateStart"`
	DateEnd                 string `json:"DateEnd"`
	Result                  Result `json:"Result"`
}

type Result struct {
	Warnings []Warning `json:"Warnings"`
	Alarms   []Alarm   `json:"Alarms"`
}

type Results []Result

type Alarm struct {
	OutlierPeriodStart time.Time `json:"OutlierPeriodStart"`
	OutlierPeriodEnd   time.Time `json:"OutlierPeriodEnd"`
	Metric             string    `json:"Metric"`
	Attribute          string    `json:"Attribute"`
}

type Alarms []Alarm

func (a Alarms) Len() int {
	return len(a)
}

func (a Alarms) Less(i, j int) bool {
	if a[i].Attribute == a[j].Attribute {
		return a[i].OutlierPeriodStart.Before(a[j].OutlierPeriodStart)
	}
	return a[i].Attribute < a[j].Attribute
}

func (a Alarms) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Warning struct {
	OutlierPeriodStart time.Time `json:"OutlierPeriodStart"`
	OutlierPeriodEnd   time.Time `json:"OutlierPeriodEnd"`
	Metric             string    `json:"Metric"`
	Attribute          string    `json:"Attribute"`
}

type Warnings []Warning

func (w Warnings) Len() int {
	return len(w)
}

func (w Warnings) Less(i, j int) bool {
	if w[i].Attribute == w[j].Attribute {
		return w[i].OutlierPeriodStart.Before(w[j].OutlierPeriodStart)
	}
	return w[i].Attribute < w[j].Attribute
}

func (w Warnings) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func AddAlarm(attributeName, metricName string, value MetricsValues) *Alarm {
	return &Alarm{
		OutlierPeriodStart: value.Date,
		OutlierPeriodEnd:   value.Date,
		Metric:             fmt.Sprintf("%s(%s)", metricName, value.Id),
		Attribute:          attributeName,
	}
}

func AddWarning(attributeName, metricName string, value MetricsValues) *Warning {
	return &Warning{
		OutlierPeriodStart: value.Date,
		OutlierPeriodEnd:   value.Date,
		Metric:             fmt.Sprintf("%s(%s)", metricName, value.Id),
		Attribute:          attributeName,
	}
}
