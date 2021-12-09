package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// nolint: noctx
func GetHTTPResponse(url string, v interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("unsuccessful status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return nil
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

func GetHTTPResponseBytes(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	return bytes, nil
}
