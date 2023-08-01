// main_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
	"encoding/json"
)

func TestUploadAndProcessDataDeduplicate(t *testing.T) {
    // Create a mock request with JSON data for testing.
    requestBody := `{"Operation": "deduplicate", "Data": [1,2,2,3,3,3,4,4,4,4,6,8,8,9,6,6,7]}`
    req, err := http.NewRequest("POST", "/upload", strings.NewReader(requestBody))
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response.
    rr := httptest.NewRecorder()


    // Call the handler function.
    uploadAndProcessData(rr, req)

    // Check the HTTP status code.
    if rr.Code != http.StatusOK {
        t.Errorf("Expected status 200 OK, got %d", rr.Code)
    }

    // Decode the JSON response into a variable.
    var dataMap OutputDeduplicate
    if err := json.Unmarshal(rr.Body.Bytes(), &dataMap); err != nil {
        t.Errorf("Error parsing JSON: %v", err)
    }

    // Validate the JSON response.
    expectedOperation := "deduplicate"
    expectedData := []int64{1, 2, 3, 4, 6, 8, 9, 7}
    if dataMap.Operation != expectedOperation {
        t.Errorf("Expected operation %q, but got %q", expectedOperation, dataMap.Operation)
    }
	if len(expectedData) != len(dataMap.Data) {
		t.Errorf("Expected data %v, but got %v", expectedData, dataMap.Data)
	}
    for i, v := range dataMap.Data {
        if v != expectedData[i] {
            t.Errorf("Expected data %v, but got %v", expectedData, dataMap.Data)
        }
    }
}

// Add more test cases for deduplicate and getPairs functions here.
