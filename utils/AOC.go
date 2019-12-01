package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getFilename() string {
	return "input.txt"
}

func readFromFile() string {
	data, err := ioutil.ReadFile(getFilename())
	check(err)

	return strings.Trim(strings.Trim(string(data), "\n"), "\r")
}

func saveToFile(data string) {
	fileName := getFilename()
	err := ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		log.Println("Error writing to " + fileName + ", err: ")
		log.Println(err)
	}
}

func getSessionID() string {
	args := os.Args[1:]

	for index, arg := range args {
		if arg == "--session" {
			if len(args) >= index+1 {
				return args[index+1]
			}
		}
	}

	sID := os.Getenv("aoc_sessionid")
	if sID > "" {
		return sID
	}

	panic("No advent of code sessionId found, either set 'aoc_sessionid' env variable or specify it with --input SESSION_ID")
}

func FetchInput(year, day int) string {
	if fileExists(getFilename()) {
		log.Println("Reusing data from file")
		return readFromFile()
	}
	client := &http.Client{}

	url := "https://adventofcode.com/" + strconv.Itoa(year) + "/day/" + strconv.Itoa(day) + "/input"
	log.Println("Fetching from ", url)
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.Header.Set("Cookie", "session="+getSessionID()+";")

	resp, err := client.Do(req)
	check(err)
	data, err := ioutil.ReadAll(resp.Body)
	check(err)
	if resp.StatusCode != 200 {
		log.Println("Got response: ", data)
		log.Println("Invalid response code ", resp.StatusCode)

		panic("Error fetching input, got code " + strconv.Itoa(resp.StatusCode))
	}

	saveToFile(string(data))
	return strings.Trim(strings.Trim(string(data), "\n"), "\r")
}
