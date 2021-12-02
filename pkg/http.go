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
		return fmt.Errorf("failed to read bytes from body: %v", err)
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
