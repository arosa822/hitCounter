package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// DailyMetrics is an embedded json list containing processed data
type DailyMetrics struct {
	DayOfMonth     string `json:"date"`
	TotalHits      int    `json:"totalHits"`
	UniqueVisitors int    `json:"unique"`
	RepeatVisits   int    `json:"repeat"`
}

// UserData is the main json struct
type UserData struct {
	Users          []string       `json:"users"`
	TimeVisit      [][]int        `json:"timeStamp"`
	HitCount       int            `json:"hitCount"`
	UniqueVisitors int            `json:"uniqueVisits"`
	Data           []DailyMetrics `json:"data"`
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
	return int64(mostRecent)
}

// processByDay processes UserData.TimeVisit and populates DailyMetrics struct
func (data *UserData) processByDay(configs Params) {

	var total, unique, repeat int = 0, 0, 0
	daily := make(map[string]DailyMetrics)
	mostRecentEntry := time.Unix(data.findMostRecent(), 0)
	query, err := strconv.Atoi(configs.Days)

	if err != nil {
		panic(err)
	}
	count := 0

	// go through each day
	for n := 0; n < query; n++ {
		// start at the top and subtract a day
		timeObject := mostRecentEntry.AddDate(0, 0, (-1 * n))
		_, m, d := timeObject.Date()
		date := fmt.Sprintf("%v-%d", m, d)
		// traverse the 2d array
		// each list in list is specific to a unique user
		// this must be done for each day in span created above
		for _, j := range data.TimeVisit {
			for _, k := range j {
				// check if item in the list matches wrapping day
				if time.Unix(int64(k), 0).Day() == d {
					total++
					unique++
					repeat = FindDupesInArray(j, d)
				} else {
					continue
				}
			}
		}
		// if anything over 1 visit means it is a repeat visit
		if unique > 1 {
			unique = unique - repeat
		}
		// create the map object with data fields
		daily[strconv.Itoa(count)] = DailyMetrics{DayOfMonth: date, TotalHits: total, UniqueVisitors: unique, RepeatVisits: repeat}
		// push the map object into data.Data slice
		data.Data = append(data.Data, daily[strconv.Itoa(count)])

		total, repeat, unique = 0, 0, 0
		count++
	}
}

// FindDupesInArray returns the number of duplicate entries
// takes a list of unix times and specified day
func FindDupesInArray(array []int, day int) int {
	count := 0
	for i := 0; i < len(array); i++ {
		if time.Unix(int64(array[i]), 0).Day() == day {
			count++
		}
	}
	return count
}

// readLines reads a whole file into memory and returns a list of strings
// containing each line.
func readLines(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

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

func writeToFile(json *[]uint8, location string) {
	err := ioutil.WriteFile(location, *json, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	var jsonString []byte

	// load up config parameters
	configs := getConfig()

	// read in data
	listOfUsers, err := readLines(configs.File)
	if err != nil {
		panic(err)
	}

	// filter data and create a map object
	filteredList := processByTime(&listOfUsers)

	//  convert map into data structure
	siteVisits := convertMapToStruct(&filteredList)

	// process data using methods
	siteVisits.processByDay(configs)

	// encode structure into json format
	jsonString, err = json.Marshal(siteVisits)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonString))

	// write to file
	writeToFile(&jsonString, configs.Output)
}
