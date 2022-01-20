package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHTTPResponse(url string, result interface{}) error {
	bodyBytes, err := GetHTTPResponseBytes(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return nil
}

// nolint
func GetHTTPResponseBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unsuccessful status code: %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	return bodyBytes, nil
}
