package repository

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/marcellribeiro/odoscope_attrrange3/entities"
)

func MakeSampleData() entities.Datasets {
	attributes := []string{"first_attr"}
	metricesList := []string{"Revenue", "Basket"}
	siteId := "brax"
	timeAgo := 30
	timeStep := 1
	outliersDetectionMethod := "3-sigmas"
	minVisitorsPerTimeStep := 30
	outliersMultiplier := 2
	strongOutliersMultiplier := 3
	minVisitors := minVisitorsPerTimeStep * timeAgo
	sampleRows := 1000

	return entities.Datasets{
		SiteId:                  siteId,
		TimeAgo:                 timeAgo,
		TimeStep:                timeStep,
		OutliersDetectionMethod: outliersDetectionMethod,
		MetricesList:            metricesList,
		MinVisitorsPerTimeStep:  minVisitorsPerTimeStep,
		OutliersDetection: entities.OutliersDetection{
			OutliersMultiplier:       outliersMultiplier,
			StrongOutliersMultiplier: strongOutliersMultiplier,
		},
		Attributes: makeSampleAttributes(minVisitors, sampleRows, attributes, metricesList, timeAgo),
	}
}

func makeSampleAttributes(minVisitors int, sampleRows int, attributes []string, metricesList []string, timeAgo int) []entities.Attributes {
	sampleRowsPerAttributes := sampleRows / len(attributes)                   //1000
	sampleMetricsPerAttributes := sampleRowsPerAttributes / len(metricesList) //500

	attributesResult := make([]entities.Attributes, len(attributes))
	for i := range attributesResult {
		attributesResult[i] = entities.Attributes{
			Name:    attributes[i],
			Metrics: makeSampleMetrics(minVisitors, sampleMetricsPerAttributes, attributes, metricesList, timeAgo),
		}
	}

	return attributesResult
}

func makeSampleMetrics(minVisitors int, sampleMetricsPerAttributes int, attributes []string, metricesList []string, timeAgo int) []entities.Metrics {
	metricsResult := make([]entities.Metrics, len(metricesList))
	for i := range metricsResult {
		metricsResult[i] = entities.Metrics{
			Name:   metricesList[i],
			Values: makeSampleValues(minVisitors, sampleMetricsPerAttributes, attributes, metricesList, timeAgo),
		}
	}

	return metricsResult
}

func makeSampleValues(minVisitors int, sampleMetricsPerAttributes int, attributes []string, metricesList []string, timeAgo int) []entities.MetricsValues {
	valuesResult := make([]entities.MetricsValues, sampleMetricsPerAttributes)
	for i := range valuesResult {
		uid := uuid.NewString()

		valuesResult[i] = entities.MetricsValues{
			Id:    uid,
			Value: makeSampleMetricValue(minVisitors),
			Date:  makeSampleTime(timeAgo),
		}
	}

	return valuesResult
}

func makeSampleMetricValue(minVisitors int) float64 {
	min := 5.0
	max := 100.0
	r := min + rand.Float64()*(max-min)
	return r
}

func makeSampleTime(timeAgo int) string {
	t := time.Now()
	timeAgoNegative := 0 - timeAgo

	max := time.Now().Unix()
	min := t.AddDate(0, 0, timeAgoNegative).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	temporaryTime := time.Unix(sec, 0)

	return fmt.Sprintf("%d", temporaryTime)
}
