package entities

type OutliersResult struct {
	SiteId                  string `json:"siteId"`
	TimeAgo                 string `json:"TimeAgo"`
	TimeStep                string `json:"TimeStep"`
	OutliersDetectionMethod string `json:"OutliersDetectionMethod"`
	checkTimeStart          string `json:"checkTimeStart"`
	checkTimeEnd            string `json:"checkTimeEnd"`
	DateStart               string `json:"DateStart"`
	DateEnd                 string `json:"DateEnd"`
	Result                  Result `json:"Result"`
}

type Result struct {
	Warnings []Warning `json:"Warnings"`
	Alarms   []Alarm   `json:"Alarms"`
}

type Alarm struct {
	OutlierPeriodStart string `json:"OutlierPeriodStart"`
	OutlierPeriodEnd   string `json:"OutlierPeriodEnd"`
	Metric             string `json:"Metric"`
	Attribute          string `json:"Attribute"`
}

type Warning struct {
	OutlierPeriodStart string `json:"OutlierPeriodStart"`
	OutlierPeriodEnd   string `json:"OutlierPeriodEnd"`
	Metric             string `json:"Metric"`
	Attribute          string `json:"Attribute"`
}
