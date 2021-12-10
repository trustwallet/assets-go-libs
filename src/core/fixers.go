package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/trustwallet/assets-go-libs/pkg/file"
)

func (s *Service) FixInfoJSON(file *file.AssetFile) error {
	jsonFile, err := os.Open(file.Info.Path())
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, data, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file.Info.Path(), prettyJSON.Bytes(), 0644)
}
