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

// readLines reads a whole file into memory and returns a slice of its lines.
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

// cleanDupes creates a map and removes duplicates
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

func writeToFile(json *[]uint8) {
	err := ioutil.WriteFile("output.json", *json, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {

	listOfUsers, err := readLines("./sample.txt")
	if err != nil {
		panic(err)
	}

	filteredList := processByTime(&listOfUsers)

	jsonString, err := json.Marshal(filteredList)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonString))

	// write to file
	writeToFile(&jsonString)

}
