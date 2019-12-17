package main

// Params are all the file configurations passeed in during runtime
type Params struct {
	Days   string `json:"days"`
	File   string `json:"file"`
	Output string `json:"outputFile"`
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
