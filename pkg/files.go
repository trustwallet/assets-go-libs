package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func IsFileAllowedInPR(path string) bool {
	if strings.HasSuffix(path, "tokenlist.json") {
		return false
	}
	if strings.HasPrefix(path, "blockchains") && strings.Index(path, "assets") > 0 {
		return true
	}
	if strings.HasPrefix(path, "blockchains") && strings.HasSuffix(path, "allowlist.json") {
		return true
	}
	if strings.HasPrefix(path, "blockchains") && strings.HasSuffix(path, "validators/list.json") {
		return true
	}
	if strings.HasPrefix(path, "dapps") {
		return true
	}

	return false
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func CreateDirPath(path string) error {
	dirPath := filepath.Dir(path)

	return os.MkdirAll(dirPath, os.ModePerm)
}

func CreatePNGFromURL(logoURL, logoPath string) error {
	imgBytes, err := GetHTTPResponseBytes(logoURL)
	if err != nil {
		return err
	}

	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return fmt.Errorf("failed to decode image bytes: %v", err)
	}

	out, err := os.Create(logoPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}

	return nil
}

func CreateJSONFile(path string, payload interface{}) error {
	file, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal json: %v", err)
	}

	err = ioutil.WriteFile(path, file, 0600)
	if err != nil {
		return fmt.Errorf("failed to write json to file: %v", err)
	}

	return nil
}

func ReadJSONFile(path string, result interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to read data from file: %v", err)
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return nil
}
