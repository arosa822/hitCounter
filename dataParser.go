package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type DailyMetrics struct {
	DayOfMonth     string `json:"date"`
	TotalHits      int    `json:"totalHits"`
	UniqueVisitors int    `json:"unique"`
	RepeatVisits   int    `json:"repeat"`
}

type UserData struct {
	Users          []string       `json:"users"`
	TimeVisit      [][]int        `json:"timeStamp"`
	HitCount       int            `json:"hitCount"`
	UniqueVisitors int            `json:"uniqueVisits"`
	Data           []DailyMetrics `'json:"T-7D"`
}

// readLines reads a whole file into memory and returns a list of strings
// containing each line.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// cleanDupes maps a key value pair of users and unixTime specific to each user
func processByTime(list *[]string) map[string][]int {
	mapOfUsers := map[string][]int{}

	for _, line := range *list {
		entry := strings.Split(line, ",")

		// entry must have an IP and time in order to be added to map
		if len(entry) < 2 {
			continue
		} else {
			// remove whitespace
			entry[1] = strings.Join(strings.Fields(entry[1]), "")
			// convert unix string to int
			unixTime, err := strconv.Atoi(entry[1])
			if err != nil {
				panic(err)
			}

			// append the map: entry[0] =  IP
			// duplicates are appended to the key value entry[0]
			mapOfUsers[entry[0]] = append(mapOfUsers[entry[0]], unixTime)

		}
	}

	return mapOfUsers
}

// convertMapToStruct takes in a map object as an argument and returns a
// structure of type UserData
func convertMapToStruct(unstructData *map[string][]int) UserData {
	//data := make(map[string]UserData)
	var data UserData
	count := 0

	// iterate over the map and throw in struct
	for key, value := range *unstructData {
		data.Users = append(data.Users, key)
		data.TimeVisit = append(data.TimeVisit, value)
		// add the length of each slice containing a time stamp to the count
		count = count + len(value)

	}

	data.UniqueVisitors = len(data.Users)
	data.HitCount = count
	return data
}

func writeToFile(json *[]uint8) {
	err := ioutil.WriteFile("output.json", *json, 0644)
	if err != nil {
		panic(err)
	}
}

// sorting method for UserData to find the most recent time entry
func (data *UserData) findMostRecent() int64 {
	mostRecent := data.TimeVisit[0][0]

	for _, i := range data.TimeVisit {
		for _, j := range i {
			if j > mostRecent {
				mostRecent = j
			}
		}
	}

	// debug code
	// fmt.Println(mostRecent)
	// mostRecentTimeObj := time.Unix(int64(mostRecent), 0)
	// fmt.Println(mostRecentTimeObj)
	// fmt.Println(mostRecentTimeObj.Date())

	return int64(mostRecent)
}

func main() {
	var jsonString []byte

	// load up config parameters
	configs := getConfig()

	listOfUsers, err := readLines(configs.File)
	if err != nil {
		panic(err)
	}

	filteredList := processByTime(&listOfUsers)

	siteVisits := convertMapToStruct(&filteredList)

	jsonString, err = json.Marshal(siteVisits)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonString))

	// write to file
	writeToFile(&jsonString)

	siteVisits.findMostRecent()

}
