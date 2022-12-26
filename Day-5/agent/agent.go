package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Job struct {
	Id   int
	Name string
}

type Result struct {
	Id     int
	Status string
}

func execute(job Job, c chan Result) {
	log.Println("executing ", job.Name)

	// Write to the channel
	c <- Result{job.Id, "SUCCESS"}
}

func getDataFromJson(filepath string) []Job {
	// Let's first read the `jobs.json` file
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload []Job
	json.Unmarshal(content, &payload)
	return payload
}

func sendDataToJson(filepath string, data []Result) {
	var marshalDataJson, _ = json.Marshal(data)
	ioutil.WriteFile(filepath, marshalDataJson, 0644)
}

func main() {
	jsonData := getDataFromJson("./jobs.json")
	dataLength := len(jsonData)

	channel := make(chan Result)

	for i := 0; i < dataLength; i++ {
		go execute(jsonData[i], channel)
	}

	var results []Result
	for i := 0; i < dataLength; i++ {
		results = append(results, <-channel)
	}

	sendDataToJson("./status.json", results)
}
