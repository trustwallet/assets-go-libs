package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func GetHTTPResponse(url string, v interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "failed to make GET request")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("failed to obtain json")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read bytes from body")
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal json")
	}

	return nil
}

func GetHTTPResponseCode(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, errors.Wrap(err, "failed to make GET request")
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}
