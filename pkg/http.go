package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// nolint: noctx
func GetHTTPResponse(url string, result interface{}) error {
	bodyBytes, err := GetHTTPResponseBytes(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return nil
}

// nolint: noctx
func GetHTTPResponseBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unsuccessful status code: %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	return bodyBytes, nil
}

// nolint: noctx
func GetHTTPResponseCode(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}

// nolint: noctx
func PostHTTPResponse(url string, payload []byte, result interface{}) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create POST request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do POST request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	err = json.Unmarshal(bodyBytes, result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return nil
}
