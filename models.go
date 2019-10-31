package main

// DailyMetrics is an embedded json list containing processed data
type DailyMetrics struct {
	DayOfMonth     int64 `json:"date"`
	TotalHits      int   `json:"totalHits"`
	UniqueVisitors int   `json:"unique"`
	RepeatVisits   int   `json:"repeat"`
}

// UserData is the main json struct
type UserData struct {
	Users          []string       `json:"_"`
	TimeVisit      [][]int        `json:"_"`
	HitCount       int            `json:"hitCount"`
	UniqueVisitors int            `json:"uniqueVisits"`
	Data           []DailyMetrics `json:"data"`
}
