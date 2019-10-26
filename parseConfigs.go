package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Params struct {
	Days string `json:"days"`
	File string `json:"file"`
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
