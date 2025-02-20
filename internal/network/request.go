package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func HttpPost(url string, req map[string]interface{}) (*http.Response, error) {
	// Convert the request map to JSON
	jsonReq, err := json.Marshal(req)
	if err != nil {
		//utils.Print.Errorf("fail to convert req to json")
		return nil, fmt.Errorf("received nil req")
	}

	// Create a new POST request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		//utils.Print.Errorf("fail to create req")
		return nil, fmt.Errorf("req create failed")
	}

	// Set the Content-Type header to application/json
	request.Header.Set("Content-Type", "application/json")

	// Create a client and send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		//utils.Print.Errorf("fail to sent req")
		return nil, fmt.Errorf("req sent failed")
	}
	// log.Println("HttpPost")
	// log.Println(response)
	return response, nil
}
