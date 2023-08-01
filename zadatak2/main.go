package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
)

type Input struct {
	Operation	string 	`json:"operation"`
	Data		[]int64 `json:"data"`
}

type OutputDeduplicate struct {
	ID			string  `json:"id"`
	Operation	string 	`json:"operation"`
	Data		[]int64 `json:"data"`
}

type OutputGetPairs struct {
	ID			string  `json:"id"`
	Operation	string 	`json:"operation"`
	Data		[][]int64 `json:"data"`
}

func uploadAndProcessData(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is POST (or any other method you expect to send JSON data).
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    // Parse the JSON data from the request body.
    var dataMap Input
    err := json.NewDecoder(r.Body).Decode(&dataMap)
    if err != nil {
        http.Error(w, "Error parsing JSON", http.StatusBadRequest)
        return
    }
	fmt.Println("dataMap :", dataMap)
	operation := dataMap.Operation
	data := dataMap.Data
	fmt.Println("uploadAndProcessData: ", operation, data)

	if operation != "deduplicate" && operation != "getPairs" {
		fmt.Fprintln(w, "Operation " + operation + " does not exist!")
		fmt.Fprintln(w, "Valid operations are: deduplicate and getPairs")
		return
	}

	var output any

	if operation == "deduplicate" {
		output = deduplicate(data, operation)
	} else if operation == "getPairs" {
		output = getPairs(data, operation)
	}



	outputJSON, err := json.Marshal(output)
    if err != nil {
        http.Error(w, "Error parsing JSON", http.StatusBadRequest)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	w.Write(outputJSON)
}

func deduplicate(data []int64, operation string)OutputDeduplicate {
	m := make(map[int64]int8)
	currIndex := 0
	for i := 0; i < len(data); i++ {
		num := m[data[i]]
		if num == 0 {
			m[data[i]] = 1
			data[currIndex] = data[i]
			currIndex++
		}
	}
	var output OutputDeduplicate
	output.ID = uuid.New().String()
	output.Data = data[:currIndex]
	output.Operation = operation
	return output
}

func getPairs(data []int64, operation string)OutputGetPairs {
	m := make(map[int64]int64)
	for i := 0; i < len(data); i++ {
		m[data[i]]++
	}
	dataProcessed := [][]int64{}
	for key, cnt := range m {
		dataProcessed = append(dataProcessed, []int64{key, cnt})
    }
	var output OutputGetPairs
	output.ID = uuid.New().String()
	output.Data = dataProcessed
	output.Operation = operation
	return output
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadAndProcessData)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Server started!")
	setupRoutes()
}