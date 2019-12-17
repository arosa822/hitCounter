package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// load up config parameters, this should be the only global setting defined
var configs Params = getConfig()

// getMetrics processes data for the target file specified in configs.json
func getMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, processFile(configs))
}

// spin server spins up the api endpoint
func spinServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getMetrics", getMetrics)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// parseArgs checks the flag at runtime to determine the course of action
func parseArgs() string {
	args := os.Args
	if len(os.Args) == 1 {
		return "NULL"
	}
	return args[1]

}

func main() {

	switch parseArgs() {

	case "-t": // test output without spinnning up server
		fmt.Println(processFile(configs))
	case "-r": // runs the api and spins up the server
		spinServer()
		fmt.Println("Server initialized on port 8080...")

	default:
		fmt.Printf("Usage:\n\t-r: (run) - spins up the server exposing to port 8080" +
			"\n\t-t: (test) - test outputJSON without spinning up the server\n")

	}
}
