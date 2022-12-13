package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	fileModeReadWrite = 0600 //nolint
	indent            = "    "
	prefix            = ""
)

func PrepareJSONData(payload interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(payload, prefix, indent)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}

	// The solution of escaping special HTML characters in golang json.marshal.
	data = bytes.ReplaceAll(data, []byte("\\u003c"), []byte("<"))
	data = bytes.ReplaceAll(data, []byte("\\u003e"), []byte(">"))
	data = bytes.ReplaceAll(data, []byte("\\u0026"), []byte("&"))

	return data, nil
}

func CreateJSONFile(path string, data []byte) error {
	err := os.WriteFile(path, data, fileModeReadWrite)
	if err != nil {
		return fmt.Errorf("failed to write json to file: %w", err)
	}

	return nil
}

func ReadJSONFile(path string, result interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read data from file: %w", err)
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return nil
}

func FormatJSONFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read data from file: %w", err)
	}

	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, data, prefix, indent)
	if err != nil {
		return fmt.Errorf("failed to indent json: %w", err)
	}

	err = os.WriteFile(path, prettyJSON.Bytes(), fileModeReadWrite)
	if err != nil {
		return fmt.Errorf("failed to write json to file: %w", err)
	}

	return nil
}
