package entities

type Config struct {
	Datasets []Datasets `slicetructure:"Datasets"`
}

type Datasets struct {
	SiteId                  string              `string:"siteId"`
	TimeAgo                 string              `string:"TimeAgo"`
	TimeStep                string              `string:"TimeStep"`
	OutliersDetectionMethod string              `string:"OutliersDetectionMethod"`
	MetricesList            []string            `slicetructure:"MetricesList"`
	MinVisitorsPerTimeStep  int                 `int:"MinVisitorsPerTimeStep"`
	OutliersDetection       []OutliersDetection `slicetructure:"OutliersDetection"`
}

type OutliersDetection struct {
	OutliersMultiplier       int `int:"OutliersMultiplier"`
	StrongOutliersMultiplier int `int:"StrongOutliersMultiplier"`
}
