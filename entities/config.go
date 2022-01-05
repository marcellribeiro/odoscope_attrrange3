package entities

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
	Id    string  `string:"id"`
	Value float64 `float64:"value"`
	Date  string  `string:"id"`
}
