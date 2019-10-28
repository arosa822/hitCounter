package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Params are all the file configurations from configs.json
type Params struct {
	Days   string `json:"days"`
	File   string `json:"file"`
	Output string `json:"outputFile"`
}

func getConfig() Params {

	var params Params

	jsonFile, err := os.Open("configs.json")

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(bytes, &params)

	return params
}
