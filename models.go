package main

// Params are all the file configurations passeed in during runtime
type Params struct {
	Days   string `json:"days"`
	File   string `json:"file"`
	Output string `json:"outputFile"`
}

// DailyMetrics is an embedded json list containing processed data
type DailyMetrics struct {
	DayOfMonth     int64 `json:"date"`
	TotalHits      int   `json:"totalHits"`
	UniqueVisitors int   `json:"unique"`
	RepeatVisits   int   `json:"repeat"`
}

// UserData is the main json struct
type UserData struct {
	Head           []string `json:"head"`
	Users          []string `json:"_"`
	TimeVisit      [][]int  `json:"_"`
	HitCount       int      `json:"_"`
	UniqueVisitors int      `json:"_"`
	Data           [][]int  `json:"data"`
}
