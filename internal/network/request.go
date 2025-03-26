package network

import (
	"bytes"
	"io"
	"net/http"
)

func HttpPost(url string, req []byte) ([]byte, error) {

	// Create a new POST request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(req))
	if err != nil {
		//utils.Print.Errorf("fail to create req")
		return nil, err
	}

	// Set the Content-Type header to application/json
	request.Header.Set("Content-Type", "application/json")

	// Create a client and send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		//utils.Print.Errorf("fail to sent req")
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	// log.Println("HttpPost")
	// log.Println(response)
	return body, nil
}
