// Assisted by watsonx Code Assistant

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

// Counter holds the current value to be incremented
var Counter int = 0

// IncrementHandler handles requests to increment the counter
func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	var input string
	if err := readRequestBody(r, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse input to int
	value, err := strconv.Atoi(input)
	if err != nil {
		http.Error(w, "Invalid input. Please provide a number.", http.StatusBadRequest)
		return
	}

	// Increment counter
	Counter += value

	// Respond with the new counter value
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Counter: %d", Counter)
}

// readRequestBody reads the request body into the provided pointer
func readRequestBody(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	str := string(body)
	err = json.Unmarshal([]byte(str), v)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	http.HandleFunc("/increment", IncrementHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
