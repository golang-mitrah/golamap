package golamap

import (
	"encoding/json"
	"io"
	"net/http"
)

func MakeExternalRequest(method, url, requestID, oauthToken string, responseObj interface{}) error {
	// Create a new request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Request-Id", requestID)
	if oauthToken != "" {
		req.Header.Add("Authorization", oauthToken)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse the JSON response
	err = json.Unmarshal(body, responseObj)
	if err != nil {
		return err
	}

	return nil
}

// ParseJSONBody parses the JSON request body into the provided struct.
func ParseJSONBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
